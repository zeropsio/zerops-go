// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type RequestedPortsJsonObject struct {
	IsActive       types.Bool                             `json:"isActive"`
	RequestedPorts RequestedPortsJsonObjectRequestedPorts `json:"requestedPorts"`
}

type RequestedPortsJsonObjectRequestedPorts []ServicePort

func (dto RequestedPortsJsonObjectRequestedPorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServicePort(dto))
}
