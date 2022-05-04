// Generated ZEROPS sdk

package enum

type EsTransactionDebitGroupByEnum string

const (
	EsTransactionDebitGroupByEnumStackId   = EsTransactionDebitGroupByEnum("stackId")
	EsTransactionDebitGroupByEnumMetricId  = EsTransactionDebitGroupByEnum("metricId")
	EsTransactionDebitGroupByEnumProjectId = EsTransactionDebitGroupByEnum("projectId")
)

func NewEsTransactionDebitGroupByEnumFromString(value string) (out EsTransactionDebitGroupByEnum, err error) {
	return EsTransactionDebitGroupByEnum(value), nil
}

func (enum EsTransactionDebitGroupByEnum) String() string {
	return string(enum)
}

func (enum EsTransactionDebitGroupByEnum) Native() string {
	return string(enum)
}

func (enum EsTransactionDebitGroupByEnum) Values() []EsTransactionDebitGroupByEnum {
	return EsTransactionDebitGroupByEnumAll()
}

func (enum EsTransactionDebitGroupByEnum) PublicValues() []EsTransactionDebitGroupByEnum {
	return EsTransactionDebitGroupByEnumAllPublic()
}

func (enum EsTransactionDebitGroupByEnum) PrivateValues() []EsTransactionDebitGroupByEnum {
	return EsTransactionDebitGroupByEnumAllPrivate()
}

func (enum EsTransactionDebitGroupByEnum) DefaultValue() EsTransactionDebitGroupByEnum {
	return EsTransactionDebitGroupByEnumDefault()
}

func (enum EsTransactionDebitGroupByEnum) Is(values ...EsTransactionDebitGroupByEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func EsTransactionDebitGroupByEnumAllStrings() []string {
	return []string{
		string(EsTransactionDebitGroupByEnumStackId), string(EsTransactionDebitGroupByEnumMetricId), string(EsTransactionDebitGroupByEnumProjectId),
	}
}

func EsTransactionDebitGroupByEnumAll() []EsTransactionDebitGroupByEnum {
	return []EsTransactionDebitGroupByEnum{
		EsTransactionDebitGroupByEnumStackId, EsTransactionDebitGroupByEnumMetricId, EsTransactionDebitGroupByEnumProjectId,
	}
}

func EsTransactionDebitGroupByEnumAllPublic() []EsTransactionDebitGroupByEnum {
	return []EsTransactionDebitGroupByEnum{
		EsTransactionDebitGroupByEnumStackId, EsTransactionDebitGroupByEnumMetricId, EsTransactionDebitGroupByEnumProjectId,
	}
}

func EsTransactionDebitGroupByEnumAllPrivate() []EsTransactionDebitGroupByEnum {
	return []EsTransactionDebitGroupByEnum{}
}

func EsTransactionDebitGroupByEnumDefault() EsTransactionDebitGroupByEnum {
	return EsTransactionDebitGroupByEnumStackId
}

func (enum EsTransactionDebitGroupByEnum) IsStackId() bool {
	return enum.Is(EsTransactionDebitGroupByEnumStackId)
}

func (enum EsTransactionDebitGroupByEnum) IsMetricId() bool {
	return enum.Is(EsTransactionDebitGroupByEnumMetricId)
}

func (enum EsTransactionDebitGroupByEnum) IsProjectId() bool {
	return enum.Is(EsTransactionDebitGroupByEnumProjectId)
}
