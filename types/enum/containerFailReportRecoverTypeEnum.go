// Generated ZEROPS sdk

package enum

type ContainerFailReportRecoverTypeEnum string

const (
	ContainerFailReportRecoverTypeEnumContainer = ContainerFailReportRecoverTypeEnum("CONTAINER")
	ContainerFailReportRecoverTypeEnumService   = ContainerFailReportRecoverTypeEnum("SERVICE")
	ContainerFailReportRecoverTypeEnumNone      = ContainerFailReportRecoverTypeEnum("NONE")
)

func NewContainerFailReportRecoverTypeEnumFromString(value string) (out ContainerFailReportRecoverTypeEnum, err error) {
	return ContainerFailReportRecoverTypeEnum(value), nil
}

func (enum ContainerFailReportRecoverTypeEnum) String() string {
	return string(enum)
}

func (enum ContainerFailReportRecoverTypeEnum) Native() string {
	return string(enum)
}

func (enum ContainerFailReportRecoverTypeEnum) Values() []ContainerFailReportRecoverTypeEnum {
	return ContainerFailReportRecoverTypeEnumAll()
}

func (enum ContainerFailReportRecoverTypeEnum) PublicValues() []ContainerFailReportRecoverTypeEnum {
	return ContainerFailReportRecoverTypeEnumAllPublic()
}

func (enum ContainerFailReportRecoverTypeEnum) PrivateValues() []ContainerFailReportRecoverTypeEnum {
	return ContainerFailReportRecoverTypeEnumAllPrivate()
}

func (enum ContainerFailReportRecoverTypeEnum) DefaultValue() ContainerFailReportRecoverTypeEnum {
	return ContainerFailReportRecoverTypeEnumDefault()
}

func (enum ContainerFailReportRecoverTypeEnum) Is(values ...ContainerFailReportRecoverTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ContainerFailReportRecoverTypeEnumAllStrings() []string {
	return []string{
		string(ContainerFailReportRecoverTypeEnumContainer), string(ContainerFailReportRecoverTypeEnumService), string(ContainerFailReportRecoverTypeEnumNone),
	}
}

func ContainerFailReportRecoverTypeEnumAll() []ContainerFailReportRecoverTypeEnum {
	return []ContainerFailReportRecoverTypeEnum{
		ContainerFailReportRecoverTypeEnumContainer, ContainerFailReportRecoverTypeEnumService, ContainerFailReportRecoverTypeEnumNone,
	}
}

func ContainerFailReportRecoverTypeEnumAllPublic() []ContainerFailReportRecoverTypeEnum {
	return []ContainerFailReportRecoverTypeEnum{
		ContainerFailReportRecoverTypeEnumContainer, ContainerFailReportRecoverTypeEnumService, ContainerFailReportRecoverTypeEnumNone,
	}
}

func ContainerFailReportRecoverTypeEnumAllPrivate() []ContainerFailReportRecoverTypeEnum {
	return []ContainerFailReportRecoverTypeEnum{}
}

func ContainerFailReportRecoverTypeEnumDefault() ContainerFailReportRecoverTypeEnum {
	return ""
}

func (enum ContainerFailReportRecoverTypeEnum) IsContainer() bool {
	return enum.Is(ContainerFailReportRecoverTypeEnumContainer)
}

func (enum ContainerFailReportRecoverTypeEnum) IsService() bool {
	return enum.Is(ContainerFailReportRecoverTypeEnumService)
}

func (enum ContainerFailReportRecoverTypeEnum) IsNone() bool {
	return enum.Is(ContainerFailReportRecoverTypeEnumNone)
}
