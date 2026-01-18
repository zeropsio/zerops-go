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

type GetRecipeService struct {
	Name               types.String                       `json:"name"`
	TypeId             stringId.ServiceStackTypeId        `json:"typeId"`
	TypeName           types.String                       `json:"typeName"`
	TypeVersionId      stringId.ServiceStackTypeVersionId `json:"typeVersionId"`
	TypeVersionName    types.String                       `json:"typeVersionName"`
	ExactVersionNumber types.String                       `json:"exactVersionNumber"`
	Category           enum.ServiceStackTypeCategoryEnum  `json:"category"`
	Mode               enum.ServiceStackModeEnum          `json:"mode"`
	GitRepo            types.StringNull                   `json:"gitRepo"`
	Content            types.StringNull                   `json:"content"`
	ZeropsYaml         types.StringNull                   `json:"zeropsYaml"`
	Extracts           types.Map                          `json:"extracts"`
	Autoscaling        *CustomAutoscaling                 `json:"autoscaling"`
	Ports              GetRecipeServicePorts              `json:"ports"`
	ObjectStorageSize  types.IntNull                      `json:"objectStorageSize"`
	IsUtility          types.Bool                         `json:"isUtility"`
}

type GetRecipeServicePorts []ServicePort

func (dto GetRecipeServicePorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServicePort(dto))
}
