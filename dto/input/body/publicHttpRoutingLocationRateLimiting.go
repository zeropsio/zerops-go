// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingLocationRateLimiting)(nil)

type PublicHttpRoutingLocationRateLimiting struct {
	Enabled  types.Bool                                        `json:"enabled"`
	Key      enum.PublicHttpRoutingLocationRateLimitingKeyEnum `json:"key"`
	Burst    types.Int                                         `json:"burst"`
	Rate     types.Int                                         `json:"rate"`
	Zone     types.EmptyString                                 `json:"zone"`
	ZoneSize types.Int                                         `json:"zoneSize"`
}

func (dto PublicHttpRoutingLocationRateLimiting) GetEnabled() types.Bool {
	return dto.Enabled
}
func (dto PublicHttpRoutingLocationRateLimiting) GetKey() enum.PublicHttpRoutingLocationRateLimitingKeyEnum {
	return dto.Key
}
func (dto PublicHttpRoutingLocationRateLimiting) GetBurst() types.Int {
	return dto.Burst
}
func (dto PublicHttpRoutingLocationRateLimiting) GetRate() types.Int {
	return dto.Rate
}
func (dto PublicHttpRoutingLocationRateLimiting) GetZone() types.EmptyString {
	return dto.Zone
}
func (dto PublicHttpRoutingLocationRateLimiting) GetZoneSize() types.Int {
	return dto.ZoneSize
}

func (dto *PublicHttpRoutingLocationRateLimiting) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Enabled  *types.Bool
		Key      *enum.PublicHttpRoutingLocationRateLimitingKeyEnum
		Burst    *types.Int
		Rate     *types.Int
		Zone     *types.EmptyString
		ZoneSize *types.Int
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingLocationRateLimiting", err)
	}
	var errorList validator.ErrorList
	if aux.Enabled == nil {
		errorList = errorList.With(validator.NewError("enabled", "field is required"))
	}
	if aux.Key == nil {
		errorList = errorList.With(validator.NewError("key", "field is required"))
	}
	if aux.Burst == nil {
		errorList = errorList.With(validator.NewError("burst", "field is required"))
	}
	if aux.Rate == nil {
		errorList = errorList.With(validator.NewError("rate", "field is required"))
	}
	if aux.Zone == nil {
		errorList = errorList.With(validator.NewError("zone", "field is required"))
	}
	if aux.ZoneSize == nil {
		errorList = errorList.With(validator.NewError("zoneSize", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Enabled = *aux.Enabled
	dto.Key = *aux.Key
	dto.Burst = *aux.Burst
	dto.Rate = *aux.Rate
	dto.Zone = *aux.Zone
	dto.ZoneSize = *aux.ZoneSize

	return nil
}
