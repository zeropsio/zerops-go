// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type PublicHttpRoutingProcess struct {
	Id         uuid.PublicHttpRoutingId          `json:"id"`
	ClientId   uuid.ClientId                     `json:"clientId"`
	ProjectId  uuid.ProjectId                    `json:"projectId"`
	SslEnabled types.Bool                        `json:"sslEnabled"`
	Domains    PublicHttpRoutingProcessDomains   `json:"domains"`
	Locations  PublicHttpRoutingProcessLocations `json:"locations"`
	Created    types.DateTime                    `json:"created"`
	LastUpdate types.DateTime                    `json:"lastUpdate"`
	IsEditable types.Bool                        `json:"isEditable"`
	CdnEnabled types.Bool                        `json:"cdnEnabled"`
	Process    Process                           `json:"process"`
}

type PublicHttpRoutingProcessDomains []PublicHttpRoutingDomain

func (dto PublicHttpRoutingProcessDomains) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingDomain(dto))
}

type PublicHttpRoutingProcessLocations []PublicHttpRoutingLocation

func (dto PublicHttpRoutingProcessLocations) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingLocation(dto))
}
