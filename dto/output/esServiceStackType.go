// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type EsServiceStackType struct {
	Id                           stringId.ServiceStackTypeId                   `json:"id"`
	Name                         types.String                                  `json:"name"`
	Description                  types.Text                                    `json:"description"`
	DefaultServiceStackVersionId stringId.ServiceStackTypeVersionIdNull        `json:"defaultServiceStackVersionId"`
	Category                     enum.ServiceStackTypeCategoryEnum             `json:"category"`
	Created                      types.DateTime                                `json:"created"`
	LastUpdate                   types.DateTime                                `json:"lastUpdate"`
	DefaultServiceStackVersion   *ServiceStackTypeVersionLight                 `json:"defaultServiceStackVersion"`
	ServiceStackTypeVersionList  EsServiceStackTypeServiceStackTypeVersionList `json:"serviceStackTypeVersionList"`
}

type EsServiceStackTypeServiceStackTypeVersionList []ServiceStackTypeVersionLight

func (dto EsServiceStackTypeServiceStackTypeVersionList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackTypeVersionLight(dto))
}
