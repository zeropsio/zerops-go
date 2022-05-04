// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*EsExport)(nil)

type EsExport struct {
	Search  EsExportSearch    `json:"search"`
	Columns types.StringArray `json:"columns"`
	Sort    EsExportSort      `json:"sort"`
	Format  types.StringNull  `json:"format"`
}

func (dto EsExport) GetSearch() EsExportSearch {
	return dto.Search
}
func (dto EsExport) GetColumns() types.StringArray {
	return dto.Columns
}
func (dto EsExport) GetSort() EsExportSort {
	return dto.Sort
}
func (dto EsExport) GetFormat() types.StringNull {
	return dto.Format
}

type EsExportSearch []EsSearchItem

func (dto EsExportSearch) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSearchItem(dto))
}

type EsExportSort []EsSortItem

func (dto EsExportSort) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSortItem(dto))
}

func (dto *EsExport) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Search  *EsExportSearch
		Columns *types.StringArray
		Sort    *EsExportSort
		Format  types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsExport", err)
	}
	var errorList validator.ErrorList
	if aux.Search == nil {
		errorList = errorList.With(validator.NewError("search", "field is required"))
	}
	if aux.Columns == nil {
		errorList = errorList.With(validator.NewError("columns", "field is required"))
	}
	if aux.Sort == nil {
		errorList = errorList.With(validator.NewError("sort", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Search = *aux.Search
	dto.Columns = *aux.Columns
	dto.Sort = *aux.Sort
	dto.Format = aux.Format

	return nil
}
