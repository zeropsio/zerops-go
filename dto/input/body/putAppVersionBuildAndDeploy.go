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
	BuildConfigContent types.MediumText `json:"buildConfigContent"`
	Source             types.StringNull `json:"source"`
}

func (dto PutAppVersionBuildAndDeploy) GetBuildConfigContent() types.MediumText {
	return dto.BuildConfigContent
}
func (dto PutAppVersionBuildAndDeploy) GetSource() types.StringNull {
	return dto.Source
}

func (dto *PutAppVersionBuildAndDeploy) UnmarshalJSON(b []byte) error {
	var aux = struct {
		BuildConfigContent *types.MediumText
		Source             types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutAppVersionBuildAndDeploy", err)
	}
	var errorList validator.ErrorList
	if aux.BuildConfigContent == nil {
		errorList = errorList.With(validator.NewError("buildConfigContent", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.BuildConfigContent = *aux.BuildConfigContent
	dto.Source = aux.Source

	return nil
}
