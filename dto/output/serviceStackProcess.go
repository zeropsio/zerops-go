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

type ServiceStackProcess struct {
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
	CdnEnabled                types.Bool                         `json:"cdnEnabled"`
	NestId                    stringId.NestId                    `json:"nestId"`
	Ports                     ServiceStackProcessPorts           `json:"ports"`
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
	ConnectedStacks           ServiceStackProcessConnectedStacks `json:"connectedStacks"`
	UserData                  ServiceStackProcessUserData        `json:"userData"`
	ActiveAppVersion          *GetAppVersion                     `json:"activeAppVersion"`
	CoreService               *ServiceStackLight                 `json:"coreService"`
	Process                   Process                            `json:"process"`
}

type ServiceStackProcessPorts []ServicePort

func (dto ServiceStackProcessPorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServicePort(dto))
}

type ServiceStackProcessConnectedStacks []ServiceStackConnectedServiceStack

func (dto ServiceStackProcessConnectedStacks) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackConnectedServiceStack(dto))
}

type ServiceStackProcessUserData []UserData

func (dto ServiceStackProcessUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserData(dto))
}
