// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type FileBrowsingAccessToken types.UuidShort

func NewFileBrowsingAccessTokenFromString(value string) (out FileBrowsingAccessToken, err error) {
	return FileBrowsingAccessToken(value), nil
}

func (parameter FileBrowsingAccessToken) Native() string {
	return string(parameter)
}

func (parameter FileBrowsingAccessToken) TypedString() types.String {
	return types.String(parameter)
}

type FileBrowsingAccessTokenNull struct {
	value  FileBrowsingAccessToken
	filled bool
}

func (parameter FileBrowsingAccessToken) FileBrowsingAccessTokenNull() FileBrowsingAccessTokenNull {
	return NewFileBrowsingAccessTokenNull(parameter)
}

func NewFileBrowsingAccessTokenNull(value FileBrowsingAccessToken) FileBrowsingAccessTokenNull {
	return FileBrowsingAccessTokenNull{
		filled: true,
		value:  value,
	}
}

func NewFileBrowsingAccessTokenNullFromString(value string) FileBrowsingAccessTokenNull {
	return FileBrowsingAccessTokenNull{
		filled: true,
		value:  FileBrowsingAccessToken(value),
	}
}

func (parameter FileBrowsingAccessTokenNull) Get() (FileBrowsingAccessToken, bool) {
	return parameter.value, parameter.filled
}

func (parameter FileBrowsingAccessTokenNull) Filled() bool {
	return parameter.filled
}

func (parameter FileBrowsingAccessTokenNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *FileBrowsingAccessTokenNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
