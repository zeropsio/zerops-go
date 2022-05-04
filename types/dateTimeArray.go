// Generated ZEROPS sdk

package types

import (
	"encoding/json"
	"errors"
	"time"
)

type DateTimeArray []time.Time

func NewDateTimeArray(value []time.Time) DateTimeArray {
	return DateTimeArray(value)
}

func (parameter DateTimeArray) Native() []time.Time {
	return []time.Time(parameter)
}

func (parameter DateTimeArray) MarshalJSON() ([]byte, error) {
	if parameter == nil {
		return nil, errors.New("value can't be nil")
	}
	return json.Marshal([]time.Time(parameter))
}

func (parameter *DateTimeArray) UnmarshalJSON(data []byte) error {
	var arr []time.Time
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	if arr == nil {
		arr = make([]time.Time, 0)
	}
	*parameter = DateTimeArray(arr)
	return nil
}
