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

type ServiceStack struct {
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
	BuildCache                *BuildCache                        `json:"buildCache"`
	Ports                     ServiceStackPorts                  `json:"ports"`
	RequestedPorts            RequestedPortsJsonObject           `json:"requestedPorts"`
	Created                   types.DateTime                     `json:"created"`
	LastUpdate                types.DateTime                     `json:"lastUpdate"`
	Mode                      enum.ServiceStackModeEnum          `json:"mode"`
	CustomPortsEnabled        types.Bool                         `json:"customPortsEnabled"`
	SubdomainAccess           types.Bool                         `json:"subdomainAccess"`
	ReloadAvailable           types.Bool                         `json:"reloadAvailable"`
	VersionNumber             types.EmptyString                  `json:"versionNumber"`
	ProjectId                 uuid.ProjectId                     `json:"projectId"`
	Project                   ProjectLight                       `json:"project"`
	ConnectedStacks           ServiceStackConnectedStacks        `json:"connectedStacks"`
	UserData                  ServiceStackUserData               `json:"userData"`
	ActiveAppVersion          *GetAppVersion                     `json:"activeAppVersion"`
	CoreService               *ServiceStackLight                 `json:"coreService"`
}

type ServiceStackPorts []ServicePort

func (dto ServiceStackPorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServicePort(dto))
}

type ServiceStackConnectedStacks []ServiceStackConnectedServiceStack

func (dto ServiceStackConnectedStacks) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackConnectedServiceStack(dto))
}

type ServiceStackUserData []UserData

func (dto ServiceStackUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserData(dto))
}
