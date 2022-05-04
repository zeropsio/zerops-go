// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*EsSortItem)(nil)

type EsSortItem struct {
	Name      types.String   `json:"name"`
	Ascending types.BoolNull `json:"ascending"`
}

func (dto EsSortItem) GetName() types.String {
	return dto.Name
}
func (dto EsSortItem) GetAscending() types.BoolNull {
	return dto.Ascending
}

func (dto *EsSortItem) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Name      *types.String
		Ascending types.BoolNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsSortItem", err)
	}
	var errorList validator.ErrorList
	if aux.Name == nil {
		errorList = errorList.With(validator.NewError("name", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Name = *aux.Name
	dto.Ascending = aux.Ascending

	return nil
}
