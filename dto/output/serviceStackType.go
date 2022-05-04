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

type ServiceStackType struct {
	Id                          stringId.ServiceStackTypeId                 `json:"id"`
	Name                        types.String                                `json:"name"`
	Description                 types.Text                                  `json:"description"`
	Created                     types.DateTime                              `json:"created"`
	LastUpdate                  types.DateTime                              `json:"lastUpdate"`
	Category                    enum.ServiceStackTypeCategoryEnum           `json:"category"`
	ServiceStackTypeVersionList ServiceStackTypeServiceStackTypeVersionList `json:"serviceStackTypeVersionList"`
}

type ServiceStackTypeServiceStackTypeVersionList []ServiceStackTypeVersion

func (dto ServiceStackTypeServiceStackTypeVersionList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackTypeVersion(dto))
}
