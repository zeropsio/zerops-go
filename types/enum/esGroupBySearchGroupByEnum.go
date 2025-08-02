// Generated ZEROPS sdk

package enum

type EsGroupBySearchGroupByEnum string

const (
	EsGroupBySearchGroupByEnumProjectId = EsGroupBySearchGroupByEnum("projectId")
	EsGroupBySearchGroupByEnumStackId   = EsGroupBySearchGroupByEnum("stackId")
	EsGroupBySearchGroupByEnumMetricId  = EsGroupBySearchGroupByEnum("metricId")
)

func NewEsGroupBySearchGroupByEnumFromString(value string) (out EsGroupBySearchGroupByEnum, err error) {
	return EsGroupBySearchGroupByEnum(value), nil
}

func (enum EsGroupBySearchGroupByEnum) String() string {
	return string(enum)
}

func (enum EsGroupBySearchGroupByEnum) Native() string {
	return string(enum)
}

func (enum EsGroupBySearchGroupByEnum) Values() []EsGroupBySearchGroupByEnum {
	return EsGroupBySearchGroupByEnumAll()
}

func (enum EsGroupBySearchGroupByEnum) PublicValues() []EsGroupBySearchGroupByEnum {
	return EsGroupBySearchGroupByEnumAllPublic()
}

func (enum EsGroupBySearchGroupByEnum) PrivateValues() []EsGroupBySearchGroupByEnum {
	return EsGroupBySearchGroupByEnumAllPrivate()
}

func (enum EsGroupBySearchGroupByEnum) DefaultValue() EsGroupBySearchGroupByEnum {
	return EsGroupBySearchGroupByEnumDefault()
}

func (enum EsGroupBySearchGroupByEnum) Is(values ...EsGroupBySearchGroupByEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func EsGroupBySearchGroupByEnumAllStrings() []string {
	return []string{
		string(EsGroupBySearchGroupByEnumProjectId), string(EsGroupBySearchGroupByEnumStackId), string(EsGroupBySearchGroupByEnumMetricId),
	}
}

func EsGroupBySearchGroupByEnumAll() []EsGroupBySearchGroupByEnum {
	return []EsGroupBySearchGroupByEnum{
		EsGroupBySearchGroupByEnumProjectId, EsGroupBySearchGroupByEnumStackId, EsGroupBySearchGroupByEnumMetricId,
	}
}

func EsGroupBySearchGroupByEnumAllPublic() []EsGroupBySearchGroupByEnum {
	return []EsGroupBySearchGroupByEnum{
		EsGroupBySearchGroupByEnumProjectId, EsGroupBySearchGroupByEnumStackId, EsGroupBySearchGroupByEnumMetricId,
	}
}

func EsGroupBySearchGroupByEnumAllPrivate() []EsGroupBySearchGroupByEnum {
	return []EsGroupBySearchGroupByEnum{}
}

func EsGroupBySearchGroupByEnumDefault() EsGroupBySearchGroupByEnum {
	return ""
}

func (enum EsGroupBySearchGroupByEnum) IsProjectId() bool {
	return enum.Is(EsGroupBySearchGroupByEnumProjectId)
}

func (enum EsGroupBySearchGroupByEnum) IsStackId() bool {
	return enum.Is(EsGroupBySearchGroupByEnumStackId)
}

func (enum EsGroupBySearchGroupByEnum) IsMetricId() bool {
	return enum.Is(EsGroupBySearchGroupByEnumMetricId)
}
