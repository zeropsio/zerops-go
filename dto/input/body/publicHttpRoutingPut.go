// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingPut)(nil)

type PublicHttpRoutingPut struct {
	SslEnabled types.Bool                    `json:"sslEnabled"`
	Domains    types.StringArray             `json:"domains"`
	Locations  PublicHttpRoutingPutLocations `json:"locations"`
}

func (dto PublicHttpRoutingPut) GetSslEnabled() types.Bool {
	return dto.SslEnabled
}
func (dto PublicHttpRoutingPut) GetDomains() types.StringArray {
	return dto.Domains
}
func (dto PublicHttpRoutingPut) GetLocations() PublicHttpRoutingPutLocations {
	return dto.Locations
}

type PublicHttpRoutingPutLocations []PublicHttpRoutingLocationPut

func (dto PublicHttpRoutingPutLocations) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingLocationPut(dto))
}

func (dto *PublicHttpRoutingPut) UnmarshalJSON(b []byte) error {
	var aux = struct {
		SslEnabled *types.Bool
		Domains    *types.StringArray
		Locations  *PublicHttpRoutingPutLocations
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingPut", err)
	}
	var errorList validator.ErrorList
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
	dto.SslEnabled = *aux.SslEnabled
	dto.Domains = *aux.Domains
	dto.Locations = *aux.Locations

	return nil
}
