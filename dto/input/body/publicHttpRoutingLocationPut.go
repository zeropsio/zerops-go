// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingLocationPut)(nil)

type PublicHttpRoutingLocationPut struct {
	Locations PublicHttpRoutingLocationPutLocations `json:"locations"`
}

func (dto PublicHttpRoutingLocationPut) GetLocations() PublicHttpRoutingLocationPutLocations {
	return dto.Locations
}

type PublicHttpRoutingLocationPutLocations []PublicHttpRoutingLocation

func (dto PublicHttpRoutingLocationPutLocations) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingLocation(dto))
}

func (dto *PublicHttpRoutingLocationPut) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Locations *PublicHttpRoutingLocationPutLocations
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingLocationPut", err)
	}
	var errorList validator.ErrorList
	if aux.Locations == nil {
		errorList = errorList.With(validator.NewError("locations", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Locations = *aux.Locations

	return nil
}
