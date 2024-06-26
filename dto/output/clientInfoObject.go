// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ClientInfoObject struct {
	Id               types.String             `json:"id"`
	CompanyName      types.String             `json:"companyName"`
	CompanyNumber    types.StringNull         `json:"companyNumber"`
	VatNumber        types.StringNull         `json:"vatNumber"`
	Created          types.DateTime           `json:"created"`
	LastUpdate       types.DateTime           `json:"lastUpdate"`
	LastPublicUpdate types.DateTime           `json:"lastPublicUpdate"`
	Status           types.String             `json:"status"`
	VatCountry       types.String             `json:"vatCountry"`
	VatPayer         types.Bool               `json:"vatPayer"`
	VatRate          types.Float              `json:"vatRate"`
	VatMode          types.StringNull         `json:"vatMode"`
	Currency         ClientInfoObjectCurrency `json:"currency"`
	Language         ClientInfoLanguageObject `json:"language"`
}
