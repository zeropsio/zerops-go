// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type DeviceLoginRequestId types.UuidShort

func NewDeviceLoginRequestIdFromString(value string) (out DeviceLoginRequestId, err error) {
	return DeviceLoginRequestId(value), nil
}

func (parameter DeviceLoginRequestId) Native() string {
	return string(parameter)
}

func (parameter DeviceLoginRequestId) TypedString() types.String {
	return types.String(parameter)
}

type DeviceLoginRequestIdNull struct {
	value  DeviceLoginRequestId
	filled bool
}

func (parameter DeviceLoginRequestId) DeviceLoginRequestIdNull() DeviceLoginRequestIdNull {
	return NewDeviceLoginRequestIdNull(parameter)
}

func NewDeviceLoginRequestIdNull(value DeviceLoginRequestId) DeviceLoginRequestIdNull {
	return DeviceLoginRequestIdNull{
		filled: true,
		value:  value,
	}
}

func NewDeviceLoginRequestIdNullFromString(value string) DeviceLoginRequestIdNull {
	return DeviceLoginRequestIdNull{
		filled: true,
		value:  DeviceLoginRequestId(value),
	}
}

func (parameter DeviceLoginRequestIdNull) Get() (DeviceLoginRequestId, bool) {
	return parameter.value, parameter.filled
}

func (parameter DeviceLoginRequestIdNull) Filled() bool {
	return parameter.filled
}

func (parameter DeviceLoginRequestIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *DeviceLoginRequestIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
