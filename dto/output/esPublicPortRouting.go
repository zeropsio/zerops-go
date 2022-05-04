// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsPublicPortRouting struct {
	Id                uuid.PublicPortRoutingId                 `json:"id"`
	ClientId          uuid.ClientId                            `json:"clientId"`
	ProjectId         uuid.ProjectId                           `json:"projectId"`
	ServiceStackId    uuid.ServiceStackId                      `json:"serviceStackId"`
	PublicIpType      enum.PublicPortRoutingPublicIpTypeEnum   `json:"publicIpType"`
	PublicPort        types.Int                                `json:"publicPort"`
	InternalPort      types.Int                                `json:"internalPort"`
	InternalProtocol  enum.ServicePortProtocolEnum             `json:"internalProtocol"`
	FirewallPolicy    enum.PublicPortRoutingFirewallPolicyEnum `json:"firewallPolicy"`
	FirewallIpRanges  types.StringArray                        `json:"firewallIpRanges"`
	FirewallAllowMyIp types.Bool                               `json:"firewallAllowMyIp"`
	Created           types.DateTime                           `json:"created"`
	LastUpdate        types.DateTime                           `json:"lastUpdate"`
	IsSynced          types.Bool                               `json:"isSynced"`
	DeleteOnSync      types.Bool                               `json:"deleteOnSync"`
}
