package amazonmws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"sort"
	"bytes"
	//"bitbucket.org/zombiezen/cardcpx/natsort"
)

type AmazonMWSAPI struct {
	AccessKey     string
	SecretKey     string
	Host          string
	AuthToken     string
	MarketplaceId string
	SellerId      string
}

func (api AmazonMWSAPI) genSignAndFetch(Action string, ActionPath string, Parameters map[string]string) (string, error) {
	genUrl, err := GenerateAmazonUrl(api, Action, ActionPath, Parameters)
	if err != nil {
		return "", err
	}

	SetTimestamp(genUrl)

	signedurl, err := SignAmazonUrl(genUrl, api)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(signedurl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GenerateAmazonUrl(api AmazonMWSAPI, Action string, ActionPath string, Parameters map[string]string) (finalUrl *url.URL, err error) {
	result, err := url.Parse(api.Host)
	if err != nil {
		return nil, err
	}

	result.Host = api.Host
	result.Scheme = "https"
	result.Path = ActionPath

	values := url.Values{}
	values.Add("Action", Action)

	if (api.AuthToken != "") {
		values.Add("MWSAuthToken", api.AuthToken)
	}

	values.Add("AWSAccessKeyId", api.AccessKey)
	values.Add("SellerId", api.SellerId)
	values.Add("SignatureVersion", "2")
	values.Add("SignatureMethod", "HmacSHA256")
	values.Add("Version", "2011-10-01")

	for k, v := range Parameters {
		values.Set(k, v)
	}

	params := values.Encode()
	result.RawQuery = params

	return result, nil
}

func SetTimestamp(origUrl *url.URL) (err error) {
	values, err := url.ParseQuery(origUrl.RawQuery)
	if err != nil {
		return err
	}
	values.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))
	origUrl.RawQuery = values.Encode()

	return nil
}

func SignAmazonUrl(origUrl *url.URL, api AmazonMWSAPI) (signedUrl string, err error) {
	escapeUrl := strings.Replace(origUrl.RawQuery, ",", "%2C", -1)
	escapeUrl = strings.Replace(escapeUrl, ":", "%3A", -1)

	params := strings.Split(escapeUrl, "&")
	paramMap := make(map[string]string)
	keys := make([]string, len(params))

	for k, v := range params {
		parts := strings.Split(v, "=")
		paramMap[parts[0]] = parts[1]
		keys[k] = parts[0]
	}
	sort.Strings(keys)

	sortedParams := make([]string, len(params))
	for k, _ := range params {
		var buffer bytes.Buffer
		buffer.WriteString(keys[k])
		buffer.WriteString("=")
		buffer.WriteString(paramMap[keys[k]])
		sortedParams[k] = buffer.String()
	}

	stringParams := strings.Join(sortedParams, "&")

	toSign := fmt.Sprintf("GET\n%s\n%s\n%s", origUrl.Host, origUrl.Path, stringParams)

	hasher := hmac.New(sha256.New, []byte(api.SecretKey))
	_, err = hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	newParams := fmt.Sprintf("%s&Signature=%s", stringParams, hash)

	origUrl.RawQuery = newParams

	return origUrl.String(), nil
}
