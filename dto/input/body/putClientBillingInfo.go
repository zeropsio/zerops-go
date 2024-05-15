// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutClientBillingInfo)(nil)

type PutClientBillingInfo struct {
	VatNumber               types.StringNull   `json:"vatNumber"`
	CompanyNumber           types.StringNull   `json:"companyNumber"`
	CompanyName             types.String       `json:"companyName"`
	InvoiceAddressStreet    types.String       `json:"invoiceAddressStreet"`
	InvoiceAddressCity      types.String       `json:"invoiceAddressCity"`
	InvoiceAddressPostcode  types.String       `json:"invoiceAddressPostcode"`
	InvoiceAddressCountryId stringId.CountryId `json:"invoiceAddressCountryId"`
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
func (dto PutClientBillingInfo) GetInvoiceAddressStreet() types.String {
	return dto.InvoiceAddressStreet
}
func (dto PutClientBillingInfo) GetInvoiceAddressCity() types.String {
	return dto.InvoiceAddressCity
}
func (dto PutClientBillingInfo) GetInvoiceAddressPostcode() types.String {
	return dto.InvoiceAddressPostcode
}
func (dto PutClientBillingInfo) GetInvoiceAddressCountryId() stringId.CountryId {
	return dto.InvoiceAddressCountryId
}

func (dto *PutClientBillingInfo) UnmarshalJSON(b []byte) error {
	var aux = struct {
		VatNumber               types.StringNull
		CompanyNumber           types.StringNull
		CompanyName             *types.String
		InvoiceAddressStreet    *types.String
		InvoiceAddressCity      *types.String
		InvoiceAddressPostcode  *types.String
		InvoiceAddressCountryId *stringId.CountryId
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutClientBillingInfo", err)
	}
	var errorList validator.ErrorList
	if aux.CompanyName == nil {
		errorList = errorList.With(validator.NewError("companyName", "field is required"))
	}
	if aux.InvoiceAddressStreet == nil {
		errorList = errorList.With(validator.NewError("invoiceAddressStreet", "field is required"))
	}
	if aux.InvoiceAddressCity == nil {
		errorList = errorList.With(validator.NewError("invoiceAddressCity", "field is required"))
	}
	if aux.InvoiceAddressPostcode == nil {
		errorList = errorList.With(validator.NewError("invoiceAddressPostcode", "field is required"))
	}
	if aux.InvoiceAddressCountryId == nil {
		errorList = errorList.With(validator.NewError("invoiceAddressCountryId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.VatNumber = aux.VatNumber
	dto.CompanyNumber = aux.CompanyNumber
	dto.CompanyName = *aux.CompanyName
	dto.InvoiceAddressStreet = *aux.InvoiceAddressStreet
	dto.InvoiceAddressCity = *aux.InvoiceAddressCity
	dto.InvoiceAddressPostcode = *aux.InvoiceAddressPostcode
	dto.InvoiceAddressCountryId = *aux.InvoiceAddressCountryId

	return nil
}
