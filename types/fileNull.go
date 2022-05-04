// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

var _ json.Unmarshaler

type FileNull struct {
	value  File
	filled bool
}

func NewFileNull(value string) FileNull {
	return FileNull{
		value:  NewFile(value),
		filled: true,
	}
}

func (parameter FileNull) Get() (File, bool) {
	return parameter.value, parameter.filled
}

func (parameter FileNull) Filled() bool {
	return parameter.filled
}

func (parameter FileNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		bytes, err := json.Marshal(parameter.value)
		return bytes, err
	}

	return []byte("null"), nil
}

func (parameter *FileNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
