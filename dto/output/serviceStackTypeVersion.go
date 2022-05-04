// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type ServiceStackTypeVersion struct {
	Id                 stringId.ServiceStackTypeVersionId     `json:"id"`
	Name               types.String                           `json:"name"`
	IsBuild            types.Bool                             `json:"isBuild"`
	ServiceStackTypeId stringId.ServiceStackTypeId            `json:"serviceStackTypeId"`
	Status             enum.ServiceStackTypeVersionStatusEnum `json:"status"`
	ReleaseDate        types.DateTime                         `json:"releaseDate"`
	Created            types.DateTime                         `json:"created"`
	LastUpdate         types.DateTime                         `json:"lastUpdate"`
	Description        types.Text                             `json:"description"`
	HaDriverId         stringId.DriverIdNull                  `json:"haDriverId"`
	NonHaDriverId      stringId.DriverIdNull                  `json:"nonHaDriverId"`
	VersionNumber      types.EmptyString                      `json:"versionNumber"`
	ExactVersionNumber types.EmptyString                      `json:"exactVersionNumber"`
	Config             ServiceStackTypeVersionConfig          `json:"config"`
}
