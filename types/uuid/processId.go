// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ProcessId types.UuidShort

func NewProcessIdFromString(value string) (out ProcessId, err error) {
	return ProcessId(value), nil
}

func (parameter ProcessId) Native() string {
	return string(parameter)
}

func (parameter ProcessId) TypedString() types.String {
	return types.String(parameter)
}

type ProcessIdNull struct {
	value  ProcessId
	filled bool
}

func (parameter ProcessId) ProcessIdNull() ProcessIdNull {
	return NewProcessIdNull(parameter)
}

func NewProcessIdNull(value ProcessId) ProcessIdNull {
	return ProcessIdNull{
		filled: true,
		value:  value,
	}
}

func NewProcessIdNullFromString(value string) ProcessIdNull {
	return ProcessIdNull{
		filled: true,
		value:  ProcessId(value),
	}
}

func (parameter ProcessIdNull) Get() (ProcessId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ProcessIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ProcessIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ProcessIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
