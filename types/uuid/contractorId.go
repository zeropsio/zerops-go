// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ContractorId types.UuidShort

func NewContractorIdFromString(value string) (out ContractorId, err error) {
	return ContractorId(value), nil
}

func (parameter ContractorId) Native() string {
	return string(parameter)
}

func (parameter ContractorId) TypedString() types.String {
	return types.String(parameter)
}

type ContractorIdNull struct {
	value  ContractorId
	filled bool
}

func (parameter ContractorId) ContractorIdNull() ContractorIdNull {
	return NewContractorIdNull(parameter)
}

func NewContractorIdNull(value ContractorId) ContractorIdNull {
	return ContractorIdNull{
		filled: true,
		value:  value,
	}
}

func NewContractorIdNullFromString(value string) ContractorIdNull {
	return ContractorIdNull{
		filled: true,
		value:  ContractorId(value),
	}
}

func (parameter ContractorIdNull) Get() (ContractorId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ContractorIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ContractorIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ContractorIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
