// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ServiceStackImport)(nil)

type ServiceStackImport struct {
	Yaml            types.Text     `json:"yaml"`
	RecipeSource    types.TextNull `json:"recipeSource"`
	RecipeSourceUrl types.TextNull `json:"recipeSourceUrl"`
}

func (dto ServiceStackImport) GetYaml() types.Text {
	return dto.Yaml
}
func (dto ServiceStackImport) GetRecipeSource() types.TextNull {
	return dto.RecipeSource
}
func (dto ServiceStackImport) GetRecipeSourceUrl() types.TextNull {
	return dto.RecipeSourceUrl
}

func (dto *ServiceStackImport) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Yaml            *types.Text
		RecipeSource    types.TextNull
		RecipeSourceUrl types.TextNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ServiceStackImport", err)
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
	dto.RecipeSourceUrl = aux.RecipeSourceUrl

	return nil
}
