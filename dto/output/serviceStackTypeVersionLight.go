// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type ServiceStackTypeVersionLight struct {
	Id                 stringId.ServiceStackTypeVersionId     `json:"id"`
	ServiceStackTypeId stringId.ServiceStackTypeId            `json:"serviceStackTypeId"`
	Name               types.String                           `json:"name"`
	IsBuild            types.Bool                             `json:"isBuild"`
	UpdateUrl          types.Text                             `json:"updateUrl"`
	Status             enum.ServiceStackTypeVersionStatusEnum `json:"status"`
	ReleaseDate        types.DateTime                         `json:"releaseDate"`
}
