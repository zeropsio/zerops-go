// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"errors"
)

type Int64Array []int64

func NewInt64Array(value []int64) Int64Array {
	return Int64Array(value)
}

func (parameter Int64Array) Native() []int64 {
	return []int64(parameter)
}

func (parameter Int64Array) MarshalJSON() ([]byte, error) {
	if parameter == nil {
		return nil, errors.New("value can't be nil")
	}
	return json.Marshal([]int64(parameter))
}

func (parameter *Int64Array) UnmarshalJSON(data []byte) error {
	var arr []int64
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	if arr == nil {
		arr = make([]int64, 0)
	}
	*parameter = Int64Array(arr)
	return nil
}
