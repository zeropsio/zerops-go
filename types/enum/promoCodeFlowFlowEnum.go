// Generated ZEROPS sdk

package enum

type PromoCodeFlowFlowEnum string

const (
	PromoCodeFlowFlowEnumTopUp      = PromoCodeFlowFlowEnum("TOP_UP")
	PromoCodeFlowFlowEnumVerified   = PromoCodeFlowFlowEnum("VERIFIED")
	PromoCodeFlowFlowEnumStandalone = PromoCodeFlowFlowEnum("STANDALONE")
)

func NewPromoCodeFlowFlowEnumFromString(value string) (out PromoCodeFlowFlowEnum, err error) {
	return PromoCodeFlowFlowEnum(value), nil
}

func (enum PromoCodeFlowFlowEnum) String() string {
	return string(enum)
}

func (enum PromoCodeFlowFlowEnum) Native() string {
	return string(enum)
}

func (enum PromoCodeFlowFlowEnum) Values() []PromoCodeFlowFlowEnum {
	return PromoCodeFlowFlowEnumAll()
}

func (enum PromoCodeFlowFlowEnum) PublicValues() []PromoCodeFlowFlowEnum {
	return PromoCodeFlowFlowEnumAllPublic()
}

func (enum PromoCodeFlowFlowEnum) PrivateValues() []PromoCodeFlowFlowEnum {
	return PromoCodeFlowFlowEnumAllPrivate()
}

func (enum PromoCodeFlowFlowEnum) DefaultValue() PromoCodeFlowFlowEnum {
	return PromoCodeFlowFlowEnumDefault()
}

func (enum PromoCodeFlowFlowEnum) Is(values ...PromoCodeFlowFlowEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PromoCodeFlowFlowEnumAllStrings() []string {
	return []string{
		string(PromoCodeFlowFlowEnumTopUp), string(PromoCodeFlowFlowEnumVerified), string(PromoCodeFlowFlowEnumStandalone),
	}
}

func PromoCodeFlowFlowEnumAll() []PromoCodeFlowFlowEnum {
	return []PromoCodeFlowFlowEnum{
		PromoCodeFlowFlowEnumTopUp, PromoCodeFlowFlowEnumVerified, PromoCodeFlowFlowEnumStandalone,
	}
}

func PromoCodeFlowFlowEnumAllPublic() []PromoCodeFlowFlowEnum {
	return []PromoCodeFlowFlowEnum{
		PromoCodeFlowFlowEnumTopUp, PromoCodeFlowFlowEnumVerified, PromoCodeFlowFlowEnumStandalone,
	}
}

func PromoCodeFlowFlowEnumAllPrivate() []PromoCodeFlowFlowEnum {
	return []PromoCodeFlowFlowEnum{}
}

func PromoCodeFlowFlowEnumDefault() PromoCodeFlowFlowEnum {
	return ""
}

func (enum PromoCodeFlowFlowEnum) IsTopUp() bool {
	return enum.Is(PromoCodeFlowFlowEnumTopUp)
}

func (enum PromoCodeFlowFlowEnum) IsVerified() bool {
	return enum.Is(PromoCodeFlowFlowEnumVerified)
}

func (enum PromoCodeFlowFlowEnum) IsStandalone() bool {
	return enum.Is(PromoCodeFlowFlowEnumStandalone)
}
