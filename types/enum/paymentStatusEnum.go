// Generated ZEROPS sdk

package enum

type PaymentStatusEnum string

const (
	PaymentStatusEnumNew                  = PaymentStatusEnum("NEW")
	PaymentStatusEnumProcessing           = PaymentStatusEnum("PROCESSING")
	PaymentStatusEnumProcessingFailed     = PaymentStatusEnum("PROCESSING_FAILED")
	PaymentStatusEnumPaymentDeclined      = PaymentStatusEnum("PAYMENT_DECLINED")
	PaymentStatusEnumPaymentDeclineFailed = PaymentStatusEnum("PAYMENT_DECLINE_FAILED")
	PaymentStatusEnumPaymentSuccess       = PaymentStatusEnum("PAYMENT_SUCCESS")
	PaymentStatusEnumBalanceUpdated       = PaymentStatusEnum("BALANCE_UPDATED")
	PaymentStatusEnumBalanceUpdateFailed  = PaymentStatusEnum("BALANCE_UPDATE_FAILED")
	PaymentStatusEnumInvoiced             = PaymentStatusEnum("INVOICED")
	PaymentStatusEnumInvoicingFailed      = PaymentStatusEnum("INVOICING_FAILED")
	PaymentStatusEnumFinished             = PaymentStatusEnum("FINISHED")
	PaymentStatusEnumCanceled             = PaymentStatusEnum("CANCELED")
	PaymentStatusEnumDeclined             = PaymentStatusEnum("DECLINED")
)

func NewPaymentStatusEnumFromString(value string) (out PaymentStatusEnum, err error) {
	return PaymentStatusEnum(value), nil
}

func (enum PaymentStatusEnum) String() string {
	return string(enum)
}

func (enum PaymentStatusEnum) Native() string {
	return string(enum)
}

func (enum PaymentStatusEnum) Values() []PaymentStatusEnum {
	return PaymentStatusEnumAll()
}

func (enum PaymentStatusEnum) PublicValues() []PaymentStatusEnum {
	return PaymentStatusEnumAllPublic()
}

func (enum PaymentStatusEnum) PrivateValues() []PaymentStatusEnum {
	return PaymentStatusEnumAllPrivate()
}

func (enum PaymentStatusEnum) DefaultValue() PaymentStatusEnum {
	return PaymentStatusEnumDefault()
}

func (enum PaymentStatusEnum) Is(values ...PaymentStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PaymentStatusEnumAllStrings() []string {
	return []string{
		string(PaymentStatusEnumNew), string(PaymentStatusEnumProcessing), string(PaymentStatusEnumProcessingFailed), string(PaymentStatusEnumPaymentDeclined), string(PaymentStatusEnumPaymentDeclineFailed), string(PaymentStatusEnumPaymentSuccess), string(PaymentStatusEnumBalanceUpdated), string(PaymentStatusEnumBalanceUpdateFailed), string(PaymentStatusEnumInvoiced), string(PaymentStatusEnumInvoicingFailed), string(PaymentStatusEnumFinished), string(PaymentStatusEnumCanceled), string(PaymentStatusEnumDeclined),
	}
}

func PaymentStatusEnumAll() []PaymentStatusEnum {
	return []PaymentStatusEnum{
		PaymentStatusEnumNew, PaymentStatusEnumProcessing, PaymentStatusEnumProcessingFailed, PaymentStatusEnumPaymentDeclined, PaymentStatusEnumPaymentDeclineFailed, PaymentStatusEnumPaymentSuccess, PaymentStatusEnumBalanceUpdated, PaymentStatusEnumBalanceUpdateFailed, PaymentStatusEnumInvoiced, PaymentStatusEnumInvoicingFailed, PaymentStatusEnumFinished, PaymentStatusEnumCanceled, PaymentStatusEnumDeclined,
	}
}

func PaymentStatusEnumAllPublic() []PaymentStatusEnum {
	return []PaymentStatusEnum{
		PaymentStatusEnumNew, PaymentStatusEnumProcessing, PaymentStatusEnumProcessingFailed, PaymentStatusEnumPaymentDeclined, PaymentStatusEnumPaymentDeclineFailed, PaymentStatusEnumPaymentSuccess, PaymentStatusEnumBalanceUpdated, PaymentStatusEnumBalanceUpdateFailed, PaymentStatusEnumInvoiced, PaymentStatusEnumInvoicingFailed, PaymentStatusEnumFinished, PaymentStatusEnumCanceled, PaymentStatusEnumDeclined,
	}
}

func PaymentStatusEnumAllPrivate() []PaymentStatusEnum {
	return []PaymentStatusEnum{}
}

func PaymentStatusEnumDefault() PaymentStatusEnum {
	return PaymentStatusEnumNew
}

func (enum PaymentStatusEnum) IsNew() bool {
	return enum.Is(PaymentStatusEnumNew)
}

func (enum PaymentStatusEnum) IsProcessing() bool {
	return enum.Is(PaymentStatusEnumProcessing)
}

func (enum PaymentStatusEnum) IsProcessingFailed() bool {
	return enum.Is(PaymentStatusEnumProcessingFailed)
}

func (enum PaymentStatusEnum) IsPaymentDeclined() bool {
	return enum.Is(PaymentStatusEnumPaymentDeclined)
}

func (enum PaymentStatusEnum) IsPaymentDeclineFailed() bool {
	return enum.Is(PaymentStatusEnumPaymentDeclineFailed)
}

func (enum PaymentStatusEnum) IsPaymentSuccess() bool {
	return enum.Is(PaymentStatusEnumPaymentSuccess)
}

func (enum PaymentStatusEnum) IsBalanceUpdated() bool {
	return enum.Is(PaymentStatusEnumBalanceUpdated)
}

func (enum PaymentStatusEnum) IsBalanceUpdateFailed() bool {
	return enum.Is(PaymentStatusEnumBalanceUpdateFailed)
}

func (enum PaymentStatusEnum) IsInvoiced() bool {
	return enum.Is(PaymentStatusEnumInvoiced)
}

func (enum PaymentStatusEnum) IsInvoicingFailed() bool {
	return enum.Is(PaymentStatusEnumInvoicingFailed)
}

func (enum PaymentStatusEnum) IsFinished() bool {
	return enum.Is(PaymentStatusEnumFinished)
}

func (enum PaymentStatusEnum) IsCanceled() bool {
	return enum.Is(PaymentStatusEnumCanceled)
}

func (enum PaymentStatusEnum) IsDeclined() bool {
	return enum.Is(PaymentStatusEnumDeclined)
}
