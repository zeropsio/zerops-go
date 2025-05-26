// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutServiceStackObjectStorageSize)(nil)

type PutServiceStackObjectStorageSize struct {
	DiskGBytes types.Int        `json:"diskGBytes"`
	Policy     types.StringNull `json:"policy"`
	RawPolicy  types.TextNull   `json:"rawPolicy"`
	CdnEnabled types.BoolNull   `json:"cdnEnabled"`
}

func (dto PutServiceStackObjectStorageSize) GetDiskGBytes() types.Int {
	return dto.DiskGBytes
}
func (dto PutServiceStackObjectStorageSize) GetPolicy() types.StringNull {
	return dto.Policy
}
func (dto PutServiceStackObjectStorageSize) GetRawPolicy() types.TextNull {
	return dto.RawPolicy
}
func (dto PutServiceStackObjectStorageSize) GetCdnEnabled() types.BoolNull {
	return dto.CdnEnabled
}

func (dto *PutServiceStackObjectStorageSize) UnmarshalJSON(b []byte) error {
	var aux = struct {
		DiskGBytes *types.Int
		Policy     types.StringNull
		RawPolicy  types.TextNull
		CdnEnabled types.BoolNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutServiceStackObjectStorageSize", err)
	}
	var errorList validator.ErrorList
	if aux.DiskGBytes == nil {
		errorList = errorList.With(validator.NewError("diskGBytes", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.DiskGBytes = *aux.DiskGBytes
	dto.Policy = aux.Policy
	dto.RawPolicy = aux.RawPolicy
	dto.CdnEnabled = aux.CdnEnabled

	return nil
}
