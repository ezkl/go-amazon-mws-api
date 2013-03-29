# Amazon Marketplace Web Services (MWS) API

This Amazon MWS API client is based heavily on [go-amazon-product-api](ttps://github.com/DDRBoxman/go-amazon-product-api).

```go
package main

import (
       "fmt"
       "github.com/ezkl/go-amazon-mws-api"
)

func main() {
       var api amazonmws.AmazonMWSAPI

       api.AccessKey = ""
       api.SecretKey = ""
       api.Host = "mws.amazonservices.com"
       api.MarketplaceId = "ATVPDKIKX0DER"
       api.SellerId = ""

       asins = []string{"0195019199"}

       result,err := api.GetLowestOfferListingsForASIN(asins)

       if (err != nil) {
           fmt.Println(err)
       }

       fmt.Println(result)
}
```
