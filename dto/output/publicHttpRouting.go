// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type PublicHttpRouting struct {
	Id           uuid.PublicHttpRoutingId   `json:"id"`
	ClientId     uuid.ClientId              `json:"clientId"`
	ProjectId    uuid.ProjectId             `json:"projectId"`
	SslEnabled   types.Bool                 `json:"sslEnabled"`
	Domains      PublicHttpRoutingDomains   `json:"domains"`
	Locations    PublicHttpRoutingLocations `json:"locations"`
	Created      types.DateTime             `json:"created"`
	LastUpdate   types.DateTime             `json:"lastUpdate"`
	IsSynced     types.Bool                 `json:"isSynced"`
	IsEditable   types.Bool                 `json:"isEditable"`
	DeleteOnSync types.Bool                 `json:"deleteOnSync"`
}

type PublicHttpRoutingDomains []PublicHttpRoutingDomain

func (dto PublicHttpRoutingDomains) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingDomain(dto))
}

type PublicHttpRoutingLocations []PublicHttpRoutingLocation

func (dto PublicHttpRoutingLocations) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingLocation(dto))
}
