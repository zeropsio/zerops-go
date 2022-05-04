// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type PaymentSource struct {
	Id                 types.String `json:"id"`
	Object             types.String `json:"object"`
	AddressCity        types.String `json:"addressCity"`
	AddressCountry     types.String `json:"addressCountry"`
	AddressLine1       types.String `json:"addressLine1"`
	AddressLine1Check  types.String `json:"addressLine1Check"`
	AddressLine2       types.String `json:"addressLine2"`
	AddressState       types.String `json:"addressState"`
	AddressZip         types.String `json:"addressZip"`
	AddressZipCheck    types.String `json:"addressZipCheck"`
	Brand              types.String `json:"brand"`
	Country            types.String `json:"country"`
	CvcCheck           types.String `json:"cvcCheck"`
	ExpMonth           types.String `json:"expMonth"`
	ExpYear            types.String `json:"expYear"`
	Funding            types.String `json:"funding"`
	Last4              types.String `json:"last4"`
	Name               types.String `json:"name"`
	TokenizationMethod types.String `json:"tokenizationMethod"`
	Metadata           types.Map    `json:"metadata"`
}
