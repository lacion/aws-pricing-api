package schema

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type AWSServiceCatalog struct {
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	map[string]AWSServiceCatalog_Product
	Terms		map[string]map[string]AWSServiceCatalog_Term
}
type AWSServiceCatalog_Product struct {	Sku	string
	ProductFamily	string
	Attributes	AWSServiceCatalog_Product_Attributes
}
type AWSServiceCatalog_Product_Attributes struct {	Servicecode	string
	Location	string
	LocationType	string
	Usagetype	string
	Operation	string
	WithActiveUsers	string
}

type AWSServiceCatalog_Term struct {
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions AWSServiceCatalog_Term_PriceDimensions
	TermAttributes AWSServiceCatalog_Term_TermAttributes
}

type AWSServiceCatalog_Term_PriceDimensions struct {
	RateCode	string
	RateType	string
	Description	string
	BeginRange	string
	EndRange	string
	Unit	string
	PricePerUnit	AWSServiceCatalog_Term_PricePerUnit
	AppliesTo	[]interface{}
}

type AWSServiceCatalog_Term_PricePerUnit struct {
	USD	string
}

type AWSServiceCatalog_Term_TermAttributes struct {

}
func (a *AWSServiceCatalog) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AWSServiceCatalog/current/index.json"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, a)
	if err != nil {
		return err
	}

	return nil
}