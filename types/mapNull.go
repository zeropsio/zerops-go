// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

type TypeScope int

const (
	TypeScopeInput = TypeScope(iota)
	TypeScopeOutput
	TypeScopePath
	TypeScopeQuery
)

type Base struct {
	value  interface{}
	filled bool
}

func (parameter Base) Filled() bool {
	return parameter.filled
}

func (parameter Base) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *Base) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
