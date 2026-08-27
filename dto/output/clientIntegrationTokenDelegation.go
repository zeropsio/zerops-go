// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientIntegrationTokenDelegation struct {
	Id               uuid.IntegrationTokenDelegationId    `json:"id"`
	ClientId         uuid.ClientId                        `json:"clientId"`
	ClientUserId     uuid.ClientUserId                    `json:"clientUserId"`
	UserId           uuid.UserId                          `json:"userId"`
	TokenId          uuid.UserId                          `json:"tokenId"`
	TokenPermissions IntegrationTokenDelegationPermission `json:"tokenPermissions"`
	Created          types.DateTime                       `json:"created"`
	LastUpdate       types.DateTime                       `json:"lastUpdate"`
}
