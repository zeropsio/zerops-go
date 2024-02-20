// Generated ZEROPS sdk

package stringId

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type LanguageId types.StringId

func (parameter LanguageId) LanguageIdNull() LanguageIdNull {
	return NewLanguageIdNull(parameter)
}

func NewLanguageIdNullFromString(value string) LanguageIdNull {
	return LanguageIdNull{
		filled: true,
		value:  LanguageId(value),
	}
}

func NewLanguageIdNull(value LanguageId) LanguageIdNull {
	return LanguageIdNull{
		filled: true,
		value:  value,
	}
}

func NewLanguageIdFromString(value string) (out LanguageId, err error) {
	return LanguageId(value), nil
}

func (parameter LanguageId) Native() string {
	return string(parameter)
}

type LanguageIdNull struct {
	value  LanguageId
	filled bool
}

func (parameter LanguageIdNull) Get() (LanguageId, bool) {
	return parameter.value, parameter.filled
}

func (parameter LanguageIdNull) Filled() bool {
	return parameter.filled
}

func (parameter LanguageIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *LanguageIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
