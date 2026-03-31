// Generated ZEROPS sdk

package enum

type InvoiceTypeEnum string

const (
	InvoiceTypeEnumInvoice    = InvoiceTypeEnum("INVOICE")
	InvoiceTypeEnumReceipt    = InvoiceTypeEnum("RECEIPT")
	InvoiceTypeEnumCreditNote = InvoiceTypeEnum("CREDIT_NOTE")
)

func NewInvoiceTypeEnumFromString(value string) (out InvoiceTypeEnum, err error) {
	return InvoiceTypeEnum(value), nil
}

func (enum InvoiceTypeEnum) String() string {
	return string(enum)
}

func (enum InvoiceTypeEnum) Native() string {
	return string(enum)
}

func (enum InvoiceTypeEnum) Values() []InvoiceTypeEnum {
	return InvoiceTypeEnumAll()
}

func (enum InvoiceTypeEnum) PublicValues() []InvoiceTypeEnum {
	return InvoiceTypeEnumAllPublic()
}

func (enum InvoiceTypeEnum) PrivateValues() []InvoiceTypeEnum {
	return InvoiceTypeEnumAllPrivate()
}

func (enum InvoiceTypeEnum) DefaultValue() InvoiceTypeEnum {
	return InvoiceTypeEnumDefault()
}

func (enum InvoiceTypeEnum) Is(values ...InvoiceTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func InvoiceTypeEnumAllStrings() []string {
	return []string{
		string(InvoiceTypeEnumInvoice), string(InvoiceTypeEnumReceipt), string(InvoiceTypeEnumCreditNote),
	}
}

func InvoiceTypeEnumAll() []InvoiceTypeEnum {
	return []InvoiceTypeEnum{
		InvoiceTypeEnumInvoice, InvoiceTypeEnumReceipt, InvoiceTypeEnumCreditNote,
	}
}

func InvoiceTypeEnumAllPublic() []InvoiceTypeEnum {
	return []InvoiceTypeEnum{
		InvoiceTypeEnumInvoice, InvoiceTypeEnumReceipt, InvoiceTypeEnumCreditNote,
	}
}

func InvoiceTypeEnumAllPrivate() []InvoiceTypeEnum {
	return []InvoiceTypeEnum{}
}

func InvoiceTypeEnumDefault() InvoiceTypeEnum {
	return ""
}

func (enum InvoiceTypeEnum) IsInvoice() bool {
	return enum.Is(InvoiceTypeEnumInvoice)
}

func (enum InvoiceTypeEnum) IsReceipt() bool {
	return enum.Is(InvoiceTypeEnumReceipt)
}

func (enum InvoiceTypeEnum) IsCreditNote() bool {
	return enum.Is(InvoiceTypeEnumCreditNote)
}
