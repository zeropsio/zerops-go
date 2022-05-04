// Generated ZEROPS sdk

package enum

type GitlabUrlActionEnum string

const (
	GitlabUrlActionEnumRegistration = GitlabUrlActionEnum("REGISTRATION")
	GitlabUrlActionEnumLogin        = GitlabUrlActionEnum("LOGIN")
	GitlabUrlActionEnumRepository   = GitlabUrlActionEnum("REPOSITORY")
)

func NewGitlabUrlActionEnumFromString(value string) (out GitlabUrlActionEnum, err error) {
	return GitlabUrlActionEnum(value), nil
}

func (enum GitlabUrlActionEnum) String() string {
	return string(enum)
}

func (enum GitlabUrlActionEnum) Native() string {
	return string(enum)
}

func (enum GitlabUrlActionEnum) Values() []GitlabUrlActionEnum {
	return GitlabUrlActionEnumAll()
}

func (enum GitlabUrlActionEnum) PublicValues() []GitlabUrlActionEnum {
	return GitlabUrlActionEnumAllPublic()
}

func (enum GitlabUrlActionEnum) PrivateValues() []GitlabUrlActionEnum {
	return GitlabUrlActionEnumAllPrivate()
}

func (enum GitlabUrlActionEnum) DefaultValue() GitlabUrlActionEnum {
	return GitlabUrlActionEnumDefault()
}

func (enum GitlabUrlActionEnum) Is(values ...GitlabUrlActionEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func GitlabUrlActionEnumAllStrings() []string {
	return []string{
		string(GitlabUrlActionEnumRegistration), string(GitlabUrlActionEnumLogin), string(GitlabUrlActionEnumRepository),
	}
}

func GitlabUrlActionEnumAll() []GitlabUrlActionEnum {
	return []GitlabUrlActionEnum{
		GitlabUrlActionEnumRegistration, GitlabUrlActionEnumLogin, GitlabUrlActionEnumRepository,
	}
}

func GitlabUrlActionEnumAllPublic() []GitlabUrlActionEnum {
	return []GitlabUrlActionEnum{
		GitlabUrlActionEnumRegistration, GitlabUrlActionEnumLogin, GitlabUrlActionEnumRepository,
	}
}

func GitlabUrlActionEnumAllPrivate() []GitlabUrlActionEnum {
	return []GitlabUrlActionEnum{}
}

func GitlabUrlActionEnumDefault() GitlabUrlActionEnum {
	return GitlabUrlActionEnumRegistration
}

func (enum GitlabUrlActionEnum) IsRegistration() bool {
	return enum.Is(GitlabUrlActionEnumRegistration)
}

func (enum GitlabUrlActionEnum) IsLogin() bool {
	return enum.Is(GitlabUrlActionEnumLogin)
}

func (enum GitlabUrlActionEnum) IsRepository() bool {
	return enum.Is(GitlabUrlActionEnumRepository)
}
