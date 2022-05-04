// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type AppVersionId types.UuidShort

func NewAppVersionIdFromString(value string) (out AppVersionId, err error) {
	return AppVersionId(value), nil
}

func (parameter AppVersionId) Native() string {
	return string(parameter)
}

func (parameter AppVersionId) TypedString() types.String {
	return types.String(parameter)
}

type AppVersionIdNull struct {
	value  AppVersionId
	filled bool
}

func (parameter AppVersionId) AppVersionIdNull() AppVersionIdNull {
	return NewAppVersionIdNull(parameter)
}

func NewAppVersionIdNull(value AppVersionId) AppVersionIdNull {
	return AppVersionIdNull{
		filled: true,
		value:  value,
	}
}

func NewAppVersionIdNullFromString(value string) AppVersionIdNull {
	return AppVersionIdNull{
		filled: true,
		value:  AppVersionId(value),
	}
}

func (parameter AppVersionIdNull) Get() (AppVersionId, bool) {
	return parameter.value, parameter.filled
}

func (parameter AppVersionIdNull) Filled() bool {
	return parameter.filled
}

func (parameter AppVersionIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *AppVersionIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
