// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type ObjectStorageId types.UuidShort

func NewObjectStorageIdFromString(value string) (out ObjectStorageId, err error) {
	return ObjectStorageId(value), nil
}

func (parameter ObjectStorageId) Native() string {
	return string(parameter)
}

func (parameter ObjectStorageId) TypedString() types.String {
	return types.String(parameter)
}

type ObjectStorageIdNull struct {
	value  ObjectStorageId
	filled bool
}

func (parameter ObjectStorageId) ObjectStorageIdNull() ObjectStorageIdNull {
	return NewObjectStorageIdNull(parameter)
}

func NewObjectStorageIdNull(value ObjectStorageId) ObjectStorageIdNull {
	return ObjectStorageIdNull{
		filled: true,
		value:  value,
	}
}

func NewObjectStorageIdNullFromString(value string) ObjectStorageIdNull {
	return ObjectStorageIdNull{
		filled: true,
		value:  ObjectStorageId(value),
	}
}

func (parameter ObjectStorageIdNull) Get() (ObjectStorageId, bool) {
	return parameter.value, parameter.filled
}

func (parameter ObjectStorageIdNull) Filled() bool {
	return parameter.filled
}

func (parameter ObjectStorageIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *ObjectStorageIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
