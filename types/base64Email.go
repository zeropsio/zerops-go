// Generated ZEROPS sdk

package types

import (
	"encoding/base64"
	"errors"
)

type Base64Email string

func NewBase64Email(value string) Base64Email {
	return Base64Email(value)
}

func (parameter Base64Email) Native() string {
	return string(parameter)
}

func (parameter Base64Email) Decoded() (string, error) {
	decodedEmail, err := base64.StdEncoding.DecodeString(string(parameter))
	if err != nil {
		return "", errors.New("value is not valid base64 format, " + err.Error())
	}

	return string(decodedEmail), nil
}
