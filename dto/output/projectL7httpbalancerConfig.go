// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ProjectL7httpbalancerConfig struct {
	Values ProjectL7httpbalancerConfigValues `json:"values"`
}

type ProjectL7httpbalancerConfigValues []ProjectL7httpbalancerConfigValue

func (dto ProjectL7httpbalancerConfigValues) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectL7httpbalancerConfigValue(dto))
}
