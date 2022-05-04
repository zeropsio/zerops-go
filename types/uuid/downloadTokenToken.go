// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type DownloadTokenToken types.UuidShort

func NewDownloadTokenTokenFromString(value string) (out DownloadTokenToken, err error) {
	return DownloadTokenToken(value), nil
}

func (parameter DownloadTokenToken) Native() string {
	return string(parameter)
}

func (parameter DownloadTokenToken) TypedString() types.String {
	return types.String(parameter)
}

type DownloadTokenTokenNull struct {
	value  DownloadTokenToken
	filled bool
}

func (parameter DownloadTokenToken) DownloadTokenTokenNull() DownloadTokenTokenNull {
	return NewDownloadTokenTokenNull(parameter)
}

func NewDownloadTokenTokenNull(value DownloadTokenToken) DownloadTokenTokenNull {
	return DownloadTokenTokenNull{
		filled: true,
		value:  value,
	}
}

func NewDownloadTokenTokenNullFromString(value string) DownloadTokenTokenNull {
	return DownloadTokenTokenNull{
		filled: true,
		value:  DownloadTokenToken(value),
	}
}

func (parameter DownloadTokenTokenNull) Get() (DownloadTokenToken, bool) {
	return parameter.value, parameter.filled
}

func (parameter DownloadTokenTokenNull) Filled() bool {
	return parameter.filled
}

func (parameter DownloadTokenTokenNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *DownloadTokenTokenNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
