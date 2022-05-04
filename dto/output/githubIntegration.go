// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type GithubIntegration struct {
	UserId                types.String                        `json:"userId"`
	AuthorizationRequired types.Bool                          `json:"authorizationRequired"`
	BranchName            types.StringNull                    `json:"branchName"`
	EventType             enum.GithubIntegrationEventTypeEnum `json:"eventType"`
	IsActive              types.Bool                          `json:"isActive"`
	RepositoryFullName    types.String                        `json:"repositoryFullName"`
	WebhookId             types.IntNull                       `json:"webhookId"`
}
