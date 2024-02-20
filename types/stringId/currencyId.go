// Generated ZEROPS sdk

package stringId

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type CurrencyId types.StringId

func (parameter CurrencyId) CurrencyIdNull() CurrencyIdNull {
	return NewCurrencyIdNull(parameter)
}

func NewCurrencyIdNullFromString(value string) CurrencyIdNull {
	return CurrencyIdNull{
		filled: true,
		value:  CurrencyId(value),
	}
}

func NewCurrencyIdNull(value CurrencyId) CurrencyIdNull {
	return CurrencyIdNull{
		filled: true,
		value:  value,
	}
}

func NewCurrencyIdFromString(value string) (out CurrencyId, err error) {
	return CurrencyId(value), nil
}

func (parameter CurrencyId) Native() string {
	return string(parameter)
}

type CurrencyIdNull struct {
	value  CurrencyId
	filled bool
}

func (parameter CurrencyIdNull) Get() (CurrencyId, bool) {
	return parameter.value, parameter.filled
}

func (parameter CurrencyIdNull) Filled() bool {
	return parameter.filled
}

func (parameter CurrencyIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *CurrencyIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
