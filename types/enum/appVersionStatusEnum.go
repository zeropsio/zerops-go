// Generated ZEROPS sdk

package enum

type AppVersionStatusEnum string

const (
	AppVersionStatusEnumUploading              = AppVersionStatusEnum("UPLOADING")
	AppVersionStatusEnumWaitingToDeploy        = AppVersionStatusEnum("WAITING_TO_DEPLOY")
	AppVersionStatusEnumDeploying              = AppVersionStatusEnum("DEPLOYING")
	AppVersionStatusEnumDeployFailed           = AppVersionStatusEnum("DEPLOY_FAILED")
	AppVersionStatusEnumActive                 = AppVersionStatusEnum("ACTIVE")
	AppVersionStatusEnumBackup                 = AppVersionStatusEnum("BACKUP")
	AppVersionStatusEnumBuilding               = AppVersionStatusEnum("BUILDING")
	AppVersionStatusEnumWaitingToBuild         = AppVersionStatusEnum("WAITING_TO_BUILD")
	AppVersionStatusEnumBuildFailed            = AppVersionStatusEnum("BUILD_FAILED")
	AppVersionStatusEnumPreparingRuntime       = AppVersionStatusEnum("PREPARING_RUNTIME")
	AppVersionStatusEnumPreparingRuntimeFailed = AppVersionStatusEnum("PREPARING_RUNTIME_FAILED")
	AppVersionStatusEnumBuildValidationFailed  = AppVersionStatusEnum("BUILD_VALIDATION_FAILED")
)

func NewAppVersionStatusEnumFromString(value string) (out AppVersionStatusEnum, err error) {
	return AppVersionStatusEnum(value), nil
}

func (enum AppVersionStatusEnum) String() string {
	return string(enum)
}

func (enum AppVersionStatusEnum) Native() string {
	return string(enum)
}

func (enum AppVersionStatusEnum) Values() []AppVersionStatusEnum {
	return AppVersionStatusEnumAll()
}

func (enum AppVersionStatusEnum) PublicValues() []AppVersionStatusEnum {
	return AppVersionStatusEnumAllPublic()
}

func (enum AppVersionStatusEnum) PrivateValues() []AppVersionStatusEnum {
	return AppVersionStatusEnumAllPrivate()
}

func (enum AppVersionStatusEnum) DefaultValue() AppVersionStatusEnum {
	return AppVersionStatusEnumDefault()
}

func (enum AppVersionStatusEnum) Is(values ...AppVersionStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func AppVersionStatusEnumAllStrings() []string {
	return []string{
		string(AppVersionStatusEnumUploading), string(AppVersionStatusEnumWaitingToDeploy), string(AppVersionStatusEnumDeploying), string(AppVersionStatusEnumDeployFailed), string(AppVersionStatusEnumActive), string(AppVersionStatusEnumBackup), string(AppVersionStatusEnumBuilding), string(AppVersionStatusEnumWaitingToBuild), string(AppVersionStatusEnumBuildFailed), string(AppVersionStatusEnumPreparingRuntime), string(AppVersionStatusEnumPreparingRuntimeFailed), string(AppVersionStatusEnumBuildValidationFailed),
	}
}

func AppVersionStatusEnumAll() []AppVersionStatusEnum {
	return []AppVersionStatusEnum{
		AppVersionStatusEnumUploading, AppVersionStatusEnumWaitingToDeploy, AppVersionStatusEnumDeploying, AppVersionStatusEnumDeployFailed, AppVersionStatusEnumActive, AppVersionStatusEnumBackup, AppVersionStatusEnumBuilding, AppVersionStatusEnumWaitingToBuild, AppVersionStatusEnumBuildFailed, AppVersionStatusEnumPreparingRuntime, AppVersionStatusEnumPreparingRuntimeFailed, AppVersionStatusEnumBuildValidationFailed,
	}
}

func AppVersionStatusEnumAllPublic() []AppVersionStatusEnum {
	return []AppVersionStatusEnum{
		AppVersionStatusEnumUploading, AppVersionStatusEnumWaitingToDeploy, AppVersionStatusEnumDeploying, AppVersionStatusEnumDeployFailed, AppVersionStatusEnumActive, AppVersionStatusEnumBackup, AppVersionStatusEnumBuilding, AppVersionStatusEnumWaitingToBuild, AppVersionStatusEnumBuildFailed, AppVersionStatusEnumPreparingRuntime, AppVersionStatusEnumPreparingRuntimeFailed, AppVersionStatusEnumBuildValidationFailed,
	}
}

func AppVersionStatusEnumAllPrivate() []AppVersionStatusEnum {
	return []AppVersionStatusEnum{}
}

func AppVersionStatusEnumDefault() AppVersionStatusEnum {
	return AppVersionStatusEnumUploading
}

func (enum AppVersionStatusEnum) IsUploading() bool {
	return enum.Is(AppVersionStatusEnumUploading)
}

func (enum AppVersionStatusEnum) IsWaitingToDeploy() bool {
	return enum.Is(AppVersionStatusEnumWaitingToDeploy)
}

func (enum AppVersionStatusEnum) IsDeploying() bool {
	return enum.Is(AppVersionStatusEnumDeploying)
}

func (enum AppVersionStatusEnum) IsDeployFailed() bool {
	return enum.Is(AppVersionStatusEnumDeployFailed)
}

func (enum AppVersionStatusEnum) IsActive() bool {
	return enum.Is(AppVersionStatusEnumActive)
}

func (enum AppVersionStatusEnum) IsBackup() bool {
	return enum.Is(AppVersionStatusEnumBackup)
}

func (enum AppVersionStatusEnum) IsBuilding() bool {
	return enum.Is(AppVersionStatusEnumBuilding)
}

func (enum AppVersionStatusEnum) IsWaitingToBuild() bool {
	return enum.Is(AppVersionStatusEnumWaitingToBuild)
}

func (enum AppVersionStatusEnum) IsBuildFailed() bool {
	return enum.Is(AppVersionStatusEnumBuildFailed)
}

func (enum AppVersionStatusEnum) IsPreparingRuntime() bool {
	return enum.Is(AppVersionStatusEnumPreparingRuntime)
}

func (enum AppVersionStatusEnum) IsPreparingRuntimeFailed() bool {
	return enum.Is(AppVersionStatusEnumPreparingRuntimeFailed)
}

func (enum AppVersionStatusEnum) IsBuildValidationFailed() bool {
	return enum.Is(AppVersionStatusEnumBuildValidationFailed)
}
