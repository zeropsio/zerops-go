// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutAppVersionBuildAndDeploy)(nil)

type PutAppVersionBuildAndDeploy struct {
	ZeropsYaml types.MediumText `json:"zeropsYaml"`
	Source     types.StringNull `json:"source"`
}

func (dto PutAppVersionBuildAndDeploy) GetZeropsYaml() types.MediumText {
	return dto.ZeropsYaml
}
func (dto PutAppVersionBuildAndDeploy) GetSource() types.StringNull {
	return dto.Source
}

func (dto *PutAppVersionBuildAndDeploy) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ZeropsYaml *types.MediumText
		Source     types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutAppVersionBuildAndDeploy", err)
	}
	var errorList validator.ErrorList
	if aux.ZeropsYaml == nil {
		errorList = errorList.With(validator.NewError("zeropsYaml", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ZeropsYaml = *aux.ZeropsYaml
	dto.Source = aux.Source

	return nil
}
