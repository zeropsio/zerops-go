// Generated ZEROPS sdk

package types

import (
	"encoding/base64"
	"errors"
	"net/url"
)

type UrlEncodedBase64Email string

func NewUrlEncodedBase64Email(value string) UrlEncodedBase64Email {
	return UrlEncodedBase64Email(value)
}

func (parameter UrlEncodedBase64Email) Native() string {
	return string(parameter)
}

func (parameter UrlEncodedBase64Email) Decoded() (string, error) {
	urlDecoded, err := url.QueryUnescape(string(parameter))
	if err != nil {
		return "", errors.New("value is not valid url format, " + err.Error())
	}

	base64Decoded, err := base64.StdEncoding.DecodeString(urlDecoded)
	if err != nil {
		return "", errors.New("value is not valid base64 format, " + err.Error())
	}

	return string(base64Decoded), nil
}
