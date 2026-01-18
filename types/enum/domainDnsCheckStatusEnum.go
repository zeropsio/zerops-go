// Generated ZEROPS sdk

package enum

type DomainDnsCheckStatusEnum string

const (
	DomainDnsCheckStatusEnumPending  = DomainDnsCheckStatusEnum("PENDING")
	DomainDnsCheckStatusEnumChecking = DomainDnsCheckStatusEnum("CHECKING")
	DomainDnsCheckStatusEnumOk       = DomainDnsCheckStatusEnum("OK")
	DomainDnsCheckStatusEnumFailed   = DomainDnsCheckStatusEnum("FAILED")
)

func NewDomainDnsCheckStatusEnumFromString(value string) (out DomainDnsCheckStatusEnum, err error) {
	return DomainDnsCheckStatusEnum(value), nil
}

func (enum DomainDnsCheckStatusEnum) String() string {
	return string(enum)
}

func (enum DomainDnsCheckStatusEnum) Native() string {
	return string(enum)
}

func (enum DomainDnsCheckStatusEnum) Values() []DomainDnsCheckStatusEnum {
	return DomainDnsCheckStatusEnumAll()
}

func (enum DomainDnsCheckStatusEnum) PublicValues() []DomainDnsCheckStatusEnum {
	return DomainDnsCheckStatusEnumAllPublic()
}

func (enum DomainDnsCheckStatusEnum) PrivateValues() []DomainDnsCheckStatusEnum {
	return DomainDnsCheckStatusEnumAllPrivate()
}

func (enum DomainDnsCheckStatusEnum) DefaultValue() DomainDnsCheckStatusEnum {
	return DomainDnsCheckStatusEnumDefault()
}

func (enum DomainDnsCheckStatusEnum) Is(values ...DomainDnsCheckStatusEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func DomainDnsCheckStatusEnumAllStrings() []string {
	return []string{
		string(DomainDnsCheckStatusEnumPending), string(DomainDnsCheckStatusEnumChecking), string(DomainDnsCheckStatusEnumOk), string(DomainDnsCheckStatusEnumFailed),
	}
}

func DomainDnsCheckStatusEnumAll() []DomainDnsCheckStatusEnum {
	return []DomainDnsCheckStatusEnum{
		DomainDnsCheckStatusEnumPending, DomainDnsCheckStatusEnumChecking, DomainDnsCheckStatusEnumOk, DomainDnsCheckStatusEnumFailed,
	}
}

func DomainDnsCheckStatusEnumAllPublic() []DomainDnsCheckStatusEnum {
	return []DomainDnsCheckStatusEnum{
		DomainDnsCheckStatusEnumPending, DomainDnsCheckStatusEnumChecking, DomainDnsCheckStatusEnumOk, DomainDnsCheckStatusEnumFailed,
	}
}

func DomainDnsCheckStatusEnumAllPrivate() []DomainDnsCheckStatusEnum {
	return []DomainDnsCheckStatusEnum{}
}

func DomainDnsCheckStatusEnumDefault() DomainDnsCheckStatusEnum {
	return DomainDnsCheckStatusEnumPending
}

func (enum DomainDnsCheckStatusEnum) IsPending() bool {
	return enum.Is(DomainDnsCheckStatusEnumPending)
}

func (enum DomainDnsCheckStatusEnum) IsChecking() bool {
	return enum.Is(DomainDnsCheckStatusEnumChecking)
}

func (enum DomainDnsCheckStatusEnum) IsOk() bool {
	return enum.Is(DomainDnsCheckStatusEnumOk)
}

func (enum DomainDnsCheckStatusEnum) IsFailed() bool {
	return enum.Is(DomainDnsCheckStatusEnumFailed)
}
