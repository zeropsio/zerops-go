// Generated ZEROPS sdk

package enum

type ClientWelcomeCreditWithheldEnum string

const (
	ClientWelcomeCreditWithheldEnumDatacenterIp = ClientWelcomeCreditWithheldEnum("DATACENTER_IP")
)

func NewClientWelcomeCreditWithheldEnumFromString(value string) (out ClientWelcomeCreditWithheldEnum, err error) {
	return ClientWelcomeCreditWithheldEnum(value), nil
}

func (enum ClientWelcomeCreditWithheldEnum) String() string {
	return string(enum)
}

func (enum ClientWelcomeCreditWithheldEnum) Native() string {
	return string(enum)
}

func (enum ClientWelcomeCreditWithheldEnum) Values() []ClientWelcomeCreditWithheldEnum {
	return ClientWelcomeCreditWithheldEnumAll()
}

func (enum ClientWelcomeCreditWithheldEnum) PublicValues() []ClientWelcomeCreditWithheldEnum {
	return ClientWelcomeCreditWithheldEnumAllPublic()
}

func (enum ClientWelcomeCreditWithheldEnum) PrivateValues() []ClientWelcomeCreditWithheldEnum {
	return ClientWelcomeCreditWithheldEnumAllPrivate()
}

func (enum ClientWelcomeCreditWithheldEnum) DefaultValue() ClientWelcomeCreditWithheldEnum {
	return ClientWelcomeCreditWithheldEnumDefault()
}

func (enum ClientWelcomeCreditWithheldEnum) Is(values ...ClientWelcomeCreditWithheldEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func ClientWelcomeCreditWithheldEnumAllStrings() []string {
	return []string{
		string(ClientWelcomeCreditWithheldEnumDatacenterIp),
	}
}

func ClientWelcomeCreditWithheldEnumAll() []ClientWelcomeCreditWithheldEnum {
	return []ClientWelcomeCreditWithheldEnum{
		ClientWelcomeCreditWithheldEnumDatacenterIp,
	}
}

func ClientWelcomeCreditWithheldEnumAllPublic() []ClientWelcomeCreditWithheldEnum {
	return []ClientWelcomeCreditWithheldEnum{
		ClientWelcomeCreditWithheldEnumDatacenterIp,
	}
}

func ClientWelcomeCreditWithheldEnumAllPrivate() []ClientWelcomeCreditWithheldEnum {
	return []ClientWelcomeCreditWithheldEnum{}
}

func ClientWelcomeCreditWithheldEnumDefault() ClientWelcomeCreditWithheldEnum {
	return ""
}

func (enum ClientWelcomeCreditWithheldEnum) IsDatacenterIp() bool {
	return enum.Is(ClientWelcomeCreditWithheldEnumDatacenterIp)
}
