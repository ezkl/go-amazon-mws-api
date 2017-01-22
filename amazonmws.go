// amazonmws provides methods for interacting with the Amazon Marketplace Services API.
package amazonmws

import (
	"fmt"
	"strconv"
	"log"
)

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


/*
THIS IS OUR SELLZEE PROP STUFF
 */
func (api AmazonMWSAPI) GetLowestPricedOffersForASIN(asin string, itemCondition string) (string, error) {
	params := make(map[string]string)

	params["ASIN"] = asin
	params["MarketplaceId"] = string(api.MarketplaceId)
	params["ItemCondition"] = itemCondition

	return api.genSignAndFetch("GetLowestPricedOffersForASIN", "/Products/2011-10-01", params)
}


func (api AmazonMWSAPI) GetMyFeesEstimate(items []string, uuid string, listingPrice float64, shipping float64) (string, error) {
	params := make(map[string]string)

	listingPriceStr := fmt.Sprintf("%.2f", listingPrice)
	shippingStr := fmt.Sprintf("%.2f", shipping)
	log.Println("---- DEBUG ----")
	log.Println(listingPriceStr)
	log.Println(shippingStr)
	log.Println("---- ----")
	for c := 0; c < len(items); c++ {
		d := strconv.Itoa(c + 1)
		asin := items[c]
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".MarketplaceId"] = string(api.MarketplaceId)
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".IdType"] = "ASIN"
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".IdValue"] = asin
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".IsAmazonFulfilled"] = "true"
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".Identifier"] = uuid
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".PriceToEstimateFees.ListingPrice.CurrencyCode"] = "USD"
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".PriceToEstimateFees.ListingPrice.Amount"] = listingPriceStr
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".PriceToEstimateFees.Shipping.CurrencyCode"] = "USD"
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".PriceToEstimateFees.Shipping.Amount"] = shippingStr
		params["FeesEstimateRequestList.FeesEstimateRequest." + d + ".PriceToEstimateFees.Points.PointsNumber"] = "0"
	}

	return api.genSignAndFetch("GetMyFeesEstimate", "/Products/2011-10-01", params)
}
