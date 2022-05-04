// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type BillingInfo struct {
	VatPayer                types.Bool       `json:"vatPayer"`
	VatNumber               types.StringNull `json:"vatNumber"`
	CompanyNumber           types.StringNull `json:"companyNumber"`
	CompanyName             types.String     `json:"companyName"`
	InvoiceAddressStreet    types.StringNull `json:"invoiceAddressStreet"`
	InvoiceAddressCity      types.StringNull `json:"invoiceAddressCity"`
	InvoiceAddressPostcode  types.StringNull `json:"invoiceAddressPostcode"`
	InvoiceAddressCountryId types.StringNull `json:"invoiceAddressCountryId"`
}
