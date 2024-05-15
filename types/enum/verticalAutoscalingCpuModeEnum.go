// Generated ZEROPS sdk

package enum

type VerticalAutoscalingCpuModeEnum string

const (
	VerticalAutoscalingCpuModeEnumDedicated = VerticalAutoscalingCpuModeEnum("DEDICATED")
	VerticalAutoscalingCpuModeEnumShared    = VerticalAutoscalingCpuModeEnum("SHARED")
)

func NewVerticalAutoscalingCpuModeEnumFromString(value string) (out VerticalAutoscalingCpuModeEnum, err error) {
	return VerticalAutoscalingCpuModeEnum(value), nil
}

func (enum VerticalAutoscalingCpuModeEnum) String() string {
	return string(enum)
}

func (enum VerticalAutoscalingCpuModeEnum) Native() string {
	return string(enum)
}

func (enum VerticalAutoscalingCpuModeEnum) Values() []VerticalAutoscalingCpuModeEnum {
	return VerticalAutoscalingCpuModeEnumAll()
}

func (enum VerticalAutoscalingCpuModeEnum) PublicValues() []VerticalAutoscalingCpuModeEnum {
	return VerticalAutoscalingCpuModeEnumAllPublic()
}

func (enum VerticalAutoscalingCpuModeEnum) PrivateValues() []VerticalAutoscalingCpuModeEnum {
	return VerticalAutoscalingCpuModeEnumAllPrivate()
}

func (enum VerticalAutoscalingCpuModeEnum) DefaultValue() VerticalAutoscalingCpuModeEnum {
	return VerticalAutoscalingCpuModeEnumDefault()
}

func (enum VerticalAutoscalingCpuModeEnum) Is(values ...VerticalAutoscalingCpuModeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func VerticalAutoscalingCpuModeEnumAllStrings() []string {
	return []string{
		string(VerticalAutoscalingCpuModeEnumDedicated), string(VerticalAutoscalingCpuModeEnumShared),
	}
}

func VerticalAutoscalingCpuModeEnumAll() []VerticalAutoscalingCpuModeEnum {
	return []VerticalAutoscalingCpuModeEnum{
		VerticalAutoscalingCpuModeEnumDedicated, VerticalAutoscalingCpuModeEnumShared,
	}
}

func VerticalAutoscalingCpuModeEnumAllPublic() []VerticalAutoscalingCpuModeEnum {
	return []VerticalAutoscalingCpuModeEnum{
		VerticalAutoscalingCpuModeEnumDedicated, VerticalAutoscalingCpuModeEnumShared,
	}
}

func VerticalAutoscalingCpuModeEnumAllPrivate() []VerticalAutoscalingCpuModeEnum {
	return []VerticalAutoscalingCpuModeEnum{}
}

func VerticalAutoscalingCpuModeEnumDefault() VerticalAutoscalingCpuModeEnum {
	return VerticalAutoscalingCpuModeEnumShared
}

func (enum VerticalAutoscalingCpuModeEnum) IsDedicated() bool {
	return enum.Is(VerticalAutoscalingCpuModeEnumDedicated)
}

func (enum VerticalAutoscalingCpuModeEnum) IsShared() bool {
	return enum.Is(VerticalAutoscalingCpuModeEnumShared)
}
