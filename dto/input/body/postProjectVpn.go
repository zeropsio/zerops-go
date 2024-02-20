// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostProjectVpn)(nil)

type PostProjectVpn struct {
	VpnPublicKey types.String `json:"vpnPublicKey"`
}

func (dto PostProjectVpn) GetVpnPublicKey() types.String {
	return dto.VpnPublicKey
}

func (dto *PostProjectVpn) UnmarshalJSON(b []byte) error {
	var aux = struct {
		VpnPublicKey *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostProjectVpn", err)
	}
	var errorList validator.ErrorList
	if aux.VpnPublicKey == nil {
		errorList = errorList.With(validator.NewError("vpnPublicKey", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.VpnPublicKey = *aux.VpnPublicKey

	return nil
}
