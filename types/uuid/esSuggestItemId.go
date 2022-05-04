// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type EsSuggestItemId types.UuidShort

func NewEsSuggestItemIdFromString(value string) (out EsSuggestItemId, err error) {
	return EsSuggestItemId(value), nil
}

func (parameter EsSuggestItemId) Native() string {
	return string(parameter)
}

func (parameter EsSuggestItemId) TypedString() types.String {
	return types.String(parameter)
}

type EsSuggestItemIdNull struct {
	value  EsSuggestItemId
	filled bool
}

func (parameter EsSuggestItemId) EsSuggestItemIdNull() EsSuggestItemIdNull {
	return NewEsSuggestItemIdNull(parameter)
}

func NewEsSuggestItemIdNull(value EsSuggestItemId) EsSuggestItemIdNull {
	return EsSuggestItemIdNull{
		filled: true,
		value:  value,
	}
}

func NewEsSuggestItemIdNullFromString(value string) EsSuggestItemIdNull {
	return EsSuggestItemIdNull{
		filled: true,
		value:  EsSuggestItemId(value),
	}
}

func (parameter EsSuggestItemIdNull) Get() (EsSuggestItemId, bool) {
	return parameter.value, parameter.filled
}

func (parameter EsSuggestItemIdNull) Filled() bool {
	return parameter.filled
}

func (parameter EsSuggestItemIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *EsSuggestItemIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
