// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"errors"
)

type ObjectArray []string

func NewObjectArray(value []string) ObjectArray {
	return ObjectArray(value)
}

func (parameter ObjectArray) Native() []string {
	return []string(parameter)
}

func (parameter ObjectArray) MarshalJSON() ([]byte, error) {
	if parameter == nil {
		return nil, errors.New("value can't be nil")
	}
	return json.Marshal([]string(parameter))
}

func (parameter *ObjectArray) UnmarshalJSON(data []byte) error {
	var arr []string
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	if arr == nil {
		arr = make([]string, 0)
	}
	*parameter = ObjectArray(arr)
	return nil
}
