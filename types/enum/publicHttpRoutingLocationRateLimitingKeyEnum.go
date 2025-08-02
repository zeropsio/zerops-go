// Generated ZEROPS sdk

package enum

type PublicHttpRoutingLocationRateLimitingKeyEnum string

const (
	PublicHttpRoutingLocationRateLimitingKeyEnumBinaryRemoteAddr = PublicHttpRoutingLocationRateLimitingKeyEnum("binary_remote_addr")
	PublicHttpRoutingLocationRateLimitingKeyEnumServerName       = PublicHttpRoutingLocationRateLimitingKeyEnum("server_name")
)

func NewPublicHttpRoutingLocationRateLimitingKeyEnumFromString(value string) (out PublicHttpRoutingLocationRateLimitingKeyEnum, err error) {
	return PublicHttpRoutingLocationRateLimitingKeyEnum(value), nil
}

func (enum PublicHttpRoutingLocationRateLimitingKeyEnum) String() string {
	return string(enum)
}

func (enum PublicHttpRoutingLocationRateLimitingKeyEnum) Native() string {
	return string(enum)
}

func (enum PublicHttpRoutingLocationRateLimitingKeyEnum) Values() []PublicHttpRoutingLocationRateLimitingKeyEnum {
	return PublicHttpRoutingLocationRateLimitingKeyEnumAll()
}

func (enum PublicHttpRoutingLocationRateLimitingKeyEnum) PublicValues() []PublicHttpRoutingLocationRateLimitingKeyEnum {
	return PublicHttpRoutingLocationRateLimitingKeyEnumAllPublic()
}

func (enum PublicHttpRoutingLocationRateLimitingKeyEnum) PrivateValues() []PublicHttpRoutingLocationRateLimitingKeyEnum {
	return PublicHttpRoutingLocationRateLimitingKeyEnumAllPrivate()
}

func (enum PublicHttpRoutingLocationRateLimitingKeyEnum) DefaultValue() PublicHttpRoutingLocationRateLimitingKeyEnum {
	return PublicHttpRoutingLocationRateLimitingKeyEnumDefault()
}

func (enum PublicHttpRoutingLocationRateLimitingKeyEnum) Is(values ...PublicHttpRoutingLocationRateLimitingKeyEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PublicHttpRoutingLocationRateLimitingKeyEnumAllStrings() []string {
	return []string{
		string(PublicHttpRoutingLocationRateLimitingKeyEnumBinaryRemoteAddr), string(PublicHttpRoutingLocationRateLimitingKeyEnumServerName),
	}
}

func PublicHttpRoutingLocationRateLimitingKeyEnumAll() []PublicHttpRoutingLocationRateLimitingKeyEnum {
	return []PublicHttpRoutingLocationRateLimitingKeyEnum{
		PublicHttpRoutingLocationRateLimitingKeyEnumBinaryRemoteAddr, PublicHttpRoutingLocationRateLimitingKeyEnumServerName,
	}
}

func PublicHttpRoutingLocationRateLimitingKeyEnumAllPublic() []PublicHttpRoutingLocationRateLimitingKeyEnum {
	return []PublicHttpRoutingLocationRateLimitingKeyEnum{
		PublicHttpRoutingLocationRateLimitingKeyEnumBinaryRemoteAddr, PublicHttpRoutingLocationRateLimitingKeyEnumServerName,
	}
}

func PublicHttpRoutingLocationRateLimitingKeyEnumAllPrivate() []PublicHttpRoutingLocationRateLimitingKeyEnum {
	return []PublicHttpRoutingLocationRateLimitingKeyEnum{}
}

func PublicHttpRoutingLocationRateLimitingKeyEnumDefault() PublicHttpRoutingLocationRateLimitingKeyEnum {
	return PublicHttpRoutingLocationRateLimitingKeyEnumBinaryRemoteAddr
}

func (enum PublicHttpRoutingLocationRateLimitingKeyEnum) IsBinaryRemoteAddr() bool {
	return enum.Is(PublicHttpRoutingLocationRateLimitingKeyEnumBinaryRemoteAddr)
}

func (enum PublicHttpRoutingLocationRateLimitingKeyEnum) IsServerName() bool {
	return enum.Is(PublicHttpRoutingLocationRateLimitingKeyEnumServerName)
}
