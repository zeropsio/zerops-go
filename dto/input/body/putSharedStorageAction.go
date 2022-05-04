// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types/uuid"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutSharedStorageAction)(nil)

type PutSharedStorageAction struct {
	SharedStorageId uuid.ServiceStackId `json:"sharedStorageId"`
}

func (dto PutSharedStorageAction) GetSharedStorageId() uuid.ServiceStackId {
	return dto.SharedStorageId
}

func (dto *PutSharedStorageAction) UnmarshalJSON(b []byte) error {
	var aux = struct {
		SharedStorageId *uuid.ServiceStackId
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutSharedStorageAction", err)
	}
	var errorList validator.ErrorList
	if aux.SharedStorageId == nil {
		errorList = errorList.With(validator.NewError("sharedStorageId", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.SharedStorageId = *aux.SharedStorageId

	return nil
}
