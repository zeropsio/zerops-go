// Generated ZEROPS sdk

package enum

type InvoiceStatusEnum string

const (
	InvoiceStatusEnumPaid = InvoiceStatusEnum("PAID")
)

func NewInvoiceStatusEnumFromString(value string) (out InvoiceStatusEnum, err error) {
	return InvoiceStatusEnum(value), nil
}

func (enum InvoiceStatusEnum) String() string {
	return string(enum)
}

func (enum InvoiceStatusEnum) Native() string {
	return string(enum)
}

func (enum InvoiceStatusEnum) Values() []InvoiceStatusEnum {
	return InvoiceStatusEnumAll()
}

func (enum InvoiceStatusEnum) PublicValues() []InvoiceStatusEnum {
	return InvoiceStatusEnumAllPublic()
}

func (enum InvoiceStatusEnum) PrivateValues() []InvoiceStatusEnum {
	return InvoiceStatusEnumAllPrivate()
}

func (enum InvoiceStatusEnum) DefaultValue() InvoiceStatusEnum {
	return InvoiceStatusEnumDefault()
}

func (enum InvoiceStatusEnum) Is(values ...InvoiceStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func InvoiceStatusEnumAllStrings() []string {
	return []string{
		string(InvoiceStatusEnumPaid),
	}
}

func InvoiceStatusEnumAll() []InvoiceStatusEnum {
	return []InvoiceStatusEnum{
		InvoiceStatusEnumPaid,
	}
}

func InvoiceStatusEnumAllPublic() []InvoiceStatusEnum {
	return []InvoiceStatusEnum{
		InvoiceStatusEnumPaid,
	}
}

func InvoiceStatusEnumAllPrivate() []InvoiceStatusEnum {
	return []InvoiceStatusEnum{}
}

func InvoiceStatusEnumDefault() InvoiceStatusEnum {
	return ""
}

func (enum InvoiceStatusEnum) IsPaid() bool {
	return enum.Is(InvoiceStatusEnumPaid)
}
