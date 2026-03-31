// Generated ZEROPS sdk

package stringId

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type LocationId types.StringId

func (parameter LocationId) LocationIdNull() LocationIdNull {
	return NewLocationIdNull(parameter)
}

func NewLocationIdNullFromString(value string) LocationIdNull {
	return LocationIdNull{
		filled: true,
		value:  LocationId(value),
	}
}

func NewLocationIdNull(value LocationId) LocationIdNull {
	return LocationIdNull{
		filled: true,
		value:  value,
	}
}

func NewLocationIdFromString(value string) (out LocationId, err error) {
	return LocationId(value), nil
}

func (parameter LocationId) Native() string {
	return string(parameter)
}

type LocationIdNull struct {
	value  LocationId
	filled bool
}

func (parameter LocationIdNull) Get() (LocationId, bool) {
	return parameter.value, parameter.filled
}

func (parameter LocationIdNull) Filled() bool {
	return parameter.filled
}

func (parameter LocationIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *LocationIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
