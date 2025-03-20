// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type BuildCache struct {
	BuildServiceStackId uuid.ServiceStackId       `json:"buildServiceStackId"`
	SnapshotId          uuid.BuildCacheSnapshotId `json:"snapshotId"`
	ConfigChecksum      types.String              `json:"configChecksum"`
}
