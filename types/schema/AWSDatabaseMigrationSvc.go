package schema

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/jinzhu/gorm"
)

type AWSDatabaseMigrationSvc struct {
	gorm.Model
	FormatVersion	string
	Disclaimer	string
	OfferCode	string
	Version		string
	PublicationDate	string
	Products	map[string]AWSDatabaseMigrationSvc_Product
	Terms		map[string]map[string]map[string]AWSDatabaseMigrationSvc_Term
}
type AWSDatabaseMigrationSvc_Product struct {	Sku	string
	ProductFamily	string
	Attributes	AWSDatabaseMigrationSvc_Product_Attributes
}
type AWSDatabaseMigrationSvc_Product_Attributes struct {	Servicecode	string
	LocationType	string
	DedicatedEbsThroughput	string
	Operation	string
	AvailabilityZone	string
	Location	string
	CurrentGeneration	string
	PhysicalProcessor	string
	ClockSpeed	string
	Storage	string
	Usagetype	string
	EnhancedNetworkingSupported	string
	Vcpu	string
	NetworkPerformance	string
	ProcessorFeatures	string
	InstanceType	string
	InstanceFamily	string
	Memory	string
	ProcessorArchitecture	string
}

type AWSDatabaseMigrationSvc_Term struct {
	OfferTermCode string
	Sku	string
	EffectiveDate string
	PriceDimensions map[string]AWSDatabaseMigrationSvc_Term_PriceDimensions
	TermAttributes AWSDatabaseMigrationSvc_Term_TermAttributes
}

type AWSDatabaseMigrationSvc_Term_PriceDimensions struct {
	RateCode	string
	RateType	string
	Description	string
	BeginRange	string
	EndRange	string
	Unit	string
	PricePerUnit	AWSDatabaseMigrationSvc_Term_PricePerUnit
	AppliesTo	[]interface{}
}

type AWSDatabaseMigrationSvc_Term_PricePerUnit struct {
	USD	string
}

type AWSDatabaseMigrationSvc_Term_TermAttributes struct {

}
func (a AWSDatabaseMigrationSvc) QueryProducts(q func(product AWSDatabaseMigrationSvc_Product) bool) []AWSDatabaseMigrationSvc_Product{
	ret := []AWSDatabaseMigrationSvc_Product{}
	for _, v := range a.Products {
		if q(v) {
			ret = append(ret, v)
		}
	}

	return ret
}
func (a AWSDatabaseMigrationSvc) QueryTerms(t string, q func(product AWSDatabaseMigrationSvc_Term) bool) []AWSDatabaseMigrationSvc_Term{
	ret := []AWSDatabaseMigrationSvc_Term{}
	for _, v := range a.Terms[t] {
		for _, val := range v {
			if q(val) {
				ret = append(ret, val)
			}
		}
	}

	return ret
}
func (a *AWSDatabaseMigrationSvc) Refresh() error {
	var url = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AWSDatabaseMigrationSvc/current/index.json"
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