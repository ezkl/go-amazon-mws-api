//Package amazonmws provides methods for interacting with the Amazon Marketplace Services API
package amazonmws

import (
	"fmt"
)

func (api AmazonMWSAPI) GetLowestOfferListingsForASIN(items []string) (string, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceId)

	return api.genSignAndFetch("GetLowestOfferListingsForASIN", params)
}
