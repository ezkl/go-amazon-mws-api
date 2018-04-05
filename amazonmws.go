// amazonmws provides methods for interacting with the Amazon Marketplace Services API.
package amazonmws

import (
	"fmt"
	"log"
	"strconv"
)

type RequestParams interface {
	ToQueryParams() (map[string]string, error)
	SetPrefixKey(string)
}

type FeeEstimateRequest struct {
	IdValue             string
	PriceToEstimateFees float64
	Currency            string
	MarketplaceId       string
	IdType              string
	Identifier          string
	IsAmazonFulfilled   bool
	prefixKey           string
}

type LowestOfferListingsForASIN struct {
	MarketplaceId string
	ASIN          string
	ItemCondition string
}

func (f *FeeEstimateRequest) setDefaults(marketplaceId string, isFba bool) {
	if f.Currency == "" {
		f.Currency = "USD"
	}

	if f.MarketplaceId == "" {
		f.MarketplaceId = marketplaceId
	}

	if f.IdType == "" {
		f.IdType = "ASIN"
	}

	if f.Identifier == "" {
		f.Identifier = f.IdValue
	}

	f.IsAmazonFulfilled = isFba
}

func (f *FeeEstimateRequest) SetPrefixKey(key string) {
	f.prefixKey = key
}

func (f *FeeEstimateRequest) ToQueryParams() (map[string]string, error) {
	params := make(map[string]string)

	key := f.prefixKey
	if key == "" {
		return nil, fmt.Errorf("this endpoint requires a prefix key, likeFeesEstimateRequestList.FeesEstimateRequest.1")
	}

	if f.MarketplaceId == "" {
		return nil, fmt.Errorf("MarketplaceId cannot be empty")
	}

	if f.IdType == "" {
		return nil, fmt.Errorf("IdType cannot be empty")
	}

	params[key+".MarketplaceId"] = f.MarketplaceId
	params[key+".IdType"] = f.IdType
	params[key+".IdValue"] = f.IdValue
	params[key+".Identifier"] = f.Identifier
	params[key+".PriceToEstimateFees.ListingPrice.Amount"] = strconv.FormatFloat(f.PriceToEstimateFees, 'f', 2, 32)
	params[key+".PriceToEstimateFees.ListingPrice.CurrencyCode"] = f.Currency

	// which one should have a priority?
	if f.IsAmazonFulfilled {
		params[key+".IsAmazonFulfilled"] = "true"
	} else {
		params[key+".IsAmazonFulfilled"] = "false"
	}

	return params, nil
}

func (l *LowestOfferListingsForASIN) ToQueryParams() (map[string]string, error) {
	params := make(map[string]string, 3)

	params["MarketplaceId"] = l.MarketplaceId
	params["ASIN"] = l.ASIN
	params["ItemCondition"] = l.ItemCondition

	return params, nil
}

func (l *LowestOfferListingsForASIN) SetPrefixKey(key string) {
	log.Fatalln("Method SetPrefixKey not implemented for struct LowestOfferListingsForASIN")
}

/*
GetLowestOfferListingsForASIN takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetLowestOfferListingsForASIN(items []string) (string, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.genSignAndFetch("GetLowestOfferListingsForASIN", "/Products/2011-10-01", params)
}

/*
GetCompetitivePricingForAsin takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetCompetitivePricingForASIN(items []string) (string, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.genSignAndFetch("GetCompetitivePricingForASIN", "/Products/2011-10-01", params)
}

func (api AmazonMWSAPI) GetMatchingProductForId(idType string, idList []string) (string, error) {
	params := make(map[string]string)

	for k, v := range idList {
		key := fmt.Sprintf("IdList.Id.%d", (k + 1))
		params[key] = string(v)
	}

	params["IdType"] = idType
	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.genSignAndFetch("GetMatchingProductForId", "/Products/2011-10-01", params)
}

func (api AmazonMWSAPI) GetMyFeesEstimate(items []RequestParams) (string, error) {
	params, err := createPrefixedRequestParams("FeesEstimateRequestList.FeesEstimateRequest.%d", items)
	if err != nil {
		return "", err
	}

	return api.genSignAndFetch("GetMyFeesEstimate", "/Products/2011-10-01", params)
}

func (api AmazonMWSAPI) GetLowestPricedOffersForASIN(item RequestParams) (string, error) {
	params, err := item.ToQueryParams()
	if err != nil {
		return "", err
	}

	return api.genSignAndFetch("GetLowestPricedOffersForASIN", "/Products/2011-10-01", params)
}

var createPrefixedRequestParams = func(prefix string, items []RequestParams) (map[string]string, error) {
	params := make(map[string]string)

	for index, item := range items {
		key := fmt.Sprintf(prefix, (index + 1))
		item.SetPrefixKey(key)

		queryParams, err := item.ToQueryParams()
		if err != nil {
			return nil, err
		}

		for key, param := range queryParams {
			params[key] = param
		}
	}

	return params, nil
}
