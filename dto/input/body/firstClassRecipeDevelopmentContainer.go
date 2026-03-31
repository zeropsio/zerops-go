// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*FirstClassRecipeDevelopmentContainer)(nil)

type FirstClassRecipeDevelopmentContainer struct {
	ServiceImportYaml      types.Text     `json:"serviceImportYaml"`
	RecipeSource           types.TextNull `json:"recipeSource"`
	RecipeSourceUrl        types.TextNull `json:"recipeSourceUrl"`
	CreateIntegrationToken types.BoolNull `json:"createIntegrationToken"`
}

func (dto FirstClassRecipeDevelopmentContainer) GetServiceImportYaml() types.Text {
	return dto.ServiceImportYaml
}
func (dto FirstClassRecipeDevelopmentContainer) GetRecipeSource() types.TextNull {
	return dto.RecipeSource
}
func (dto FirstClassRecipeDevelopmentContainer) GetRecipeSourceUrl() types.TextNull {
	return dto.RecipeSourceUrl
}
func (dto FirstClassRecipeDevelopmentContainer) GetCreateIntegrationToken() types.BoolNull {
	return dto.CreateIntegrationToken
}

func (dto *FirstClassRecipeDevelopmentContainer) UnmarshalJSON(b []byte) error {
	var aux = struct {
		ServiceImportYaml      *types.Text
		RecipeSource           types.TextNull
		RecipeSourceUrl        types.TextNull
		CreateIntegrationToken types.BoolNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("FirstClassRecipeDevelopmentContainer", err)
	}
	var errorList validator.ErrorList
	if aux.ServiceImportYaml == nil {
		errorList = errorList.With(validator.NewError("serviceImportYaml", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.ServiceImportYaml = *aux.ServiceImportYaml
	dto.RecipeSource = aux.RecipeSource
	dto.RecipeSourceUrl = aux.RecipeSourceUrl
	dto.CreateIntegrationToken = aux.CreateIntegrationToken

	return nil
}
