// Generated ZEROPS sdk

package enum

type PublicPortRoutingFirewallPolicyEnum string

const (
	PublicPortRoutingFirewallPolicyEnumBlacklist = PublicPortRoutingFirewallPolicyEnum("BLACKLIST")
	PublicPortRoutingFirewallPolicyEnumWhitelist = PublicPortRoutingFirewallPolicyEnum("WHITELIST")
)

func NewPublicPortRoutingFirewallPolicyEnumFromString(value string) (out PublicPortRoutingFirewallPolicyEnum, err error) {
	return PublicPortRoutingFirewallPolicyEnum(value), nil
}

func (enum PublicPortRoutingFirewallPolicyEnum) String() string {
	return string(enum)
}

func (enum PublicPortRoutingFirewallPolicyEnum) Native() string {
	return string(enum)
}

func (enum PublicPortRoutingFirewallPolicyEnum) Values() []PublicPortRoutingFirewallPolicyEnum {
	return PublicPortRoutingFirewallPolicyEnumAll()
}

func (enum PublicPortRoutingFirewallPolicyEnum) PublicValues() []PublicPortRoutingFirewallPolicyEnum {
	return PublicPortRoutingFirewallPolicyEnumAllPublic()
}

func (enum PublicPortRoutingFirewallPolicyEnum) PrivateValues() []PublicPortRoutingFirewallPolicyEnum {
	return PublicPortRoutingFirewallPolicyEnumAllPrivate()
}

func (enum PublicPortRoutingFirewallPolicyEnum) DefaultValue() PublicPortRoutingFirewallPolicyEnum {
	return PublicPortRoutingFirewallPolicyEnumDefault()
}

func (enum PublicPortRoutingFirewallPolicyEnum) Is(values ...PublicPortRoutingFirewallPolicyEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func PublicPortRoutingFirewallPolicyEnumAllStrings() []string {
	return []string{
		string(PublicPortRoutingFirewallPolicyEnumBlacklist), string(PublicPortRoutingFirewallPolicyEnumWhitelist),
	}
}

func PublicPortRoutingFirewallPolicyEnumAll() []PublicPortRoutingFirewallPolicyEnum {
	return []PublicPortRoutingFirewallPolicyEnum{
		PublicPortRoutingFirewallPolicyEnumBlacklist, PublicPortRoutingFirewallPolicyEnumWhitelist,
	}
}

func PublicPortRoutingFirewallPolicyEnumAllPublic() []PublicPortRoutingFirewallPolicyEnum {
	return []PublicPortRoutingFirewallPolicyEnum{
		PublicPortRoutingFirewallPolicyEnumBlacklist, PublicPortRoutingFirewallPolicyEnumWhitelist,
	}
}

func PublicPortRoutingFirewallPolicyEnumAllPrivate() []PublicPortRoutingFirewallPolicyEnum {
	return []PublicPortRoutingFirewallPolicyEnum{}
}

func PublicPortRoutingFirewallPolicyEnumDefault() PublicPortRoutingFirewallPolicyEnum {
	return PublicPortRoutingFirewallPolicyEnumBlacklist
}

func (enum PublicPortRoutingFirewallPolicyEnum) IsBlacklist() bool {
	return enum.Is(PublicPortRoutingFirewallPolicyEnumBlacklist)
}

func (enum PublicPortRoutingFirewallPolicyEnum) IsWhitelist() bool {
	return enum.Is(PublicPortRoutingFirewallPolicyEnumWhitelist)
}
