// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type PostAppVersion struct {
	Id             uuid.AppVersionId          `json:"id"`
	ClientId       uuid.ClientId              `json:"clientId"`
	ProjectId      uuid.ProjectId             `json:"projectId"`
	ServiceStackId uuid.ServiceStackId        `json:"serviceStackId"`
	Build          *AppVersionBuild           `json:"build"`
	Sequence       types.Int                  `json:"sequence"`
	Status         enum.AppVersionStatusEnum  `json:"status"`
	UserDataList   PostAppVersionUserDataList `json:"userDataList"`
	Created        types.DateTime             `json:"created"`
	LastUpdate     types.DateTime             `json:"lastUpdate"`
	UploadUrl      types.String               `json:"uploadUrl"`
}

type PostAppVersionUserDataList []AppVersionUserData

func (dto PostAppVersionUserDataList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AppVersionUserData(dto))
}
