// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type UserNotificationList struct {
	List       UserNotificationListList `json:"list"`
	TotalCount types.Int                `json:"totalCount"`
}

type UserNotificationListList []UserNotification

func (dto UserNotificationListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserNotification(dto))
}
