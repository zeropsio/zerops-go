// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ServiceStackLight struct {
	Id                        uuid.ServiceStackId                `json:"id"`
	Status                    enum.ServiceStackStatusEnum        `json:"status"`
	Name                      types.String                       `json:"name"`
	ServiceStackTypeInfo      ServiceStackInfoJsonObject         `json:"serviceStackTypeInfo"`
	ServiceStackTypeId        stringId.ServiceStackTypeId        `json:"serviceStackTypeId"`
	ServiceStackTypeVersionId stringId.ServiceStackTypeVersionId `json:"serviceStackTypeVersionId"`
	IsSystem                  types.Bool                         `json:"isSystem"`
	StartOnProjectStart       types.Bool                         `json:"startOnProjectStart"`
	GithubIntegration         *GithubIntegration                 `json:"githubIntegration"`
	GitlabIntegration         *GitlabIntegration                 `json:"gitlabIntegration"`
	CustomAutoscaling         *CustomAutoscaling                 `json:"customAutoscaling"`
}
