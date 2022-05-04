// Generated ZEROPS sdk

package enum

type CurrencyIdEnum string

const (
	CurrencyIdEnumUsd = CurrencyIdEnum("USD")
)

func NewCurrencyIdEnumFromString(value string) (out CurrencyIdEnum, err error) {
	return CurrencyIdEnum(value), nil
}

func (enum CurrencyIdEnum) String() string {
	return string(enum)
}

func (enum CurrencyIdEnum) Native() string {
	return string(enum)
}

func (enum CurrencyIdEnum) Values() []CurrencyIdEnum {
	return CurrencyIdEnumAll()
}

func (enum CurrencyIdEnum) PublicValues() []CurrencyIdEnum {
	return CurrencyIdEnumAllPublic()
}

func (enum CurrencyIdEnum) PrivateValues() []CurrencyIdEnum {
	return CurrencyIdEnumAllPrivate()
}

func (enum CurrencyIdEnum) DefaultValue() CurrencyIdEnum {
	return CurrencyIdEnumDefault()
}

func (enum CurrencyIdEnum) Is(values ...CurrencyIdEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func CurrencyIdEnumAllStrings() []string {
	return []string{
		string(CurrencyIdEnumUsd),
	}
}

func CurrencyIdEnumAll() []CurrencyIdEnum {
	return []CurrencyIdEnum{
		CurrencyIdEnumUsd,
	}
}

func CurrencyIdEnumAllPublic() []CurrencyIdEnum {
	return []CurrencyIdEnum{
		CurrencyIdEnumUsd,
	}
}

func CurrencyIdEnumAllPrivate() []CurrencyIdEnum {
	return []CurrencyIdEnum{}
}

func CurrencyIdEnumDefault() CurrencyIdEnum {
	return CurrencyIdEnumUsd
}

func (enum CurrencyIdEnum) IsUsd() bool {
	return enum.Is(CurrencyIdEnumUsd)
}
