// Generated ZEROPS sdk

package enum

type AppVersionGithubIntegrationEventTypeEnum string

const (
	AppVersionGithubIntegrationEventTypeEnumBranch = AppVersionGithubIntegrationEventTypeEnum("BRANCH")
	AppVersionGithubIntegrationEventTypeEnumTag    = AppVersionGithubIntegrationEventTypeEnum("TAG")
)

func NewAppVersionGithubIntegrationEventTypeEnumFromString(value string) (out AppVersionGithubIntegrationEventTypeEnum, err error) {
	return AppVersionGithubIntegrationEventTypeEnum(value), nil
}

func (enum AppVersionGithubIntegrationEventTypeEnum) String() string {
	return string(enum)
}

func (enum AppVersionGithubIntegrationEventTypeEnum) Native() string {
	return string(enum)
}

func (enum AppVersionGithubIntegrationEventTypeEnum) Values() []AppVersionGithubIntegrationEventTypeEnum {
	return AppVersionGithubIntegrationEventTypeEnumAll()
}

func (enum AppVersionGithubIntegrationEventTypeEnum) PublicValues() []AppVersionGithubIntegrationEventTypeEnum {
	return AppVersionGithubIntegrationEventTypeEnumAllPublic()
}

func (enum AppVersionGithubIntegrationEventTypeEnum) PrivateValues() []AppVersionGithubIntegrationEventTypeEnum {
	return AppVersionGithubIntegrationEventTypeEnumAllPrivate()
}

func (enum AppVersionGithubIntegrationEventTypeEnum) DefaultValue() AppVersionGithubIntegrationEventTypeEnum {
	return AppVersionGithubIntegrationEventTypeEnumDefault()
}

func (enum AppVersionGithubIntegrationEventTypeEnum) Is(values ...AppVersionGithubIntegrationEventTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func AppVersionGithubIntegrationEventTypeEnumAllStrings() []string {
	return []string{
		string(AppVersionGithubIntegrationEventTypeEnumBranch), string(AppVersionGithubIntegrationEventTypeEnumTag),
	}
}

func AppVersionGithubIntegrationEventTypeEnumAll() []AppVersionGithubIntegrationEventTypeEnum {
	return []AppVersionGithubIntegrationEventTypeEnum{
		AppVersionGithubIntegrationEventTypeEnumBranch, AppVersionGithubIntegrationEventTypeEnumTag,
	}
}

func AppVersionGithubIntegrationEventTypeEnumAllPublic() []AppVersionGithubIntegrationEventTypeEnum {
	return []AppVersionGithubIntegrationEventTypeEnum{
		AppVersionGithubIntegrationEventTypeEnumBranch, AppVersionGithubIntegrationEventTypeEnumTag,
	}
}

func AppVersionGithubIntegrationEventTypeEnumAllPrivate() []AppVersionGithubIntegrationEventTypeEnum {
	return []AppVersionGithubIntegrationEventTypeEnum{}
}

func AppVersionGithubIntegrationEventTypeEnumDefault() AppVersionGithubIntegrationEventTypeEnum {
	return AppVersionGithubIntegrationEventTypeEnumBranch
}

func (enum AppVersionGithubIntegrationEventTypeEnum) IsBranch() bool {
	return enum.Is(AppVersionGithubIntegrationEventTypeEnumBranch)
}

func (enum AppVersionGithubIntegrationEventTypeEnum) IsTag() bool {
	return enum.Is(AppVersionGithubIntegrationEventTypeEnumTag)
}
