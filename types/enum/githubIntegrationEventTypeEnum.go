// Generated ZEROPS sdk

package enum

type GithubIntegrationEventTypeEnum string

const (
	GithubIntegrationEventTypeEnumBranch = GithubIntegrationEventTypeEnum("BRANCH")
	GithubIntegrationEventTypeEnumTag    = GithubIntegrationEventTypeEnum("TAG")
)

func NewGithubIntegrationEventTypeEnumFromString(value string) (out GithubIntegrationEventTypeEnum, err error) {
	return GithubIntegrationEventTypeEnum(value), nil
}

func (enum GithubIntegrationEventTypeEnum) String() string {
	return string(enum)
}

func (enum GithubIntegrationEventTypeEnum) Native() string {
	return string(enum)
}

func (enum GithubIntegrationEventTypeEnum) Values() []GithubIntegrationEventTypeEnum {
	return GithubIntegrationEventTypeEnumAll()
}

func (enum GithubIntegrationEventTypeEnum) PublicValues() []GithubIntegrationEventTypeEnum {
	return GithubIntegrationEventTypeEnumAllPublic()
}

func (enum GithubIntegrationEventTypeEnum) PrivateValues() []GithubIntegrationEventTypeEnum {
	return GithubIntegrationEventTypeEnumAllPrivate()
}

func (enum GithubIntegrationEventTypeEnum) DefaultValue() GithubIntegrationEventTypeEnum {
	return GithubIntegrationEventTypeEnumDefault()
}

func (enum GithubIntegrationEventTypeEnum) Is(values ...GithubIntegrationEventTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func GithubIntegrationEventTypeEnumAllStrings() []string {
	return []string{
		string(GithubIntegrationEventTypeEnumBranch), string(GithubIntegrationEventTypeEnumTag),
	}
}

func GithubIntegrationEventTypeEnumAll() []GithubIntegrationEventTypeEnum {
	return []GithubIntegrationEventTypeEnum{
		GithubIntegrationEventTypeEnumBranch, GithubIntegrationEventTypeEnumTag,
	}
}

func GithubIntegrationEventTypeEnumAllPublic() []GithubIntegrationEventTypeEnum {
	return []GithubIntegrationEventTypeEnum{
		GithubIntegrationEventTypeEnumBranch, GithubIntegrationEventTypeEnumTag,
	}
}

func GithubIntegrationEventTypeEnumAllPrivate() []GithubIntegrationEventTypeEnum {
	return []GithubIntegrationEventTypeEnum{}
}

func GithubIntegrationEventTypeEnumDefault() GithubIntegrationEventTypeEnum {
	return GithubIntegrationEventTypeEnumBranch
}

func (enum GithubIntegrationEventTypeEnum) IsBranch() bool {
	return enum.Is(GithubIntegrationEventTypeEnumBranch)
}

func (enum GithubIntegrationEventTypeEnum) IsTag() bool {
	return enum.Is(GithubIntegrationEventTypeEnumTag)
}
