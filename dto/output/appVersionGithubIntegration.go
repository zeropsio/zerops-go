// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type AppVersionGithubIntegration struct {
	EventType          enum.AppVersionGithubIntegrationEventTypeEnum `json:"eventType"`
	BranchName         types.StringNull                              `json:"branchName"`
	TagRegex           types.StringNull                              `json:"tagRegex"`
	Pusher             types.String                                  `json:"pusher"`
	Commit             types.String                                  `json:"commit"`
	RepositoryFullName types.String                                  `json:"repositoryFullName"`
	TagName            types.StringNull                              `json:"tagName"`
	ZeropsYamlSetup    types.StringNull                              `json:"zeropsYamlSetup"`
}
