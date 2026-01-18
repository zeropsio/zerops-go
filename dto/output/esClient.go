// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsClient struct {
	Id         uuid.ClientIdNull `json:"id"`
	Created    types.DateTime    `json:"created"`
	LastUpdate types.DateTime    `json:"lastUpdate"`
}
