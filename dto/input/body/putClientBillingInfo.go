// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutClientBillingInfo)(nil)

type PutClientBillingInfo struct {
	VatPayer                types.Bool       `json:"vatPayer"`
	VatNumber               types.StringNull `json:"vatNumber"`
	CompanyNumber           types.StringNull `json:"companyNumber"`
	CompanyName             types.String     `json:"companyName"`
	InvoiceAddressStreet    types.StringNull `json:"invoiceAddressStreet"`
	InvoiceAddressCity      types.StringNull `json:"invoiceAddressCity"`
	InvoiceAddressPostcode  types.StringNull `json:"invoiceAddressPostcode"`
	InvoiceAddressCountryId types.String     `json:"invoiceAddressCountryId"`
}

func (dto PutClientBillingInfo) GetVatPayer() types.Bool {
	return dto.VatPayer
}
func (dto PutClientBillingInfo) GetVatNumber() types.StringNull {
	return dto.VatNumber
}
func (dto PutClientBillingInfo) GetCompanyNumber() types.StringNull {
	return dto.CompanyNumber
}
func (dto PutClientBillingInfo) GetCompanyName() types.String {
	return dto.CompanyName
}
func (dto PutClientBillingInfo) GetInvoiceAddressStreet() types.StringNull {
	return dto.InvoiceAddressStreet
}
func (dto PutClientBillingInfo) GetInvoiceAddressCity() types.StringNull {
	return dto.InvoiceAddressCity
}
func (dto PutClientBillingInfo) GetInvoiceAddressPostcode() types.StringNull {
	return dto.InvoiceAddressPostcode
}
func (dto PutClientBillingInfo) GetInvoiceAddressCountryId() types.String {
	return dto.InvoiceAddressCountryId
}

func (dto *PutClientBillingInfo) UnmarshalJSON(b []byte) error {
	var aux = struct {
		VatPayer                *types.Bool
		VatNumber               types.StringNull
		CompanyNumber           types.StringNull
		CompanyName             *types.String
		InvoiceAddressStreet    types.StringNull
		InvoiceAddressCity      types.StringNull
		InvoiceAddressPostcode  types.StringNull
		InvoiceAddressCountryId *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutClientBillingInfo", err)
	}
	var errorList validator.ErrorList
	if aux.VatPayer == nil {
		errorList = errorList.With(validator.NewError("vatPayer", "field is required"))
	}
	if aux.CompanyName == nil {
		errorList = errorList.With(validator.NewError("companyName", "field is required"))
	}
	if aux.InvoiceAddressCountryId == nil {
		errorList = errorList.With(validator.NewError("invoiceAddressCountryId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.VatPayer = *aux.VatPayer
	dto.VatNumber = aux.VatNumber
	dto.CompanyNumber = aux.CompanyNumber
	dto.CompanyName = *aux.CompanyName
	dto.InvoiceAddressStreet = aux.InvoiceAddressStreet
	dto.InvoiceAddressCity = aux.InvoiceAddressCity
	dto.InvoiceAddressPostcode = aux.InvoiceAddressPostcode
	dto.InvoiceAddressCountryId = *aux.InvoiceAddressCountryId

	return nil
}
