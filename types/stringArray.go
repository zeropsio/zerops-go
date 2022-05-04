// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"errors"
)

type StringArray []string

func NewStringArray(value []string) StringArray {
	return StringArray(value)
}

func (parameter StringArray) Native() []string {
	return []string(parameter)
}

func (parameter StringArray) MarshalJSON() ([]byte, error) {
	if parameter == nil {
		return nil, errors.New("value can't be nil")
	}
	return json.Marshal([]string(parameter))
}

func (parameter *StringArray) UnmarshalJSON(data []byte) error {
	var arr []string
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	if arr == nil {
		arr = make([]string, 0)
	}
	*parameter = StringArray(arr)
	return nil
}
