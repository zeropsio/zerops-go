// Generated ZEROPS sdk

package types

import (
	"encoding/json"
)

type ObjectMap[K comparable, V any] map[K]V

func NewObjectMap[K comparable, V any](value map[K]V) ObjectMap[K, V] {
	return value
}

func (parameter ObjectMap[K, V]) Native() map[K]V {
	return map[K]V(parameter)
}

func (parameter ObjectMap[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(parameter)
}

func (parameter *ObjectMap[K, V]) UnmarshalJSON(data []byte) error {
	var x map[K]V
	err := json.Unmarshal(data, &x)
	*parameter = x
	return err
}
