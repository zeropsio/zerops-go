// Generated ZEROPS sdk

package enum

type BillingInfoVatModeEnum string

const (
	BillingInfoVatModeEnumDomesticVatPayer    = BillingInfoVatModeEnum("DOMESTIC_VAT_PAYER")
	BillingInfoVatModeEnumDomesticNotVatPayer = BillingInfoVatModeEnum("DOMESTIC_NOT_VAT_PAYER")
	BillingInfoVatModeEnumEuReverseCharge     = BillingInfoVatModeEnum("EU_REVERSE_CHARGE")
	BillingInfoVatModeEnumEuMoss              = BillingInfoVatModeEnum("EU_MOSS")
	BillingInfoVatModeEnumEuNotVatPayer       = BillingInfoVatModeEnum("EU_NOT_VAT_PAYER")
	BillingInfoVatModeEnumNotEu               = BillingInfoVatModeEnum("NOT_EU")
	BillingInfoVatModeEnumNotEuVerified       = BillingInfoVatModeEnum("NOT_EU_VERIFIED")
)

func NewBillingInfoVatModeEnumFromString(value string) (out BillingInfoVatModeEnum, err error) {
	return BillingInfoVatModeEnum(value), nil
}

func (enum BillingInfoVatModeEnum) String() string {
	return string(enum)
}

func (enum BillingInfoVatModeEnum) Native() string {
	return string(enum)
}

func (enum BillingInfoVatModeEnum) Values() []BillingInfoVatModeEnum {
	return BillingInfoVatModeEnumAll()
}

func (enum BillingInfoVatModeEnum) PublicValues() []BillingInfoVatModeEnum {
	return BillingInfoVatModeEnumAllPublic()
}

func (enum BillingInfoVatModeEnum) PrivateValues() []BillingInfoVatModeEnum {
	return BillingInfoVatModeEnumAllPrivate()
}

func (enum BillingInfoVatModeEnum) DefaultValue() BillingInfoVatModeEnum {
	return BillingInfoVatModeEnumDefault()
}

func (enum BillingInfoVatModeEnum) Is(values ...BillingInfoVatModeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func BillingInfoVatModeEnumAllStrings() []string {
	return []string{
		string(BillingInfoVatModeEnumDomesticVatPayer), string(BillingInfoVatModeEnumDomesticNotVatPayer), string(BillingInfoVatModeEnumEuReverseCharge), string(BillingInfoVatModeEnumEuMoss), string(BillingInfoVatModeEnumEuNotVatPayer), string(BillingInfoVatModeEnumNotEu), string(BillingInfoVatModeEnumNotEuVerified),
	}
}

func BillingInfoVatModeEnumAll() []BillingInfoVatModeEnum {
	return []BillingInfoVatModeEnum{
		BillingInfoVatModeEnumDomesticVatPayer, BillingInfoVatModeEnumDomesticNotVatPayer, BillingInfoVatModeEnumEuReverseCharge, BillingInfoVatModeEnumEuMoss, BillingInfoVatModeEnumEuNotVatPayer, BillingInfoVatModeEnumNotEu, BillingInfoVatModeEnumNotEuVerified,
	}
}

func BillingInfoVatModeEnumAllPublic() []BillingInfoVatModeEnum {
	return []BillingInfoVatModeEnum{
		BillingInfoVatModeEnumDomesticVatPayer, BillingInfoVatModeEnumDomesticNotVatPayer, BillingInfoVatModeEnumEuReverseCharge, BillingInfoVatModeEnumEuMoss, BillingInfoVatModeEnumEuNotVatPayer, BillingInfoVatModeEnumNotEu, BillingInfoVatModeEnumNotEuVerified,
	}
}

func BillingInfoVatModeEnumAllPrivate() []BillingInfoVatModeEnum {
	return []BillingInfoVatModeEnum{}
}

func BillingInfoVatModeEnumDefault() BillingInfoVatModeEnum {
	return ""
}

func (enum BillingInfoVatModeEnum) IsDomesticVatPayer() bool {
	return enum.Is(BillingInfoVatModeEnumDomesticVatPayer)
}

func (enum BillingInfoVatModeEnum) IsDomesticNotVatPayer() bool {
	return enum.Is(BillingInfoVatModeEnumDomesticNotVatPayer)
}

func (enum BillingInfoVatModeEnum) IsEuReverseCharge() bool {
	return enum.Is(BillingInfoVatModeEnumEuReverseCharge)
}

func (enum BillingInfoVatModeEnum) IsEuMoss() bool {
	return enum.Is(BillingInfoVatModeEnumEuMoss)
}

func (enum BillingInfoVatModeEnum) IsEuNotVatPayer() bool {
	return enum.Is(BillingInfoVatModeEnumEuNotVatPayer)
}

func (enum BillingInfoVatModeEnum) IsNotEu() bool {
	return enum.Is(BillingInfoVatModeEnumNotEu)
}

func (enum BillingInfoVatModeEnum) IsNotEuVerified() bool {
	return enum.Is(BillingInfoVatModeEnumNotEuVerified)
}
