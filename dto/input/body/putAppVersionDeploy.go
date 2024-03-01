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
	ZeropsYaml types.MediumTextNull `json:"zeropsYaml"`
	Source     types.StringNull     `json:"source"`
}

func (dto PutAppVersionDeploy) GetZeropsYaml() types.MediumTextNull {
	return dto.ZeropsYaml
}
func (dto PutAppVersionDeploy) GetSource() types.StringNull {
	return dto.Source
}

func (dto *PutAppVersionDeploy) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ZeropsYaml types.MediumTextNull
		Source     types.StringNull
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
	dto.Source = aux.Source

	return nil
}
