// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostS3Bucket)(nil)

type PostS3Bucket struct {
	Name    types.String     `json:"name"`
	XAmzAcl types.StringNull `json:"xAmzAcl"`
}

func (dto PostS3Bucket) GetName() types.String {
	return dto.Name
}
func (dto PostS3Bucket) GetXAmzAcl() types.StringNull {
	return dto.XAmzAcl
}

func (dto *PostS3Bucket) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name    *types.String
		XAmzAcl types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostS3Bucket", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.XAmzAcl = aux.XAmzAcl

	return nil
}
