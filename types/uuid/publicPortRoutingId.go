// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type PublicPortRoutingId types.UuidShort

func NewPublicPortRoutingIdFromString(value string) (out PublicPortRoutingId, err error) {
	return PublicPortRoutingId(value), nil
}

func (parameter PublicPortRoutingId) Native() string {
	return string(parameter)
}

func (parameter PublicPortRoutingId) TypedString() types.String {
	return types.String(parameter)
}

type PublicPortRoutingIdNull struct {
	value  PublicPortRoutingId
	filled bool
}

func (parameter PublicPortRoutingId) PublicPortRoutingIdNull() PublicPortRoutingIdNull {
	return NewPublicPortRoutingIdNull(parameter)
}

func NewPublicPortRoutingIdNull(value PublicPortRoutingId) PublicPortRoutingIdNull {
	return PublicPortRoutingIdNull{
		filled: true,
		value:  value,
	}
}

func NewPublicPortRoutingIdNullFromString(value string) PublicPortRoutingIdNull {
	return PublicPortRoutingIdNull{
		filled: true,
		value:  PublicPortRoutingId(value),
	}
}

func (parameter PublicPortRoutingIdNull) Get() (PublicPortRoutingId, bool) {
	return parameter.value, parameter.filled
}

func (parameter PublicPortRoutingIdNull) Filled() bool {
	return parameter.filled
}

func (parameter PublicPortRoutingIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *PublicPortRoutingIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
