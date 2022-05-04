// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"errors"
)

type IntArray []int

func NewIntArray(value []int) IntArray {
	return IntArray(value)
}

func (parameter IntArray) Native() []int {
	return []int(parameter)
}

func (parameter IntArray) MarshalJSON() ([]byte, error) {
	if parameter == nil {
		return nil, errors.New("value can't be nil")
	}
	return json.Marshal([]int(parameter))
}

func (parameter *IntArray) UnmarshalJSON(data []byte) error {
	var arr []int
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	if arr == nil {
		arr = make([]int, 0)
	}
	*parameter = IntArray(arr)
	return nil
}
