// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostServiceStackBackupModify)(nil)

type PostServiceStackBackupModify struct {
	Tags types.StringArrayNull `json:"tags"`
}

func (dto PostServiceStackBackupModify) GetTags() types.StringArrayNull {
	return dto.Tags
}

func (dto *PostServiceStackBackupModify) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Tags types.StringArrayNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostServiceStackBackupModify", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.Tags = aux.Tags

	return nil
}
