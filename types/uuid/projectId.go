// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ProjectId types.UuidShort

func NewProjectIdFromString(value string) (out ProjectId, err error) {
	return ProjectId(value), nil
}

func (parameter ProjectId) Native() string {
	return string(parameter)
}

func (parameter ProjectId) TypedString() types.String {
	return types.String(parameter)
}

type ProjectIdNull struct {
	value  ProjectId
	filled bool
}

func (parameter ProjectId) ProjectIdNull() ProjectIdNull {
	return NewProjectIdNull(parameter)
}

func NewProjectIdNull(value ProjectId) ProjectIdNull {
	return ProjectIdNull{
		filled: true,
		value:  value,
	}
}

func NewProjectIdNullFromString(value string) ProjectIdNull {
	return ProjectIdNull{
		filled: true,
		value:  ProjectId(value),
	}
}

func (parameter ProjectIdNull) Get() (ProjectId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ProjectIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ProjectIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ProjectIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
