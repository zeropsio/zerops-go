// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"errors"
)

type FloatArray []float64

func NewFloatArray(value []float64) FloatArray {
	return FloatArray(value)
}

func (parameter FloatArray) Native() []float64 {
	return []float64(parameter)
}

func (parameter FloatArray) MarshalJSON() ([]byte, error) {
	if parameter == nil {
		return nil, errors.New("value can't be nil")
	}
	return json.Marshal([]float64(parameter))
}

func (parameter *FloatArray) UnmarshalJSON(data []byte) error {
	var arr []float64
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	if arr == nil {
		arr = make([]float64, 0)
	}
	*parameter = FloatArray(arr)
	return nil
}
