// Generated ZEROPS sdk

package enum

type AuthorAuthorTypeEnum string

const (
	AuthorAuthorTypeEnumClient     = AuthorAuthorTypeEnum("CLIENT")
	AuthorAuthorTypeEnumBackoffice = AuthorAuthorTypeEnum("BACKOFFICE")
)

func NewAuthorAuthorTypeEnumFromString(value string) (out AuthorAuthorTypeEnum, err error) {
	return AuthorAuthorTypeEnum(value), nil
}

func (enum AuthorAuthorTypeEnum) String() string {
	return string(enum)
}

func (enum AuthorAuthorTypeEnum) Native() string {
	return string(enum)
}

func (enum AuthorAuthorTypeEnum) Values() []AuthorAuthorTypeEnum {
	return AuthorAuthorTypeEnumAll()
}

func (enum AuthorAuthorTypeEnum) PublicValues() []AuthorAuthorTypeEnum {
	return AuthorAuthorTypeEnumAllPublic()
}

func (enum AuthorAuthorTypeEnum) PrivateValues() []AuthorAuthorTypeEnum {
	return AuthorAuthorTypeEnumAllPrivate()
}

func (enum AuthorAuthorTypeEnum) DefaultValue() AuthorAuthorTypeEnum {
	return AuthorAuthorTypeEnumDefault()
}

func (enum AuthorAuthorTypeEnum) Is(values ...AuthorAuthorTypeEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func AuthorAuthorTypeEnumAllStrings() []string {
	return []string{
		string(AuthorAuthorTypeEnumClient), string(AuthorAuthorTypeEnumBackoffice),
	}
}

func AuthorAuthorTypeEnumAll() []AuthorAuthorTypeEnum {
	return []AuthorAuthorTypeEnum{
		AuthorAuthorTypeEnumClient, AuthorAuthorTypeEnumBackoffice,
	}
}

func AuthorAuthorTypeEnumAllPublic() []AuthorAuthorTypeEnum {
	return []AuthorAuthorTypeEnum{}
}

func AuthorAuthorTypeEnumAllPrivate() []AuthorAuthorTypeEnum {
	return []AuthorAuthorTypeEnum{
		AuthorAuthorTypeEnumClient, AuthorAuthorTypeEnumBackoffice,
	}
}

func AuthorAuthorTypeEnumDefault() AuthorAuthorTypeEnum {
	return AuthorAuthorTypeEnumClient
}

func (enum AuthorAuthorTypeEnum) IsClient() bool {
	return enum.Is(AuthorAuthorTypeEnumClient)
}

func (enum AuthorAuthorTypeEnum) IsBackoffice() bool {
	return enum.Is(AuthorAuthorTypeEnumBackoffice)
}
