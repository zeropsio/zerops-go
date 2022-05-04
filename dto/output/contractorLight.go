// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ContractorLight struct {
	Id            uuid.ContractorId `json:"id"`
	CompanyName   types.String      `json:"companyName"`
	IsActive      types.Bool        `json:"isActive"`
	CompanyNumber types.String      `json:"companyNumber"`
	VatNumber     types.String      `json:"vatNumber"`
}
