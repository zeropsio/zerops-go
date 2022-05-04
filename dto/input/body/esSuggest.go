// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*EsSuggest)(nil)

type EsSuggest struct {
	Search EsSuggestSearch  `json:"search"`
	Column types.StringNull `json:"column"`
	Text   types.StringNull `json:"text"`
}

func (dto EsSuggest) GetSearch() EsSuggestSearch {
	return dto.Search
}
func (dto EsSuggest) GetColumn() types.StringNull {
	return dto.Column
}
func (dto EsSuggest) GetText() types.StringNull {
	return dto.Text
}

type EsSuggestSearch []EsSearchItem

func (dto EsSuggestSearch) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSearchItem(dto))
}

func (dto *EsSuggest) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Search *EsSuggestSearch
		Column types.StringNull
		Text   types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsSuggest", err)
	}
	var errorList validator.ErrorList
	if aux.Search == nil {
		errorList = errorList.With(validator.NewError("search", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Search = *aux.Search
	dto.Column = aux.Column
	dto.Text = aux.Text

	return nil
}
