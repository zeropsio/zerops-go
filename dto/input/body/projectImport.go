// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ProjectImport)(nil)

type ProjectImport struct {
	Yaml         types.Text     `json:"yaml"`
	RecipeSource types.TextNull `json:"recipeSource"`
}

func (dto ProjectImport) GetYaml() types.Text {
	return dto.Yaml
}
func (dto ProjectImport) GetRecipeSource() types.TextNull {
	return dto.RecipeSource
}

func (dto *ProjectImport) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Yaml         *types.Text
		RecipeSource types.TextNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ProjectImport", err)
	}
	var errorList validator.ErrorList
	if aux.Yaml == nil {
		errorList = errorList.With(validator.NewError("yaml", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Yaml = *aux.Yaml
	dto.RecipeSource = aux.RecipeSource

	return nil
}
