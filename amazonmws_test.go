package amazonmws

import (
	//"fmt"
	"net/url"
	"testing"
)

func Test_SignAmazonUrl_1(t *testing.T) {

	urlString := `https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.10=1561711969&ASINList.ASIN.11=1561712027&ASINList.ASIN.12=2841051498&ASINList.ASIN.13=1561712043&ASINList.ASIN.14=1562413473&ASINList.ASIN.15=2729702806&ASINList.ASIN.16=2729702776&ASINList.ASIN.17=1561718939&ASINList.ASIN.18=2841090930&ASINList.ASIN.19=156171951X&ASINList.ASIN.2=1561712930&ASINList.ASIN.20=2729702032&ASINList.ASIN.3=1561713066&ASINList.ASIN.4=2729701737&ASINList.ASIN.5=1561711837&ASINList.ASIN.6=1561711845&ASINList.ASIN.7=1561711896&ASINList.ASIN.8=1561711810&ASINList.ASIN.9=1561712019&AWSAccessKeyId=AKIAJLUHOXLR5S2L6A6A&Action=GetLowestOfferListingsForASIN&MarketplaceId=APJ6JRA9NG5V4&SellerId=A2APQUVDBVWV7E&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2013-03-29T02%3A16%3A18Z`

	signedUrl := `https://mws-eu.amazonservices.com/Products/2011-10-01?ASINList.ASIN.1=1561706337&ASINList.ASIN.2=1561712930&ASINList.ASIN.3=1561713066&ASINList.ASIN.4=2729701737&ASINList.ASIN.5=1561711837&ASINList.ASIN.6=1561711845&ASINList.ASIN.7=1561711896&ASINList.ASIN.8=1561711810&ASINList.ASIN.9=1561712019&ASINList.ASIN.10=1561711969&ASINList.ASIN.11=1561712027&ASINList.ASIN.12=2841051498&ASINList.ASIN.13=1561712043&ASINList.ASIN.14=1562413473&ASINList.ASIN.15=2729702806&ASINList.ASIN.16=2729702776&ASINList.ASIN.17=1561718939&ASINList.ASIN.18=2841090930&ASINList.ASIN.19=156171951X&ASINList.ASIN.20=2729702032&AWSAccessKeyId=AKIAJLUHOXLR5S2L6A6A&Action=GetLowestOfferListingsForASIN&MarketplaceId=APJ6JRA9NG5V4&SellerId=A2APQUVDBVWV7E&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=2013-03-29T02%3A16%3A18Z&Signature=2O9DpwF6%2F0x6dX6QQLMCETP42NRkqqAaOzFDsZdIGs8%3D`

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

func TestGetMyFeesEstimateQuery(t *testing.T) {
	//item := FeeEstimateRequest{ IdValue: "BOOKBOOK12", PriceToEstimateFees: 10.11 }
	//
	//request := item.requestString(1, "US");
	//
	//result := item.toQuery(1, "US");
	//
	//fmt.Println(request);
	//fmt.Println(result);
}