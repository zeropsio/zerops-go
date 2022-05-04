// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ClientExtra struct {
	Id            uuid.ClientId    `json:"id"`
	CompanyName   types.String     `json:"companyName"`
	CompanyNumber types.StringNull `json:"companyNumber"`
	VatNumber     types.StringNull `json:"vatNumber"`
	Avatar        *ClientAvatar    `json:"avatar"`
	AccountName   types.String     `json:"accountName"`
}
