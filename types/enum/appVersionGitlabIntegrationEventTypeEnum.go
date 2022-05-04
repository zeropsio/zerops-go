// Generated ZEROPS sdk

package enum

type AppVersionGitlabIntegrationEventTypeEnum string

const (
	AppVersionGitlabIntegrationEventTypeEnumBranch = AppVersionGitlabIntegrationEventTypeEnum("BRANCH")
	AppVersionGitlabIntegrationEventTypeEnumTag    = AppVersionGitlabIntegrationEventTypeEnum("TAG")
)

func NewAppVersionGitlabIntegrationEventTypeEnumFromString(value string) (out AppVersionGitlabIntegrationEventTypeEnum, err error) {
	return AppVersionGitlabIntegrationEventTypeEnum(value), nil
}

func (enum AppVersionGitlabIntegrationEventTypeEnum) String() string {
	return string(enum)
}

func (enum AppVersionGitlabIntegrationEventTypeEnum) Native() string {
	return string(enum)
}

func (enum AppVersionGitlabIntegrationEventTypeEnum) Values() []AppVersionGitlabIntegrationEventTypeEnum {
	return AppVersionGitlabIntegrationEventTypeEnumAll()
}

func (enum AppVersionGitlabIntegrationEventTypeEnum) PublicValues() []AppVersionGitlabIntegrationEventTypeEnum {
	return AppVersionGitlabIntegrationEventTypeEnumAllPublic()
}

func (enum AppVersionGitlabIntegrationEventTypeEnum) PrivateValues() []AppVersionGitlabIntegrationEventTypeEnum {
	return AppVersionGitlabIntegrationEventTypeEnumAllPrivate()
}

func (enum AppVersionGitlabIntegrationEventTypeEnum) DefaultValue() AppVersionGitlabIntegrationEventTypeEnum {
	return AppVersionGitlabIntegrationEventTypeEnumDefault()
}

func (enum AppVersionGitlabIntegrationEventTypeEnum) Is(values ...AppVersionGitlabIntegrationEventTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func AppVersionGitlabIntegrationEventTypeEnumAllStrings() []string {
	return []string{
		string(AppVersionGitlabIntegrationEventTypeEnumBranch), string(AppVersionGitlabIntegrationEventTypeEnumTag),
	}
}

func AppVersionGitlabIntegrationEventTypeEnumAll() []AppVersionGitlabIntegrationEventTypeEnum {
	return []AppVersionGitlabIntegrationEventTypeEnum{
		AppVersionGitlabIntegrationEventTypeEnumBranch, AppVersionGitlabIntegrationEventTypeEnumTag,
	}
}

func AppVersionGitlabIntegrationEventTypeEnumAllPublic() []AppVersionGitlabIntegrationEventTypeEnum {
	return []AppVersionGitlabIntegrationEventTypeEnum{
		AppVersionGitlabIntegrationEventTypeEnumBranch, AppVersionGitlabIntegrationEventTypeEnumTag,
	}
}

func AppVersionGitlabIntegrationEventTypeEnumAllPrivate() []AppVersionGitlabIntegrationEventTypeEnum {
	return []AppVersionGitlabIntegrationEventTypeEnum{}
}

func AppVersionGitlabIntegrationEventTypeEnumDefault() AppVersionGitlabIntegrationEventTypeEnum {
	return AppVersionGitlabIntegrationEventTypeEnumBranch
}

func (enum AppVersionGitlabIntegrationEventTypeEnum) IsBranch() bool {
	return enum.Is(AppVersionGitlabIntegrationEventTypeEnumBranch)
}

func (enum AppVersionGitlabIntegrationEventTypeEnum) IsTag() bool {
	return enum.Is(AppVersionGitlabIntegrationEventTypeEnumTag)
}
