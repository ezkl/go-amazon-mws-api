package amazonmws

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

type requestParamsMock struct {
	toQueryParamsHasBeenCalled int
	setPrefixKeyHasBeenCalled  int
	toQueryParamsReturnValues  func() (map[string]string, error)
}

func (m *requestParamsMock) ToQueryParams() (map[string]string, error) {
	m.toQueryParamsHasBeenCalled++
	return m.toQueryParamsReturnValues()
}

func (m *requestParamsMock) SetPrefixKey(key string) {
	m.setPrefixKeyHasBeenCalled++
}

func TestSimpleSignAmazonUrl(t *testing.T) {
	urlString := "https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.2=1561712027&ASINList.ASIN.10=1561711969"
	signedUrl := "https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.10=1561711969&ASINList.ASIN.2=1561712027&Signature=5aNYdV%2Fuc%2B98P2KwX8NhR6tzvKaHBHrLq5F4D7A9iH0%3D"

	var api AmazonMWSAPI

	api.SecretKey = "1234567890"

	url, err := url.Parse(urlString)
	if err != nil {
		t.Error("Could not parse urlstring")
	}

	resultUrl, err := SignAmazonUrl(url, api)
	if err != nil {
		t.Error("Signing failure: %v", err)
	}

	if signedUrl != resultUrl {
		t.Log(resultUrl, "\n")
		t.Error("Signed url does not match")
	}
}

func Test_SignAmazonUrl_1(t *testing.T) {

	urlString := `https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.10=1561711969&ASINList.ASIN.11=1561712027&ASINList.ASIN.12=2841051498&ASINList.ASIN.13=1561712043&ASINList.ASIN.14=1562413473&ASINList.ASIN.15=2729702806&ASINList.ASIN.16=2729702776&ASINList.ASIN.17=1561718939&ASINList.ASIN.18=2841090930&ASINList.ASIN.19=156171951X&ASINList.ASIN.2=1561712930&ASINList.ASIN.20=2729702032&ASINList.ASIN.3=1561713066&ASINList.ASIN.4=2729701737&ASINList.ASIN.5=1561711837&ASINList.ASIN.6=1561711845&ASINList.ASIN.7=1561711896&ASINList.ASIN.8=1561711810&ASINList.ASIN.9=1561712019&AWSAccessKeyId=AKIAJLUHOXLR5S2L6A6A&Action=GetLowestOfferListingsForASIN&MarketplaceId=APJ6JRA9NG5V4&SellerId=A2APQUVDBVWV7E&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2013-03-29T02%3A16%3A18Z`

	signedUrl := `https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.10=1561711969&ASINList.ASIN.11=1561712027&ASINList.ASIN.12=2841051498&ASINList.ASIN.13=1561712043&ASINList.ASIN.14=1562413473&ASINList.ASIN.15=2729702806&ASINList.ASIN.16=2729702776&ASINList.ASIN.17=1561718939&ASINList.ASIN.18=2841090930&ASINList.ASIN.19=156171951X&ASINList.ASIN.2=1561712930&ASINList.ASIN.20=2729702032&ASINList.ASIN.3=1561713066&ASINList.ASIN.4=2729701737&ASINList.ASIN.5=1561711837&ASINList.ASIN.6=1561711845&ASINList.ASIN.7=1561711896&ASINList.ASIN.8=1561711810&ASINList.ASIN.9=1561712019&AWSAccessKeyId=AKIAJLUHOXLR5S2L6A6A&Action=GetLowestOfferListingsForASIN&MarketplaceId=APJ6JRA9NG5V4&SellerId=A2APQUVDBVWV7E&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2013-03-29T02%3A16%3A18Z&Signature=rpDlBzpJ2t5nO3SLy66Y1oTAS9ZUhbH9kd639ed8g0w%3D`
	//signedUrl := `https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.2=1561712930&ASINList.ASIN.3=1561713066&ASINList.ASIN.4=2729701737&ASINList.ASIN.5=1561711837&ASINList.ASIN.6=1561711845&ASINList.ASIN.7=1561711896&ASINList.ASIN.8=1561711810&ASINList.ASIN.9=1561712019&ASINList.ASIN.10=1561711969&ASINList.ASIN.11=1561712027&ASINList.ASIN.12=2841051498&ASINList.ASIN.13=1561712043&ASINList.ASIN.14=1562413473&ASINList.ASIN.15=2729702806&ASINList.ASIN.16=2729702776&ASINList.ASIN.17=1561718939&ASINList.ASIN.18=2841090930&ASINList.ASIN.19=156171951X&ASINList.ASIN.20=2729702032&AWSAccessKeyId=AKIAJLUHOXLR5S2L6A6A&Action=GetLowestOfferListingsForASIN&MarketplaceId=APJ6JRA9NG5V4&SellerId=A2APQUVDBVWV7E&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2013-03-29T02%3A16%3A18Z&Signature=2O9DpwF6%2F0x6dX6QQLMCETP42NRkqqAaOzFDsZdIGs8%3D`

	var api AmazonMWSAPI

	api.SecretKey = "1234567890"

	url, err := url.Parse(urlString)
	if err != nil {
		t.Error("Could not parse urlstring")
	}

	resultUrl, err := SignAmazonUrl(url, api)
	if err != nil {
		t.Error("Signing failure: %v", err)
	}

	if signedUrl != resultUrl {
		t.Log(resultUrl, "\n")
		t.Error("Signed url does not match")
	}
}

func Test_GetLowestOfferListingForAsin(t *testing.T) {
}

func Test_GetLowestPricedOffersForASIN(t *testing.T) {
	item := LowestOfferListingsForASIN{
		MarketplaceId: "ATVPDKIKX0DER",
		ASIN:          "B00X4WHP5E",
		ItemCondition: "New",
	}

	t.Run("query params are generated correctly", func(t *testing.T) {
		params, _ := item.ToQueryParams()

		correctParams := map[string]string{
			"MarketplaceId": "ATVPDKIKX0DER",
			"ASIN":          "B00X4WHP5E",
			"ItemCondition": "New",
		}

		if !reflect.DeepEqual(params, correctParams) {
			t.Fatalf("Failed asserting that %v is equal to %v", params, correctParams)
		}
	})

	t.Run("ToQueryParams is being called", func(t *testing.T) {
		var api AmazonMWSAPI

		mock := &requestParamsMock{
			toQueryParamsHasBeenCalled: 0,
			setPrefixKeyHasBeenCalled:  0,
			toQueryParamsReturnValues: func() (map[string]string, error) {
				return nil, fmt.Errorf("Method has been called")
			},
		}

		_, err := api.GetLowestPricedOffersForASIN(mock)
		if err.Error() != "Method has been called" {
			t.Fatal("Method ToQueryParams() has not been called")
		}

		if mock.toQueryParamsHasBeenCalled != 1 {
			t.Fatal("Method ToQueryParams() has not been called")
		}
	})
}

func Test_GetMyFeesEstimateQuery(t *testing.T) {
	correctParams := map[string]string{
		"FeesEstimateRequestList.FeesEstimateRequest.1.MarketplaceId":                                 "ATVPDKIKX0DER",
		"FeesEstimateRequestList.FeesEstimateRequest.1.IsAmazonFulfilled":                             "true",
		"FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.ListingPrice.CurrencyCode": "USD",
		"FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.ListingPrice.Amount":       "10.11",
		"FeesEstimateRequestList.FeesEstimateRequest.1.Identifier":                                    "BOOKBOOK12",
		"FeesEstimateRequestList.FeesEstimateRequest.1.IdValue":                                       "BOOKBOOK12",
		"FeesEstimateRequestList.FeesEstimateRequest.1.IdType":                                        "ASIN",
	}

	correctKeys := []string{
		"FeesEstimateRequestList.FeesEstimateRequest.1.MarketplaceId",
		"FeesEstimateRequestList.FeesEstimateRequest.1.IsAmazonFulfilled",
		"FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.ListingPrice.CurrencyCode",
		"FeesEstimateRequestList.FeesEstimateRequest.1.PriceToEstimateFees.ListingPrice.AmountAmount",
		"FeesEstimateRequestList.FeesEstimateRequest.1.Identifier",
		"FeesEstimateRequestList.FeesEstimateRequest.1.IdValue",
		"FeesEstimateRequestList.FeesEstimateRequest.1.IdType",
	}

	t.Run("Query params with defaults settings", func(t *testing.T) {
		item := FeeEstimateRequest{
			IdValue:             "BOOKBOOK12",
			PriceToEstimateFees: 10.11,
		}

		item.setDefaults("ATVPDKIKX0DER", true)
		item.SetPrefixKey("FeesEstimateRequestList.FeesEstimateRequest.1")

		params, err := item.ToQueryParams()
		if err != nil {
			t.Fatalf("Unexpected error, %s", err.Error())
		}

		for _, key := range correctKeys {
			if params[key] != correctParams[key] {
				t.Fatalf("Expected '%s' key to have value of '%s', but got '%s'", key, correctParams[key], params[key])
			}
		}
	})

	t.Run("Query params With explicitly set values", func(t *testing.T) {
		item := FeeEstimateRequest{
			IdValue:             "BOOKBOOK12",
			PriceToEstimateFees: 10.11,
			Currency:            "USD",
			Identifier:          "BOOKBOOK12",
			IdType:              "ASIN",
			MarketplaceId:       "ATVPDKIKX0DER",
			IsAmazonFulfilled:   true,
		}

		item.SetPrefixKey("FeesEstimateRequestList.FeesEstimateRequest.1")

		params, err := item.ToQueryParams()
		if err != nil {
			t.Fatalf("Unexpected error, %s", err.Error())
		}

		for _, key := range correctKeys {
			if params[key] != correctParams[key] {
				t.Fatalf("Expected '%s' key to have value of '%s', but got '%s'", key, correctParams[key], params[key])
			}
		}
	})
}
