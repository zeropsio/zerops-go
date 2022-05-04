// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type AppVersionJsonObject struct {
	Id                   uuid.AppVersionId            `json:"id"`
	ServiceStackId       uuid.ServiceStackIdNull      `json:"serviceStackId"`
	ProjectId            uuid.ProjectIdNull           `json:"projectId"`
	Status               *enum.AppVersionStatusEnum   `json:"status"`
	Source               enum.AppVersionSourceEnum    `json:"source"`
	Sequence             types.Int                    `json:"sequence"`
	Name                 types.StringNull             `json:"name"`
	Created              types.DateTimeNull           `json:"created"`
	Build                *AppVersionBuild             `json:"build"`
	PrepareCustomRuntime *PrepareCustomRuntime        `json:"prepareCustomRuntime"`
	GithubIntegration    *AppVersionGithubIntegration `json:"githubIntegration"`
	GitlabIntegration    *AppVersionGitlabIntegration `json:"gitlabIntegration"`
	CreatedByUser        *UserJsonObject              `json:"createdByUser"`
	ActivationDate       types.DateTimeNull           `json:"activationDate"`
	PublicGitSource      *AppVersionPublicGitSource   `json:"publicGitSource"`
	ConfigContent        types.TextNull               `json:"configContent"`
}
