// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ProjectBackupList struct {
	Count  types.Int             `json:"count"`
	Total  types.Int             `json:"total"`
	Limit  types.Int             `json:"limit"`
	Offset types.Int             `json:"offset"`
	List   ProjectBackupListList `json:"list"`
}

type ProjectBackupListList []ProjectBackup

func (dto ProjectBackupListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProjectBackup(dto))
}
