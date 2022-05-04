// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"time"
)

var _ json.Unmarshaler

type DateTimeNull struct {
	value  DateTime
	filled bool
}

func NewDateTimeNull(value time.Time) DateTimeNull {
	return DateTimeNull{
		value:  NewDateTime(value),
		filled: true,
	}
}

func (parameter DateTimeNull) Get() (DateTime, bool) {
	return parameter.value, parameter.filled
}

func (parameter DateTimeNull) Filled() bool {
	return parameter.filled
}

func (parameter DateTimeNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *DateTimeNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
