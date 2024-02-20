// Generated ZEROPS sdk

package stringId

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type CountryId types.StringId

func (parameter CountryId) CountryIdNull() CountryIdNull {
	return NewCountryIdNull(parameter)
}

func NewCountryIdNullFromString(value string) CountryIdNull {
	return CountryIdNull{
		filled: true,
		value:  CountryId(value),
	}
}

func NewCountryIdNull(value CountryId) CountryIdNull {
	return CountryIdNull{
		filled: true,
		value:  value,
	}
}

func NewCountryIdFromString(value string) (out CountryId, err error) {
	return CountryId(value), nil
}

func (parameter CountryId) Native() string {
	return string(parameter)
}

type CountryIdNull struct {
	value  CountryId
	filled bool
}

func (parameter CountryIdNull) Get() (CountryId, bool) {
	return parameter.value, parameter.filled
}

func (parameter CountryIdNull) Filled() bool {
	return parameter.filled
}

func (parameter CountryIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *CountryIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
