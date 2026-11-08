// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ContainerTechnologyInfo struct {
	DetectedAt   types.DateTime                      `json:"detectedAt"`
	OsRelease    ContainerOsRelease                  `json:"osRelease"`
	Technologies ContainerTechnologyInfoTechnologies `json:"technologies"`
}

type ContainerTechnologyInfoTechnologies []ContainerTechnologyVersion

func (dto ContainerTechnologyInfoTechnologies) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ContainerTechnologyVersion(dto))
}
