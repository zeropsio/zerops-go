// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type PublicHttpRoutingLocation struct {
	Path             types.String             `json:"path"`
	Port             types.Int                `json:"port"`
	ServiceStackId   uuid.ServiceStackId      `json:"serviceStackId"`
	ServiceStackInfo LocationServiceStackInfo `json:"serviceStackInfo"`
}
