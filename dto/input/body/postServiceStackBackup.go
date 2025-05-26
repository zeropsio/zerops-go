// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostServiceStackBackup)(nil)

type PostServiceStackBackup struct {
	Tags types.StringArrayNull `json:"tags"`
}

func (dto PostServiceStackBackup) GetTags() types.StringArrayNull {
	return dto.Tags
}

func (dto *PostServiceStackBackup) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Tags types.StringArrayNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostServiceStackBackup", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.Tags = aux.Tags

	return nil
}
