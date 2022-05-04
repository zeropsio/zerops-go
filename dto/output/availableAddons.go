// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type AvailableAddons struct {
	ClientAddons  AvailableAddonsClientAddons  `json:"clientAddons"`
	ProjectAddons AvailableAddonsProjectAddons `json:"projectAddons"`
}

type AvailableAddonsClientAddons []Addon

func (dto AvailableAddonsClientAddons) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Addon(dto))
}

type AvailableAddonsProjectAddons []Addon

func (dto AvailableAddonsProjectAddons) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Addon(dto))
}
