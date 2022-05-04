// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type PublicHttpRoutingId types.UuidShort

func NewPublicHttpRoutingIdFromString(value string) (out PublicHttpRoutingId, err error) {
	return PublicHttpRoutingId(value), nil
}

func (parameter PublicHttpRoutingId) Native() string {
	return string(parameter)
}

func (parameter PublicHttpRoutingId) TypedString() types.String {
	return types.String(parameter)
}

type PublicHttpRoutingIdNull struct {
	value  PublicHttpRoutingId
	filled bool
}

func (parameter PublicHttpRoutingId) PublicHttpRoutingIdNull() PublicHttpRoutingIdNull {
	return NewPublicHttpRoutingIdNull(parameter)
}

func NewPublicHttpRoutingIdNull(value PublicHttpRoutingId) PublicHttpRoutingIdNull {
	return PublicHttpRoutingIdNull{
		filled: true,
		value:  value,
	}
}

func NewPublicHttpRoutingIdNullFromString(value string) PublicHttpRoutingIdNull {
	return PublicHttpRoutingIdNull{
		filled: true,
		value:  PublicHttpRoutingId(value),
	}
}

func (parameter PublicHttpRoutingIdNull) Get() (PublicHttpRoutingId, bool) {
	return parameter.value, parameter.filled
}

func (parameter PublicHttpRoutingIdNull) Filled() bool {
	return parameter.filled
}

func (parameter PublicHttpRoutingIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *PublicHttpRoutingIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
