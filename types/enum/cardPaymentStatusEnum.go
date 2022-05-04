// Generated ZEROPS sdk

package enum

type CardPaymentStatusEnum string

const (
	CardPaymentStatusEnumRequested            = CardPaymentStatusEnum("REQUESTED")
	CardPaymentStatusEnumProcessing           = CardPaymentStatusEnum("PROCESSING")
	CardPaymentStatusEnumPaymentSuccess       = CardPaymentStatusEnum("PAYMENT_SUCCESS")
	CardPaymentStatusEnumPaymentDeclined      = CardPaymentStatusEnum("PAYMENT_DECLINED")
	CardPaymentStatusEnumBalanceUpdating      = CardPaymentStatusEnum("BALANCE_UPDATING")
	CardPaymentStatusEnumBalanceUpdated       = CardPaymentStatusEnum("BALANCE_UPDATED")
	CardPaymentStatusEnumInvoicing            = CardPaymentStatusEnum("INVOICING")
	CardPaymentStatusEnumInvoiced             = CardPaymentStatusEnum("INVOICED")
	CardPaymentStatusEnumPaymentFailed        = CardPaymentStatusEnum("PAYMENT_FAILED")
	CardPaymentStatusEnumBalanceUpdateFailed  = CardPaymentStatusEnum("BALANCE_UPDATE_FAILED")
	CardPaymentStatusEnumInvoicingFailed      = CardPaymentStatusEnum("INVOICING_FAILED")
	CardPaymentStatusEnumFinished             = CardPaymentStatusEnum("FINISHED")
	CardPaymentStatusEnumPaymentDeclineFailed = CardPaymentStatusEnum("PAYMENT_DECLINE_FAILED")
	CardPaymentStatusEnumFailed               = CardPaymentStatusEnum("FAILED")
)

func NewCardPaymentStatusEnumFromString(value string) (out CardPaymentStatusEnum, err error) {
	return CardPaymentStatusEnum(value), nil
}

func (enum CardPaymentStatusEnum) String() string {
	return string(enum)
}

func (enum CardPaymentStatusEnum) Native() string {
	return string(enum)
}

func (enum CardPaymentStatusEnum) Values() []CardPaymentStatusEnum {
	return CardPaymentStatusEnumAll()
}

func (enum CardPaymentStatusEnum) PublicValues() []CardPaymentStatusEnum {
	return CardPaymentStatusEnumAllPublic()
}

func (enum CardPaymentStatusEnum) PrivateValues() []CardPaymentStatusEnum {
	return CardPaymentStatusEnumAllPrivate()
}

func (enum CardPaymentStatusEnum) DefaultValue() CardPaymentStatusEnum {
	return CardPaymentStatusEnumDefault()
}

func (enum CardPaymentStatusEnum) Is(values ...CardPaymentStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func CardPaymentStatusEnumAllStrings() []string {
	return []string{
		string(CardPaymentStatusEnumRequested), string(CardPaymentStatusEnumProcessing), string(CardPaymentStatusEnumPaymentSuccess), string(CardPaymentStatusEnumPaymentDeclined), string(CardPaymentStatusEnumBalanceUpdating), string(CardPaymentStatusEnumBalanceUpdated), string(CardPaymentStatusEnumInvoicing), string(CardPaymentStatusEnumInvoiced), string(CardPaymentStatusEnumPaymentFailed), string(CardPaymentStatusEnumBalanceUpdateFailed), string(CardPaymentStatusEnumInvoicingFailed), string(CardPaymentStatusEnumFinished), string(CardPaymentStatusEnumPaymentDeclineFailed), string(CardPaymentStatusEnumFailed),
	}
}

func CardPaymentStatusEnumAll() []CardPaymentStatusEnum {
	return []CardPaymentStatusEnum{
		CardPaymentStatusEnumRequested, CardPaymentStatusEnumProcessing, CardPaymentStatusEnumPaymentSuccess, CardPaymentStatusEnumPaymentDeclined, CardPaymentStatusEnumBalanceUpdating, CardPaymentStatusEnumBalanceUpdated, CardPaymentStatusEnumInvoicing, CardPaymentStatusEnumInvoiced, CardPaymentStatusEnumPaymentFailed, CardPaymentStatusEnumBalanceUpdateFailed, CardPaymentStatusEnumInvoicingFailed, CardPaymentStatusEnumFinished, CardPaymentStatusEnumPaymentDeclineFailed, CardPaymentStatusEnumFailed,
	}
}

func CardPaymentStatusEnumAllPublic() []CardPaymentStatusEnum {
	return []CardPaymentStatusEnum{
		CardPaymentStatusEnumRequested, CardPaymentStatusEnumProcessing, CardPaymentStatusEnumPaymentSuccess, CardPaymentStatusEnumPaymentDeclined, CardPaymentStatusEnumBalanceUpdating, CardPaymentStatusEnumBalanceUpdated, CardPaymentStatusEnumInvoicing, CardPaymentStatusEnumInvoiced, CardPaymentStatusEnumPaymentFailed, CardPaymentStatusEnumBalanceUpdateFailed, CardPaymentStatusEnumInvoicingFailed, CardPaymentStatusEnumFinished, CardPaymentStatusEnumPaymentDeclineFailed, CardPaymentStatusEnumFailed,
	}
}

func CardPaymentStatusEnumAllPrivate() []CardPaymentStatusEnum {
	return []CardPaymentStatusEnum{}
}

func CardPaymentStatusEnumDefault() CardPaymentStatusEnum {
	return CardPaymentStatusEnumRequested
}

func (enum CardPaymentStatusEnum) IsRequested() bool {
	return enum.Is(CardPaymentStatusEnumRequested)
}

func (enum CardPaymentStatusEnum) IsProcessing() bool {
	return enum.Is(CardPaymentStatusEnumProcessing)
}

func (enum CardPaymentStatusEnum) IsPaymentSuccess() bool {
	return enum.Is(CardPaymentStatusEnumPaymentSuccess)
}

func (enum CardPaymentStatusEnum) IsPaymentDeclined() bool {
	return enum.Is(CardPaymentStatusEnumPaymentDeclined)
}

func (enum CardPaymentStatusEnum) IsBalanceUpdating() bool {
	return enum.Is(CardPaymentStatusEnumBalanceUpdating)
}

func (enum CardPaymentStatusEnum) IsBalanceUpdated() bool {
	return enum.Is(CardPaymentStatusEnumBalanceUpdated)
}

func (enum CardPaymentStatusEnum) IsInvoicing() bool {
	return enum.Is(CardPaymentStatusEnumInvoicing)
}

func (enum CardPaymentStatusEnum) IsInvoiced() bool {
	return enum.Is(CardPaymentStatusEnumInvoiced)
}

func (enum CardPaymentStatusEnum) IsPaymentFailed() bool {
	return enum.Is(CardPaymentStatusEnumPaymentFailed)
}

func (enum CardPaymentStatusEnum) IsBalanceUpdateFailed() bool {
	return enum.Is(CardPaymentStatusEnumBalanceUpdateFailed)
}

func (enum CardPaymentStatusEnum) IsInvoicingFailed() bool {
	return enum.Is(CardPaymentStatusEnumInvoicingFailed)
}

func (enum CardPaymentStatusEnum) IsFinished() bool {
	return enum.Is(CardPaymentStatusEnumFinished)
}

func (enum CardPaymentStatusEnum) IsPaymentDeclineFailed() bool {
	return enum.Is(CardPaymentStatusEnumPaymentDeclineFailed)
}

func (enum CardPaymentStatusEnum) IsFailed() bool {
	return enum.Is(CardPaymentStatusEnumFailed)
}
