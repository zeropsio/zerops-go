// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ServiceStackProcessMariaDb struct {
	Id                                 uuid.ServiceStackId                       `json:"id"`
	Status                             enum.ServiceStackStatusEnum               `json:"status"`
	Name                               types.String                              `json:"name"`
	ServiceStackTypeInfo               ServiceStackInfoJsonObject                `json:"serviceStackTypeInfo"`
	ServiceStackTypeId                 stringId.ServiceStackTypeId               `json:"serviceStackTypeId"`
	ServiceStackTypeVersionId          stringId.ServiceStackTypeVersionId        `json:"serviceStackTypeVersionId"`
	IsSystem                           types.Bool                                `json:"isSystem"`
	StartOnProjectStart                types.Bool                                `json:"startOnProjectStart"`
	GithubIntegration                  *GithubIntegration                        `json:"githubIntegration"`
	GitlabIntegration                  *GitlabIntegration                        `json:"gitlabIntegration"`
	CustomAutoscaling                  *CustomAutoscaling                        `json:"customAutoscaling"`
	Ports                              ServiceStackProcessMariaDbPorts           `json:"ports"`
	RequestedPorts                     RequestedPortsJsonObject                  `json:"requestedPorts"`
	Created                            types.DateTime                            `json:"created"`
	LastUpdate                         types.DateTime                            `json:"lastUpdate"`
	Mode                               enum.ServiceStackModeEnum                 `json:"mode"`
	CustomPortsEnabled                 types.Bool                                `json:"customPortsEnabled"`
	SubdomainAccess                    types.Bool                                `json:"subdomainAccess"`
	ReloadAvailable                    types.Bool                                `json:"reloadAvailable"`
	VersionNumber                      types.EmptyString                         `json:"versionNumber"`
	ProjectId                          uuid.ProjectId                            `json:"projectId"`
	Project                            ProjectLight                              `json:"project"`
	ConnectedStacks                    ServiceStackProcessMariaDbConnectedStacks `json:"connectedStacks"`
	UserData                           ServiceStackProcessMariaDbUserData        `json:"userData"`
	ActiveAppVersion                   *AppVersionLight                          `json:"activeAppVersion"`
	HasUnsyncedUserDataRecord          types.Bool                                `json:"hasUnsyncedUserDataRecord"`
	HasUnsyncedPublicHttpRoutingRecord types.Bool                                `json:"hasUnsyncedPublicHttpRoutingRecord"`
	HasUnsyncedPublicPortRecord        types.Bool                                `json:"hasUnsyncedPublicPortRecord"`
	HasPublicPortRoutingAccess         types.Bool                                `json:"hasPublicPortRoutingAccess"`
	HasPublicHttpRoutingAccess         types.Bool                                `json:"hasPublicHttpRoutingAccess"`
	ActivePublicHttpRoutingCount       types.IntNull                             `json:"activePublicHttpRoutingCount"`
	ActivePublicPortRoutingCount       types.IntNull                             `json:"activePublicPortRoutingCount"`
	User                               types.String                              `json:"user"`
	Password                           types.String                              `json:"password"`
	Process                            Process                                   `json:"process"`
}

type ServiceStackProcessMariaDbPorts []ServicePort

func (dto ServiceStackProcessMariaDbPorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServicePort(dto))
}

type ServiceStackProcessMariaDbConnectedStacks []ServiceStackConnectedServiceStack

func (dto ServiceStackProcessMariaDbConnectedStacks) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackConnectedServiceStack(dto))
}

type ServiceStackProcessMariaDbUserData []UserDataLight

func (dto ServiceStackProcessMariaDbUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataLight(dto))
}
