// Generated ZEROPS sdk

package enum

type GitlabIntegrationEventTypeEnum string

const (
	GitlabIntegrationEventTypeEnumBranch = GitlabIntegrationEventTypeEnum("BRANCH")
	GitlabIntegrationEventTypeEnumTag    = GitlabIntegrationEventTypeEnum("TAG")
)

func NewGitlabIntegrationEventTypeEnumFromString(value string) (out GitlabIntegrationEventTypeEnum, err error) {
	return GitlabIntegrationEventTypeEnum(value), nil
}

func (enum GitlabIntegrationEventTypeEnum) String() string {
	return string(enum)
}

func (enum GitlabIntegrationEventTypeEnum) Native() string {
	return string(enum)
}

func (enum GitlabIntegrationEventTypeEnum) Values() []GitlabIntegrationEventTypeEnum {
	return GitlabIntegrationEventTypeEnumAll()
}

func (enum GitlabIntegrationEventTypeEnum) PublicValues() []GitlabIntegrationEventTypeEnum {
	return GitlabIntegrationEventTypeEnumAllPublic()
}

func (enum GitlabIntegrationEventTypeEnum) PrivateValues() []GitlabIntegrationEventTypeEnum {
	return GitlabIntegrationEventTypeEnumAllPrivate()
}

func (enum GitlabIntegrationEventTypeEnum) DefaultValue() GitlabIntegrationEventTypeEnum {
	return GitlabIntegrationEventTypeEnumDefault()
}

func (enum GitlabIntegrationEventTypeEnum) Is(values ...GitlabIntegrationEventTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func GitlabIntegrationEventTypeEnumAllStrings() []string {
	return []string{
		string(GitlabIntegrationEventTypeEnumBranch), string(GitlabIntegrationEventTypeEnumTag),
	}
}

func GitlabIntegrationEventTypeEnumAll() []GitlabIntegrationEventTypeEnum {
	return []GitlabIntegrationEventTypeEnum{
		GitlabIntegrationEventTypeEnumBranch, GitlabIntegrationEventTypeEnumTag,
	}
}

func GitlabIntegrationEventTypeEnumAllPublic() []GitlabIntegrationEventTypeEnum {
	return []GitlabIntegrationEventTypeEnum{
		GitlabIntegrationEventTypeEnumBranch, GitlabIntegrationEventTypeEnumTag,
	}
}

func GitlabIntegrationEventTypeEnumAllPrivate() []GitlabIntegrationEventTypeEnum {
	return []GitlabIntegrationEventTypeEnum{}
}

func GitlabIntegrationEventTypeEnumDefault() GitlabIntegrationEventTypeEnum {
	return GitlabIntegrationEventTypeEnumBranch
}

func (enum GitlabIntegrationEventTypeEnum) IsBranch() bool {
	return enum.Is(GitlabIntegrationEventTypeEnumBranch)
}

func (enum GitlabIntegrationEventTypeEnum) IsTag() bool {
	return enum.Is(GitlabIntegrationEventTypeEnumTag)
}
