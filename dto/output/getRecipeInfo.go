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
	ReadmeMd         GetRecipeInfoReadmeMd         `json:"readmeMd"`
	ZeropsYaml       GetRecipeInfoZeropsYaml       `json:"zeropsYaml"`
	ZeropsImportYaml GetRecipeInfoZeropsImportYaml `json:"zeropsImportYaml"`
}

type GetRecipeInfoReadmeMd []GetRecipeInfoItem

func (dto GetRecipeInfoReadmeMd) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]GetRecipeInfoItem(dto))
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
