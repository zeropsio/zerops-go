// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutAppVersion)(nil)

type PutAppVersion struct {
	Name types.StringNull `json:"name"`
}

func (dto PutAppVersion) GetName() types.StringNull {
	return dto.Name
}

func (dto *PutAppVersion) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutAppVersion", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = aux.Name

	return nil
}
