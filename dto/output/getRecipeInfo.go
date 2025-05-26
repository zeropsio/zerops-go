// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type GetRecipeInfo struct {
	Url              types.String                  `json:"url"`
	ZeropsYaml       GetRecipeInfoZeropsYaml       `json:"zeropsYaml"`
	ZeropsImportYaml GetRecipeInfoZeropsImportYaml `json:"zeropsImportYaml"`
}

type GetRecipeInfoZeropsYaml []GetRecipeInfoItem

func (dto GetRecipeInfoZeropsYaml) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]GetRecipeInfoItem(dto))
}

type GetRecipeInfoZeropsImportYaml []GetRecipeInfoItem

func (dto GetRecipeInfoZeropsImportYaml) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]GetRecipeInfoItem(dto))
}
