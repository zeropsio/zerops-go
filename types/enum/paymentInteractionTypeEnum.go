// Generated ZEROPS sdk

package enum

type PaymentInteractionTypeEnum string

const (
	PaymentInteractionTypeEnumOnsession  = PaymentInteractionTypeEnum("ONSESSION")
	PaymentInteractionTypeEnumOffsession = PaymentInteractionTypeEnum("OFFSESSION")
)

func NewPaymentInteractionTypeEnumFromString(value string) (out PaymentInteractionTypeEnum, err error) {
	return PaymentInteractionTypeEnum(value), nil
}

func (enum PaymentInteractionTypeEnum) String() string {
	return string(enum)
}

func (enum PaymentInteractionTypeEnum) Native() string {
	return string(enum)
}

func (enum PaymentInteractionTypeEnum) Values() []PaymentInteractionTypeEnum {
	return PaymentInteractionTypeEnumAll()
}

func (enum PaymentInteractionTypeEnum) PublicValues() []PaymentInteractionTypeEnum {
	return PaymentInteractionTypeEnumAllPublic()
}

func (enum PaymentInteractionTypeEnum) PrivateValues() []PaymentInteractionTypeEnum {
	return PaymentInteractionTypeEnumAllPrivate()
}

func (enum PaymentInteractionTypeEnum) DefaultValue() PaymentInteractionTypeEnum {
	return PaymentInteractionTypeEnumDefault()
}

func (enum PaymentInteractionTypeEnum) Is(values ...PaymentInteractionTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PaymentInteractionTypeEnumAllStrings() []string {
	return []string{
		string(PaymentInteractionTypeEnumOnsession), string(PaymentInteractionTypeEnumOffsession),
	}
}

func PaymentInteractionTypeEnumAll() []PaymentInteractionTypeEnum {
	return []PaymentInteractionTypeEnum{
		PaymentInteractionTypeEnumOnsession, PaymentInteractionTypeEnumOffsession,
	}
}

func PaymentInteractionTypeEnumAllPublic() []PaymentInteractionTypeEnum {
	return []PaymentInteractionTypeEnum{
		PaymentInteractionTypeEnumOnsession, PaymentInteractionTypeEnumOffsession,
	}
}

func PaymentInteractionTypeEnumAllPrivate() []PaymentInteractionTypeEnum {
	return []PaymentInteractionTypeEnum{}
}

func PaymentInteractionTypeEnumDefault() PaymentInteractionTypeEnum {
	return PaymentInteractionTypeEnumOnsession
}

func (enum PaymentInteractionTypeEnum) IsOnsession() bool {
	return enum.Is(PaymentInteractionTypeEnumOnsession)
}

func (enum PaymentInteractionTypeEnum) IsOffsession() bool {
	return enum.Is(PaymentInteractionTypeEnumOffsession)
}
