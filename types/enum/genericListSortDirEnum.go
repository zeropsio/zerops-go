// Generated ZEROPS sdk

package enum

type GenericListSortDirEnum string

const (
	GenericListSortDirEnumAsc  = GenericListSortDirEnum("asc")
	GenericListSortDirEnumDesc = GenericListSortDirEnum("desc")
)

func NewGenericListSortDirEnumFromString(value string) (out GenericListSortDirEnum, err error) {
	return GenericListSortDirEnum(value), nil
}

func (enum GenericListSortDirEnum) String() string {
	return string(enum)
}

func (enum GenericListSortDirEnum) Native() string {
	return string(enum)
}

func (enum GenericListSortDirEnum) Values() []GenericListSortDirEnum {
	return GenericListSortDirEnumAll()
}

func (enum GenericListSortDirEnum) PublicValues() []GenericListSortDirEnum {
	return GenericListSortDirEnumAllPublic()
}

func (enum GenericListSortDirEnum) PrivateValues() []GenericListSortDirEnum {
	return GenericListSortDirEnumAllPrivate()
}

func (enum GenericListSortDirEnum) DefaultValue() GenericListSortDirEnum {
	return GenericListSortDirEnumDefault()
}

func (enum GenericListSortDirEnum) Is(values ...GenericListSortDirEnum) bool {
	for _, value := range values {
		if enum == value {
			return true
		}
	}
	return false
}

func GenericListSortDirEnumAllStrings() []string {
	return []string{
		string(GenericListSortDirEnumAsc), string(GenericListSortDirEnumDesc),
	}
}

func GenericListSortDirEnumAll() []GenericListSortDirEnum {
	return []GenericListSortDirEnum{
		GenericListSortDirEnumAsc, GenericListSortDirEnumDesc,
	}
}

func GenericListSortDirEnumAllPublic() []GenericListSortDirEnum {
	return []GenericListSortDirEnum{
		GenericListSortDirEnumAsc, GenericListSortDirEnumDesc,
	}
}

func GenericListSortDirEnumAllPrivate() []GenericListSortDirEnum {
	return []GenericListSortDirEnum{}
}

func GenericListSortDirEnumDefault() GenericListSortDirEnum {
	return GenericListSortDirEnumDesc
}

func (enum GenericListSortDirEnum) IsAsc() bool {
	return enum.Is(GenericListSortDirEnumAsc)
}

func (enum GenericListSortDirEnum) IsDesc() bool {
	return enum.Is(GenericListSortDirEnumDesc)
}
