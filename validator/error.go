package validator

import (
	"encoding/json"
	"fmt"
	"strings"
)

func NewError(parameterName, message string) Error {
	return Error{
		ParameterName: parameterName,
		Message:       message,
	}
}

type Error struct {
	ParameterName string `json:"parameter"`
	Message       string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.ParameterName, e.Message)
}

type ErrorList []Error

type ErrorListError struct {
	ErrorList
}

func (l ErrorListError) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ErrorList)
}

func (l ErrorListError) Error() string {
	return l.ErrorList.String()
}

func (l ErrorList) With(errors ...error) ErrorList {
	return l.WithPrefix("", errors...)
}

func (l ErrorList) WithPrefix(prefix string, errorList ...error) (result ErrorList) {
	result = l

	// back compatibility
	prefix = strings.Trim(prefix, ".")

	for _, err := range errorList {
		if err == nil {
			continue
		}
		switch typedErr := err.(type) {
		case ErrorListError:
			for _, validationErr := range typedErr.ErrorList {
				// back compatibility
				parameterName := validationErr.ParameterName
				if prefix != "" {
					parameterName = prefix + "." + parameterName
				}
				result = append(result, NewError(parameterName, validationErr.Message))
			}

		case Error:
			// back compatibility
			parameterName := typedErr.ParameterName
			if prefix != "" {
				parameterName = prefix + "." + parameterName
			}
			result = append(result, NewError(parameterName, typedErr.Message))

		default:
			result = append(result, NewError(prefix, typedErr.Error()))
		}
	}
	return
}

func (l ErrorList) GetError() error {
	if l == nil {
		return nil
	}
	return ErrorListError{l}
}

func (l ErrorList) String() string {
	var sb strings.Builder
	for _, e := range l {
		sb.WriteString(e.Error())
		sb.WriteByte('\n')
	}
	return sb.String()
}
