// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostObjectStorageServiceStack)(nil)

type PostObjectStorageServiceStack struct {
	Name       types.String     `json:"name"`
	DiskGBytes types.Int        `json:"diskGBytes"`
	Policy     types.StringNull `json:"policy"`
	RawPolicy  types.TextNull   `json:"rawPolicy"`
	CdnEnabled types.BoolNull   `json:"cdnEnabled"`
}

func (dto PostObjectStorageServiceStack) GetName() types.String {
	return dto.Name
}
func (dto PostObjectStorageServiceStack) GetDiskGBytes() types.Int {
	return dto.DiskGBytes
}
func (dto PostObjectStorageServiceStack) GetPolicy() types.StringNull {
	return dto.Policy
}
func (dto PostObjectStorageServiceStack) GetRawPolicy() types.TextNull {
	return dto.RawPolicy
}
func (dto PostObjectStorageServiceStack) GetCdnEnabled() types.BoolNull {
	return dto.CdnEnabled
}

func (dto *PostObjectStorageServiceStack) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name       *types.String
		DiskGBytes *types.Int
		Policy     types.StringNull
		RawPolicy  types.TextNull
		CdnEnabled types.BoolNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostObjectStorageServiceStack", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if aux.DiskGBytes == nil {
		errorList = errorList.With(validator.NewError("diskGBytes", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.DiskGBytes = *aux.DiskGBytes
	dto.Policy = aux.Policy
	dto.RawPolicy = aux.RawPolicy
	dto.CdnEnabled = aux.CdnEnabled

	return nil
}
