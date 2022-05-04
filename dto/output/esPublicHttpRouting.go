// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsPublicHttpRouting struct {
	Id           uuid.PublicHttpRoutingId     `json:"id"`
	ClientId     uuid.ClientId                `json:"clientId"`
	ProjectId    uuid.ProjectId               `json:"projectId"`
	SslEnabled   types.Bool                   `json:"sslEnabled"`
	Domains      EsPublicHttpRoutingDomains   `json:"domains"`
	Locations    EsPublicHttpRoutingLocations `json:"locations"`
	Created      types.DateTime               `json:"created"`
	LastUpdate   types.DateTime               `json:"lastUpdate"`
	IsSynced     types.Bool                   `json:"isSynced"`
	DeleteOnSync types.Bool                   `json:"deleteOnSync"`
	IsEditable   types.Bool                   `json:"isEditable"`
}

type EsPublicHttpRoutingDomains []PublicHttpRoutingDomain

func (dto EsPublicHttpRoutingDomains) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingDomain(dto))
}

type EsPublicHttpRoutingLocations []PublicHttpRoutingLocation

func (dto EsPublicHttpRoutingLocations) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingLocation(dto))
}
