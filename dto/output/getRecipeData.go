// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type GetRecipeData struct {
	Name         types.String              `json:"name"`
	Content      types.String              `json:"content"`
	Extracts     types.Map                 `json:"extracts"`
	Environments GetRecipeDataEnvironments `json:"environments"`
	Errors       types.StringArrayNull     `json:"errors"`
}

type GetRecipeDataEnvironments []GetRecipeEnvironment

func (dto GetRecipeDataEnvironments) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]GetRecipeEnvironment(dto))
}
