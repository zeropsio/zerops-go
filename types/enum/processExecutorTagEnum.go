// Generated ZEROPS sdk

package enum

type ProcessExecutorTagEnum string

const (
	ProcessExecutorTagEnumMain     = ProcessExecutorTagEnum("MAIN")
	ProcessExecutorTagEnumL7Master = ProcessExecutorTagEnum("L7_MASTER")
)

func NewProcessExecutorTagEnumFromString(value string) (out ProcessExecutorTagEnum, err error) {
	return ProcessExecutorTagEnum(value), nil
}

func (enum ProcessExecutorTagEnum) String() string {
	return string(enum)
}

func (enum ProcessExecutorTagEnum) Native() string {
	return string(enum)
}

func (enum ProcessExecutorTagEnum) Values() []ProcessExecutorTagEnum {
	return ProcessExecutorTagEnumAll()
}

func (enum ProcessExecutorTagEnum) PublicValues() []ProcessExecutorTagEnum {
	return ProcessExecutorTagEnumAllPublic()
}

func (enum ProcessExecutorTagEnum) PrivateValues() []ProcessExecutorTagEnum {
	return ProcessExecutorTagEnumAllPrivate()
}

func (enum ProcessExecutorTagEnum) DefaultValue() ProcessExecutorTagEnum {
	return ProcessExecutorTagEnumDefault()
}

func (enum ProcessExecutorTagEnum) Is(values ...ProcessExecutorTagEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ProcessExecutorTagEnumAllStrings() []string {
	return []string{
		string(ProcessExecutorTagEnumMain), string(ProcessExecutorTagEnumL7Master),
	}
}

func ProcessExecutorTagEnumAll() []ProcessExecutorTagEnum {
	return []ProcessExecutorTagEnum{
		ProcessExecutorTagEnumMain, ProcessExecutorTagEnumL7Master,
	}
}

func ProcessExecutorTagEnumAllPublic() []ProcessExecutorTagEnum {
	return []ProcessExecutorTagEnum{
		ProcessExecutorTagEnumMain, ProcessExecutorTagEnumL7Master,
	}
}

func ProcessExecutorTagEnumAllPrivate() []ProcessExecutorTagEnum {
	return []ProcessExecutorTagEnum{}
}

func ProcessExecutorTagEnumDefault() ProcessExecutorTagEnum {
	return ""
}

func (enum ProcessExecutorTagEnum) IsMain() bool {
	return enum.Is(ProcessExecutorTagEnumMain)
}

func (enum ProcessExecutorTagEnum) IsL7Master() bool {
	return enum.Is(ProcessExecutorTagEnumL7Master)
}
