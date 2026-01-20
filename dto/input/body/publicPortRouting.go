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
var _ json.Unmarshaler = (*PublicPortRouting)(nil)

type PublicPortRouting struct {
	PublicIpType      enum.PublicPortRoutingPublicIpTypeEnum   `json:"publicIpType"`
	PublicPort        types.Int                                `json:"publicPort"`
	InternalPort      types.Int                                `json:"internalPort"`
	InternalProtocol  enum.ServicePortProtocolEnum             `json:"internalProtocol"`
	FirewallPolicy    enum.PublicPortRoutingFirewallPolicyEnum `json:"firewallPolicy"`
	FirewallIpRanges  types.StringArray                        `json:"firewallIpRanges"`
	FirewallAllowMyIp types.Bool                               `json:"firewallAllowMyIp"`
}

func (dto PublicPortRouting) GetPublicIpType() enum.PublicPortRoutingPublicIpTypeEnum {
	return dto.PublicIpType
}
func (dto PublicPortRouting) GetPublicPort() types.Int {
	return dto.PublicPort
}
func (dto PublicPortRouting) GetInternalPort() types.Int {
	return dto.InternalPort
}
func (dto PublicPortRouting) GetInternalProtocol() enum.ServicePortProtocolEnum {
	return dto.InternalProtocol
}
func (dto PublicPortRouting) GetFirewallPolicy() enum.PublicPortRoutingFirewallPolicyEnum {
	return dto.FirewallPolicy
}
func (dto PublicPortRouting) GetFirewallIpRanges() types.StringArray {
	return dto.FirewallIpRanges
}
func (dto PublicPortRouting) GetFirewallAllowMyIp() types.Bool {
	return dto.FirewallAllowMyIp
}

func (dto *PublicPortRouting) UnmarshalJSON(b []byte) error {
	var aux = struct {
		PublicIpType      *enum.PublicPortRoutingPublicIpTypeEnum
		PublicPort        *types.Int
		InternalPort      *types.Int
		InternalProtocol  *enum.ServicePortProtocolEnum
		FirewallPolicy    *enum.PublicPortRoutingFirewallPolicyEnum
		FirewallIpRanges  *types.StringArray
		FirewallAllowMyIp *types.Bool
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicPortRouting", err)
	}
	var errorList validator.ErrorList
	if aux.PublicIpType == nil {
		errorList = errorList.With(validator.NewError("publicIpType", "field is required"))
	}
	if aux.PublicPort == nil {
		errorList = errorList.With(validator.NewError("publicPort", "field is required"))
	}
	if aux.InternalPort == nil {
		errorList = errorList.With(validator.NewError("internalPort", "field is required"))
	}
	if aux.InternalProtocol == nil {
		errorList = errorList.With(validator.NewError("internalProtocol", "field is required"))
	}
	if aux.FirewallPolicy == nil {
		errorList = errorList.With(validator.NewError("firewallPolicy", "field is required"))
	}
	if aux.FirewallIpRanges == nil {
		errorList = errorList.With(validator.NewError("firewallIpRanges", "field is required"))
	}
	if aux.FirewallAllowMyIp == nil {
		errorList = errorList.With(validator.NewError("firewallAllowMyIp", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.PublicIpType = *aux.PublicIpType
	dto.PublicPort = *aux.PublicPort
	dto.InternalPort = *aux.InternalPort
	dto.InternalProtocol = *aux.InternalProtocol
	dto.FirewallPolicy = *aux.FirewallPolicy
	dto.FirewallIpRanges = *aux.FirewallIpRanges
	dto.FirewallAllowMyIp = *aux.FirewallAllowMyIp

	return nil
}
