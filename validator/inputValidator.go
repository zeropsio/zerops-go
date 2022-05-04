package validator

import (
	"encoding/json"
	"fmt"
)

func JsonValidation(name string, err error) error {
	if err == nil {
		return nil
	}
	switch typeError := err.(type) {

	case *json.UnmarshalTypeError:
		return NewError(name+"."+typeError.Field, fmt.Sprintf("value on offset %d must be of type '%s', not '%s'",
			typeError.Offset,
			typeError.Type.String(),
			typeError.Value,
		))

	case *json.SyntaxError:
		return NewError(name, fmt.Sprintf("syntax error on line %d - %s",
			typeError.Offset,
			typeError.Error(),
		))

	case Error, ErrorListError:
		return typeError

	default:
		return NewError(name, err.Error())
	}
}
