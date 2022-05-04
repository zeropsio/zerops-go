// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type AppVersionGitlabIntegration struct {
	EventType          enum.AppVersionGitlabIntegrationEventTypeEnum `json:"eventType"`
	BranchName         types.StringNull                              `json:"branchName"`
	Pusher             types.String                                  `json:"pusher"`
	Commit             types.String                                  `json:"commit"`
	RepositoryFullName types.String                                  `json:"repositoryFullName"`
	TagName            types.StringNull                              `json:"tagName"`
}
