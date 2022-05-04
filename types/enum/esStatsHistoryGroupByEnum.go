// Generated ZEROPS sdk

package enum

type EsStatsHistoryGroupByEnum string

const (
	EsStatsHistoryGroupByEnumServiceStackId = EsStatsHistoryGroupByEnum("serviceStackId")
	EsStatsHistoryGroupByEnumServiceId      = EsStatsHistoryGroupByEnum("serviceId")
	EsStatsHistoryGroupByEnumContainerId    = EsStatsHistoryGroupByEnum("containerId")
)

func NewEsStatsHistoryGroupByEnumFromString(value string) (out EsStatsHistoryGroupByEnum, err error) {
	return EsStatsHistoryGroupByEnum(value), nil
}

func (enum EsStatsHistoryGroupByEnum) String() string {
	return string(enum)
}

func (enum EsStatsHistoryGroupByEnum) Native() string {
	return string(enum)
}

func (enum EsStatsHistoryGroupByEnum) Values() []EsStatsHistoryGroupByEnum {
	return EsStatsHistoryGroupByEnumAll()
}

func (enum EsStatsHistoryGroupByEnum) PublicValues() []EsStatsHistoryGroupByEnum {
	return EsStatsHistoryGroupByEnumAllPublic()
}

func (enum EsStatsHistoryGroupByEnum) PrivateValues() []EsStatsHistoryGroupByEnum {
	return EsStatsHistoryGroupByEnumAllPrivate()
}

func (enum EsStatsHistoryGroupByEnum) DefaultValue() EsStatsHistoryGroupByEnum {
	return EsStatsHistoryGroupByEnumDefault()
}

func (enum EsStatsHistoryGroupByEnum) Is(values ...EsStatsHistoryGroupByEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func EsStatsHistoryGroupByEnumAllStrings() []string {
	return []string{
		string(EsStatsHistoryGroupByEnumServiceStackId), string(EsStatsHistoryGroupByEnumServiceId), string(EsStatsHistoryGroupByEnumContainerId),
	}
}

func EsStatsHistoryGroupByEnumAll() []EsStatsHistoryGroupByEnum {
	return []EsStatsHistoryGroupByEnum{
		EsStatsHistoryGroupByEnumServiceStackId, EsStatsHistoryGroupByEnumServiceId, EsStatsHistoryGroupByEnumContainerId,
	}
}

func EsStatsHistoryGroupByEnumAllPublic() []EsStatsHistoryGroupByEnum {
	return []EsStatsHistoryGroupByEnum{
		EsStatsHistoryGroupByEnumServiceStackId, EsStatsHistoryGroupByEnumServiceId, EsStatsHistoryGroupByEnumContainerId,
	}
}

func EsStatsHistoryGroupByEnumAllPrivate() []EsStatsHistoryGroupByEnum {
	return []EsStatsHistoryGroupByEnum{}
}

func EsStatsHistoryGroupByEnumDefault() EsStatsHistoryGroupByEnum {
	return EsStatsHistoryGroupByEnumServiceStackId
}

func (enum EsStatsHistoryGroupByEnum) IsServiceStackId() bool {
	return enum.Is(EsStatsHistoryGroupByEnumServiceStackId)
}

func (enum EsStatsHistoryGroupByEnum) IsServiceId() bool {
	return enum.Is(EsStatsHistoryGroupByEnumServiceId)
}

func (enum EsStatsHistoryGroupByEnum) IsContainerId() bool {
	return enum.Is(EsStatsHistoryGroupByEnumContainerId)
}
