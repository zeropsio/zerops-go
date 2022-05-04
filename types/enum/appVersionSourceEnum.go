// Generated ZEROPS sdk

package enum

type AppVersionSourceEnum string

const (
	AppVersionSourceEnumCli    = AppVersionSourceEnum("CLI")
	AppVersionSourceEnumGui    = AppVersionSourceEnum("GUI")
	AppVersionSourceEnumGithub = AppVersionSourceEnum("GITHUB")
	AppVersionSourceEnumGitlab = AppVersionSourceEnum("GITLAB")
	AppVersionSourceEnumGit    = AppVersionSourceEnum("GIT")
)

func NewAppVersionSourceEnumFromString(value string) (out AppVersionSourceEnum, err error) {
	return AppVersionSourceEnum(value), nil
}

func (enum AppVersionSourceEnum) String() string {
	return string(enum)
}

func (enum AppVersionSourceEnum) Native() string {
	return string(enum)
}

func (enum AppVersionSourceEnum) Values() []AppVersionSourceEnum {
	return AppVersionSourceEnumAll()
}

func (enum AppVersionSourceEnum) PublicValues() []AppVersionSourceEnum {
	return AppVersionSourceEnumAllPublic()
}

func (enum AppVersionSourceEnum) PrivateValues() []AppVersionSourceEnum {
	return AppVersionSourceEnumAllPrivate()
}

func (enum AppVersionSourceEnum) DefaultValue() AppVersionSourceEnum {
	return AppVersionSourceEnumDefault()
}

func (enum AppVersionSourceEnum) Is(values ...AppVersionSourceEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func AppVersionSourceEnumAllStrings() []string {
	return []string{
		string(AppVersionSourceEnumCli), string(AppVersionSourceEnumGui), string(AppVersionSourceEnumGithub), string(AppVersionSourceEnumGitlab), string(AppVersionSourceEnumGit),
	}
}

func AppVersionSourceEnumAll() []AppVersionSourceEnum {
	return []AppVersionSourceEnum{
		AppVersionSourceEnumCli, AppVersionSourceEnumGui, AppVersionSourceEnumGithub, AppVersionSourceEnumGitlab, AppVersionSourceEnumGit,
	}
}

func AppVersionSourceEnumAllPublic() []AppVersionSourceEnum {
	return []AppVersionSourceEnum{
		AppVersionSourceEnumCli, AppVersionSourceEnumGui, AppVersionSourceEnumGithub, AppVersionSourceEnumGitlab, AppVersionSourceEnumGit,
	}
}

func AppVersionSourceEnumAllPrivate() []AppVersionSourceEnum {
	return []AppVersionSourceEnum{}
}

func AppVersionSourceEnumDefault() AppVersionSourceEnum {
	return AppVersionSourceEnumCli
}

func (enum AppVersionSourceEnum) IsCli() bool {
	return enum.Is(AppVersionSourceEnumCli)
}

func (enum AppVersionSourceEnum) IsGui() bool {
	return enum.Is(AppVersionSourceEnumGui)
}

func (enum AppVersionSourceEnum) IsGithub() bool {
	return enum.Is(AppVersionSourceEnumGithub)
}

func (enum AppVersionSourceEnum) IsGitlab() bool {
	return enum.Is(AppVersionSourceEnumGitlab)
}

func (enum AppVersionSourceEnum) IsGit() bool {
	return enum.Is(AppVersionSourceEnumGit)
}
