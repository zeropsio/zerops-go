// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ObjectStorage struct {
	Id           uuid.ObjectStorageId `json:"id"`
	Name         types.String         `json:"name"`
	IsDefault    types.BoolNull       `json:"isDefault"`
	LocationId   stringId.LocationId  `json:"locationId"`
	LocationName types.String         `json:"locationName"`
}
