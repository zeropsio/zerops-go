// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsUserNotificationResponse struct {
	Limit     types.Int                       `json:"limit"`
	Offset    types.Int                       `json:"offset"`
	TotalHits types.Int                       `json:"totalHits"`
	Items     EsUserNotificationResponseItems `json:"items"`
}

type EsUserNotificationResponseItems []EsUserNotification

func (dto EsUserNotificationResponseItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsUserNotification(dto))
}
