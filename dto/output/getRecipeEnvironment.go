// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type GetRecipeEnvironment struct {
	Name        types.String                 `json:"name"`
	Content     types.String                 `json:"content"`
	Extracts    types.Map                    `json:"extracts"`
	Import      types.String                 `json:"import"`
	ProjectName types.String                 `json:"projectName"`
	Services    GetRecipeEnvironmentServices `json:"services"`
}

type GetRecipeEnvironmentServices []GetRecipeService

func (dto GetRecipeEnvironmentServices) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]GetRecipeService(dto))
}
