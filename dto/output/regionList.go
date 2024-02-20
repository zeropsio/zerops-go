// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type RegionList struct {
	Items RegionListItems `json:"items"`
}

type RegionListItems []Region

func (dto RegionListItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Region(dto))
}
