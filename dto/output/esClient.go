// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type EsClient struct {
	Id                   uuid.ClientId  `json:"id"`
	InstanceGroupCreated types.Bool     `json:"instanceGroupCreated"`
	Created              types.DateTime `json:"created"`
	LastUpdate           types.DateTime `json:"lastUpdate"`
}
