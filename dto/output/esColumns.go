// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsColumns struct {
	Formats     types.StringArray    `json:"formats"`
	DefaultSort EsColumnsDefaultSort `json:"defaultSort"`
	Items       EsColumnsItems       `json:"items"`
}

type EsColumnsItems []EsColumnsItem

func (dto EsColumnsItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsColumnsItem(dto))
}
