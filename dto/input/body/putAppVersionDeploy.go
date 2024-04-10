// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutAppVersionDeploy)(nil)

type PutAppVersionDeploy struct {
	ZeropsYaml      types.MediumTextNull `json:"zeropsYaml"`
	ZeropsYamlSetup types.StringNull     `json:"zeropsYamlSetup"`
}

func (dto PutAppVersionDeploy) GetZeropsYaml() types.MediumTextNull {
	return dto.ZeropsYaml
}
func (dto PutAppVersionDeploy) GetZeropsYamlSetup() types.StringNull {
	return dto.ZeropsYamlSetup
}

func (dto *PutAppVersionDeploy) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ZeropsYaml      types.MediumTextNull
		ZeropsYamlSetup types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutAppVersionDeploy", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.ZeropsYaml = aux.ZeropsYaml
	dto.ZeropsYamlSetup = aux.ZeropsYamlSetup

	return nil
}
