// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsSuggest struct {
	Text    types.String     `json:"text"`
	Items   EsSuggestItems   `json:"items"`
	Columns EsSuggestColumns `json:"columns"`
}

type EsSuggestItems []EsSuggestItem

func (dto EsSuggestItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSuggestItem(dto))
}

type EsSuggestColumns []EsColumnsItem

func (dto EsSuggestColumns) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsColumnsItem(dto))
}
