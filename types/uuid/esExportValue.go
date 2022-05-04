// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type EsExportValue types.UuidShort

func NewEsExportValueFromString(value string) (out EsExportValue, err error) {
	return EsExportValue(value), nil
}

func (parameter EsExportValue) Native() string {
	return string(parameter)
}

func (parameter EsExportValue) TypedString() types.String {
	return types.String(parameter)
}

type EsExportValueNull struct {
	value  EsExportValue
	filled bool
}

func (parameter EsExportValue) EsExportValueNull() EsExportValueNull {
	return NewEsExportValueNull(parameter)
}

func NewEsExportValueNull(value EsExportValue) EsExportValueNull {
	return EsExportValueNull{
		filled: true,
		value:  value,
	}
}

func NewEsExportValueNullFromString(value string) EsExportValueNull {
	return EsExportValueNull{
		filled: true,
		value:  EsExportValue(value),
	}
}

func (parameter EsExportValueNull) Get() (EsExportValue, bool) {
	return parameter.value, parameter.filled
}

func (parameter EsExportValueNull) Filled() bool {
	return parameter.filled
}

func (parameter EsExportValueNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *EsExportValueNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
