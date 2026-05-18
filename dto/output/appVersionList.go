// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type AppVersionList struct {
	List       AppVersionListList `json:"list"`
	TotalCount types.Int          `json:"totalCount"`
}

type AppVersionListList []GetAppVersion

func (dto AppVersionListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]GetAppVersion(dto))
}
