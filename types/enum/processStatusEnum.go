// Generated ZEROPS sdk

package enum

type ProcessStatusEnum string

const (
	ProcessStatusEnumPending     = ProcessStatusEnum("PENDING")
	ProcessStatusEnumRunning     = ProcessStatusEnum("RUNNING")
	ProcessStatusEnumRollbacking = ProcessStatusEnum("ROLLBACKING")
	ProcessStatusEnumCanceling   = ProcessStatusEnum("CANCELING")
	ProcessStatusEnumFinished    = ProcessStatusEnum("FINISHED")
	ProcessStatusEnumFailed      = ProcessStatusEnum("FAILED")
	ProcessStatusEnumCanceled    = ProcessStatusEnum("CANCELED")
)

func NewProcessStatusEnumFromString(value string) (out ProcessStatusEnum, err error) {
	return ProcessStatusEnum(value), nil
}

func (enum ProcessStatusEnum) String() string {
	return string(enum)
}

func (enum ProcessStatusEnum) Native() string {
	return string(enum)
}

func (enum ProcessStatusEnum) Values() []ProcessStatusEnum {
	return ProcessStatusEnumAll()
}

func (enum ProcessStatusEnum) PublicValues() []ProcessStatusEnum {
	return ProcessStatusEnumAllPublic()
}

func (enum ProcessStatusEnum) PrivateValues() []ProcessStatusEnum {
	return ProcessStatusEnumAllPrivate()
}

func (enum ProcessStatusEnum) DefaultValue() ProcessStatusEnum {
	return ProcessStatusEnumDefault()
}

func (enum ProcessStatusEnum) Is(values ...ProcessStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ProcessStatusEnumAllStrings() []string {
	return []string{
		string(ProcessStatusEnumPending), string(ProcessStatusEnumRunning), string(ProcessStatusEnumRollbacking), string(ProcessStatusEnumCanceling), string(ProcessStatusEnumFinished), string(ProcessStatusEnumFailed), string(ProcessStatusEnumCanceled),
	}
}

func ProcessStatusEnumAll() []ProcessStatusEnum {
	return []ProcessStatusEnum{
		ProcessStatusEnumPending, ProcessStatusEnumRunning, ProcessStatusEnumRollbacking, ProcessStatusEnumCanceling, ProcessStatusEnumFinished, ProcessStatusEnumFailed, ProcessStatusEnumCanceled,
	}
}

func ProcessStatusEnumAllPublic() []ProcessStatusEnum {
	return []ProcessStatusEnum{}
}

func ProcessStatusEnumAllPrivate() []ProcessStatusEnum {
	return []ProcessStatusEnum{
		ProcessStatusEnumPending, ProcessStatusEnumRunning, ProcessStatusEnumRollbacking, ProcessStatusEnumCanceling, ProcessStatusEnumFinished, ProcessStatusEnumFailed, ProcessStatusEnumCanceled,
	}
}

func ProcessStatusEnumDefault() ProcessStatusEnum {
	return ProcessStatusEnumPending
}

func (enum ProcessStatusEnum) IsPending() bool {
	return enum.Is(ProcessStatusEnumPending)
}

func (enum ProcessStatusEnum) IsRunning() bool {
	return enum.Is(ProcessStatusEnumRunning)
}

func (enum ProcessStatusEnum) IsRollbacking() bool {
	return enum.Is(ProcessStatusEnumRollbacking)
}

func (enum ProcessStatusEnum) IsCanceling() bool {
	return enum.Is(ProcessStatusEnumCanceling)
}

func (enum ProcessStatusEnum) IsFinished() bool {
	return enum.Is(ProcessStatusEnumFinished)
}

func (enum ProcessStatusEnum) IsFailed() bool {
	return enum.Is(ProcessStatusEnumFailed)
}

func (enum ProcessStatusEnum) IsCanceled() bool {
	return enum.Is(ProcessStatusEnumCanceled)
}
