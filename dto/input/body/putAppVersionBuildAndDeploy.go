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
	ZeropsYaml      types.MediumText `json:"zeropsYaml"`
	ZeropsYamlSetup types.StringNull `json:"zeropsYamlSetup"`
}

func (dto PutAppVersionBuildAndDeploy) GetZeropsYaml() types.MediumText {
	return dto.ZeropsYaml
}
func (dto PutAppVersionBuildAndDeploy) GetZeropsYamlSetup() types.StringNull {
	return dto.ZeropsYamlSetup
}

func (dto *PutAppVersionBuildAndDeploy) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ZeropsYaml      *types.MediumText
		ZeropsYamlSetup types.StringNull
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
	dto.ZeropsYamlSetup = aux.ZeropsYamlSetup

	return nil
}
