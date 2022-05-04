// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingPost)(nil)

type PublicHttpRoutingPost struct {
	ProjectId  uuid.ProjectId                 `json:"projectId"`
	SslEnabled types.Bool                     `json:"sslEnabled"`
	Domains    types.StringArray              `json:"domains"`
	Locations  PublicHttpRoutingPostLocations `json:"locations"`
}

func (dto PublicHttpRoutingPost) GetProjectId() uuid.ProjectId {
	return dto.ProjectId
}
func (dto PublicHttpRoutingPost) GetSslEnabled() types.Bool {
	return dto.SslEnabled
}
func (dto PublicHttpRoutingPost) GetDomains() types.StringArray {
	return dto.Domains
}
func (dto PublicHttpRoutingPost) GetLocations() PublicHttpRoutingPostLocations {
	return dto.Locations
}

type PublicHttpRoutingPostLocations []PublicHttpRoutingLocationPost

func (dto PublicHttpRoutingPostLocations) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingLocationPost(dto))
}

func (dto *PublicHttpRoutingPost) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ProjectId  *uuid.ProjectId
		SslEnabled *types.Bool
		Domains    *types.StringArray
		Locations  *PublicHttpRoutingPostLocations
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingPost", err)
	}
	var errorList validator.ErrorList
	if aux.ProjectId == nil {
		errorList = errorList.With(validator.NewError("projectId", "field is required"))
	}
	if aux.SslEnabled == nil {
		errorList = errorList.With(validator.NewError("sslEnabled", "field is required"))
	}
	if aux.Domains == nil {
		errorList = errorList.With(validator.NewError("domains", "field is required"))
	}
	if aux.Locations == nil {
		errorList = errorList.With(validator.NewError("locations", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ProjectId = *aux.ProjectId
	dto.SslEnabled = *aux.SslEnabled
	dto.Domains = *aux.Domains
	dto.Locations = *aux.Locations

	return nil
}
