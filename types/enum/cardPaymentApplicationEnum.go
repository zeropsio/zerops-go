// Generated ZEROPS sdk

package enum

type CardPaymentApplicationEnum string

const (
	CardPaymentApplicationEnumZerops = CardPaymentApplicationEnum("ZEROPS")
)

func NewCardPaymentApplicationEnumFromString(value string) (out CardPaymentApplicationEnum, err error) {
	return CardPaymentApplicationEnum(value), nil
}

func (enum CardPaymentApplicationEnum) String() string {
	return string(enum)
}

func (enum CardPaymentApplicationEnum) Native() string {
	return string(enum)
}

func (enum CardPaymentApplicationEnum) Values() []CardPaymentApplicationEnum {
	return CardPaymentApplicationEnumAll()
}

func (enum CardPaymentApplicationEnum) PublicValues() []CardPaymentApplicationEnum {
	return CardPaymentApplicationEnumAllPublic()
}

func (enum CardPaymentApplicationEnum) PrivateValues() []CardPaymentApplicationEnum {
	return CardPaymentApplicationEnumAllPrivate()
}

func (enum CardPaymentApplicationEnum) DefaultValue() CardPaymentApplicationEnum {
	return CardPaymentApplicationEnumDefault()
}

func (enum CardPaymentApplicationEnum) Is(values ...CardPaymentApplicationEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func CardPaymentApplicationEnumAllStrings() []string {
	return []string{
		string(CardPaymentApplicationEnumZerops),
	}
}

func CardPaymentApplicationEnumAll() []CardPaymentApplicationEnum {
	return []CardPaymentApplicationEnum{
		CardPaymentApplicationEnumZerops,
	}
}

func CardPaymentApplicationEnumAllPublic() []CardPaymentApplicationEnum {
	return []CardPaymentApplicationEnum{
		CardPaymentApplicationEnumZerops,
	}
}

func CardPaymentApplicationEnumAllPrivate() []CardPaymentApplicationEnum {
	return []CardPaymentApplicationEnum{}
}

func CardPaymentApplicationEnumDefault() CardPaymentApplicationEnum {
	return CardPaymentApplicationEnumZerops
}

func (enum CardPaymentApplicationEnum) IsZerops() bool {
	return enum.Is(CardPaymentApplicationEnumZerops)
}
