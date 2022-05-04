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

type EsAppVersion struct {
	Id                uuid.AppVersionId            `json:"id"`
	ClientId          uuid.ClientId                `json:"clientId"`
	ProjectId         uuid.ProjectId               `json:"projectId"`
	ServiceStackId    uuid.ServiceStackId          `json:"serviceStackId"`
	Build             *AppVersionBuild             `json:"build"`
	Source            enum.AppVersionSourceEnum    `json:"source"`
	Sequence          types.Int                    `json:"sequence"`
	Status            enum.AppVersionStatusEnum    `json:"status"`
	UserData          EsAppVersionUserData         `json:"userData"`
	GithubIntegration *AppVersionGithubIntegration `json:"githubIntegration"`
	PublicGitSource   *AppVersionPublicGitSource   `json:"publicGitSource"`
	Created           types.DateTime               `json:"created"`
	LastUpdate        types.DateTime               `json:"lastUpdate"`
}

type EsAppVersionUserData []AppVersionUserData

func (dto EsAppVersionUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AppVersionUserData(dto))
}
