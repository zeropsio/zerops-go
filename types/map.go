// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

type Map map[string]interface{}

func NewMap(value map[string]interface{}) Map {
	return Map(value)
}

func (parameter Map) Native() map[string]interface{} {
	return map[string]interface{}(parameter)
}

func (parameter Map) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}(parameter))
}

func (parameter *Map) UnmarshalJSON(data []byte) error {
	var x map[string]interface{}
	err := json.Unmarshal(data, &x)
	*parameter = Map(x)
	return err
}
