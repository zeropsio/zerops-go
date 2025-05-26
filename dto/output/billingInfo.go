// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type BillingInfo struct {
	VatNumber               types.StringNull            `json:"vatNumber"`
	CompanyNumber           types.StringNull            `json:"companyNumber"`
	CompanyName             types.String                `json:"companyName"`
	InvoiceAddressStreet    types.String                `json:"invoiceAddressStreet"`
	InvoiceAddressCity      types.String                `json:"invoiceAddressCity"`
	InvoiceAddressPostcode  types.String                `json:"invoiceAddressPostcode"`
	InvoiceAddressCountryId stringId.CountryId          `json:"invoiceAddressCountryId"`
	VatCountryId            stringId.CountryId          `json:"vatCountryId"`
	VatMode                 enum.BillingInfoVatModeEnum `json:"vatMode"`
	VatRate                 types.Float                 `json:"vatRate"`
	VatVerified             types.Bool                  `json:"vatVerified"`
}
