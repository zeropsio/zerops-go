// Generated ZEROPS sdk

package enum

type PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum string

const (
	PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllow = PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum("allow")
	PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumDeny  = PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum("deny")
)

func NewPublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumFromString(value string) (out PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum, err error) {
	return PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum(value), nil
}

func (enum PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) String() string {
	return string(enum)
}

func (enum PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) Native() string {
	return string(enum)
}

func (enum PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) Values() []PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum {
	return PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAll()
}

func (enum PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) PublicValues() []PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum {
	return PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllPublic()
}

func (enum PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) PrivateValues() []PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum {
	return PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllPrivate()
}

func (enum PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) DefaultValue() PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum {
	return PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumDefault()
}

func (enum PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) Is(values ...PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllStrings() []string {
	return []string{
		string(PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllow), string(PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumDeny),
	}
}

func PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAll() []PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum {
	return []PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum{
		PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllow, PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumDeny,
	}
}

func PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllPublic() []PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum {
	return []PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum{
		PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllow, PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumDeny,
	}
}

func PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllPrivate() []PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum {
	return []PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum{}
}

func PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumDefault() PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum {
	return PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllow
}

func (enum PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) IsAllow() bool {
	return enum.Is(PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumAllow)
}

func (enum PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnum) IsDeny() bool {
	return enum.Is(PublicHttpRoutingLocationAccessPolicyDefaultPolicyEnumDeny)
}
