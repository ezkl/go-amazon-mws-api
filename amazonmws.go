// amazonmws provides methods for interacting with the Amazon Marketplace Services API.
package amazonmws

import (
	"fmt"
	"bytes"
	"strconv"
)

type FeeEstimateRequest struct {
	IdValue string
	PriceToEstimateFees float64
	Currency string
	MarketplaceId string
	IdType string
	Identifier string
	IsAmazonFulfilled bool
}

func (f *FeeEstimateRequest) requestString(index int, key string) string {
	var buffer bytes.Buffer
	buffer.WriteString("FeesEstimateRequestList.FeesEstimateRequest.")
	buffer.WriteString(strconv.Itoa(index))
	buffer.WriteString(".")
	buffer.WriteString(key)
	return buffer.String()
}

func (f *FeeEstimateRequest) setDefaults(mid string) {
	if f.Currency == "" {
		f.Currency = "USD"
	}

	if f.MarketplaceId == "" {
		f.MarketplaceId = mid
	}

	if f.IdType == "" {
		f.IdType = "ASIN"
	}

	if f.Identifier == "" {
		f.Identifier = f.IdValue
	}

	f.IsAmazonFulfilled = true
}

func (f *FeeEstimateRequest) toQuery(index int, marketplaceId string) map[string]string {
	output := make(map[string]string)

	f.setDefaults(marketplaceId)
	output[f.requestString(index, "IdValue")] = f.IdValue
	output[f.requestString(index, "PriceToEstimateFees.ListingPrice.CurrencyCode")] = f.Currency
	output[f.requestString(index, "PriceToEstimateFees.ListingPrice.Amount")] = strconv.FormatFloat(f.PriceToEstimateFees, 'f', 2, 32)
	output[f.requestString(index, "MarketplaceId")] = f.MarketplaceId
	output[f.requestString(index, "IdType")] = f.IdType
	output[f.requestString(index, "Identifier")] = f.Identifier

	var isFba string
	if (f.IsAmazonFulfilled) {
		isFba = "1"
	} else {
		isFba = "0"
	}

	output[f.requestString(index, "IsAmazonFulfilled")] = isFba

	fmt.Printf("%#v", output);

	return output
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

func (api AmazonMWSAPI) GetMyFeesEstimate(isFba bool, items []FeeEstimateRequest) (string, error) {
	params := make(map[string]string)

	fmt.Println(params);

	return "", nil
}