// Generated ZEROPS sdk

package enum

type CardPaymentProviderEnum string

const (
	CardPaymentProviderEnumStripe = CardPaymentProviderEnum("STRIPE")
)

func NewCardPaymentProviderEnumFromString(value string) (out CardPaymentProviderEnum, err error) {
	return CardPaymentProviderEnum(value), nil
}

func (enum CardPaymentProviderEnum) String() string {
	return string(enum)
}

func (enum CardPaymentProviderEnum) Native() string {
	return string(enum)
}

func (enum CardPaymentProviderEnum) Values() []CardPaymentProviderEnum {
	return CardPaymentProviderEnumAll()
}

func (enum CardPaymentProviderEnum) PublicValues() []CardPaymentProviderEnum {
	return CardPaymentProviderEnumAllPublic()
}

func (enum CardPaymentProviderEnum) PrivateValues() []CardPaymentProviderEnum {
	return CardPaymentProviderEnumAllPrivate()
}

func (enum CardPaymentProviderEnum) DefaultValue() CardPaymentProviderEnum {
	return CardPaymentProviderEnumDefault()
}

func (enum CardPaymentProviderEnum) Is(values ...CardPaymentProviderEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func CardPaymentProviderEnumAllStrings() []string {
	return []string{
		string(CardPaymentProviderEnumStripe),
	}
}

func CardPaymentProviderEnumAll() []CardPaymentProviderEnum {
	return []CardPaymentProviderEnum{
		CardPaymentProviderEnumStripe,
	}
}

func CardPaymentProviderEnumAllPublic() []CardPaymentProviderEnum {
	return []CardPaymentProviderEnum{
		CardPaymentProviderEnumStripe,
	}
}

func CardPaymentProviderEnumAllPrivate() []CardPaymentProviderEnum {
	return []CardPaymentProviderEnum{}
}

func CardPaymentProviderEnumDefault() CardPaymentProviderEnum {
	return CardPaymentProviderEnumStripe
}

func (enum CardPaymentProviderEnum) IsStripe() bool {
	return enum.Is(CardPaymentProviderEnumStripe)
}
