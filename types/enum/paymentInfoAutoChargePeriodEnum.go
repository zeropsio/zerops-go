// Generated ZEROPS sdk

package enum

type PaymentInfoAutoChargePeriodEnum string

const (
	PaymentInfoAutoChargePeriodEnumWeek  = PaymentInfoAutoChargePeriodEnum("WEEK")
	PaymentInfoAutoChargePeriodEnumMonth = PaymentInfoAutoChargePeriodEnum("MONTH")
)

func NewPaymentInfoAutoChargePeriodEnumFromString(value string) (out PaymentInfoAutoChargePeriodEnum, err error) {
	return PaymentInfoAutoChargePeriodEnum(value), nil
}

func (enum PaymentInfoAutoChargePeriodEnum) String() string {
	return string(enum)
}

func (enum PaymentInfoAutoChargePeriodEnum) Native() string {
	return string(enum)
}

func (enum PaymentInfoAutoChargePeriodEnum) Values() []PaymentInfoAutoChargePeriodEnum {
	return PaymentInfoAutoChargePeriodEnumAll()
}

func (enum PaymentInfoAutoChargePeriodEnum) PublicValues() []PaymentInfoAutoChargePeriodEnum {
	return PaymentInfoAutoChargePeriodEnumAllPublic()
}

func (enum PaymentInfoAutoChargePeriodEnum) PrivateValues() []PaymentInfoAutoChargePeriodEnum {
	return PaymentInfoAutoChargePeriodEnumAllPrivate()
}

func (enum PaymentInfoAutoChargePeriodEnum) DefaultValue() PaymentInfoAutoChargePeriodEnum {
	return PaymentInfoAutoChargePeriodEnumDefault()
}

func (enum PaymentInfoAutoChargePeriodEnum) Is(values ...PaymentInfoAutoChargePeriodEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PaymentInfoAutoChargePeriodEnumAllStrings() []string {
	return []string{
		string(PaymentInfoAutoChargePeriodEnumWeek), string(PaymentInfoAutoChargePeriodEnumMonth),
	}
}

func PaymentInfoAutoChargePeriodEnumAll() []PaymentInfoAutoChargePeriodEnum {
	return []PaymentInfoAutoChargePeriodEnum{
		PaymentInfoAutoChargePeriodEnumWeek, PaymentInfoAutoChargePeriodEnumMonth,
	}
}

func PaymentInfoAutoChargePeriodEnumAllPublic() []PaymentInfoAutoChargePeriodEnum {
	return []PaymentInfoAutoChargePeriodEnum{
		PaymentInfoAutoChargePeriodEnumWeek, PaymentInfoAutoChargePeriodEnumMonth,
	}
}

func PaymentInfoAutoChargePeriodEnumAllPrivate() []PaymentInfoAutoChargePeriodEnum {
	return []PaymentInfoAutoChargePeriodEnum{}
}

func PaymentInfoAutoChargePeriodEnumDefault() PaymentInfoAutoChargePeriodEnum {
	return PaymentInfoAutoChargePeriodEnumWeek
}

func (enum PaymentInfoAutoChargePeriodEnum) IsWeek() bool {
	return enum.Is(PaymentInfoAutoChargePeriodEnumWeek)
}

func (enum PaymentInfoAutoChargePeriodEnum) IsMonth() bool {
	return enum.Is(PaymentInfoAutoChargePeriodEnumMonth)
}
