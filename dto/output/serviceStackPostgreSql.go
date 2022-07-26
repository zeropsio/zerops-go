// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ServiceStackPostgreSql struct {
	Id                                 uuid.ServiceStackId                   `json:"id"`
	Status                             enum.ServiceStackStatusEnum           `json:"status"`
	Name                               types.String                          `json:"name"`
	ServiceStackTypeInfo               ServiceStackInfoJsonObject            `json:"serviceStackTypeInfo"`
	ServiceStackTypeId                 stringId.ServiceStackTypeId           `json:"serviceStackTypeId"`
	ServiceStackTypeVersionId          stringId.ServiceStackTypeVersionId    `json:"serviceStackTypeVersionId"`
	IsSystem                           types.Bool                            `json:"isSystem"`
	StartOnProjectStart                types.Bool                            `json:"startOnProjectStart"`
	GithubIntegration                  *GithubIntegration                    `json:"githubIntegration"`
	GitlabIntegration                  *GitlabIntegration                    `json:"gitlabIntegration"`
	CustomAutoscaling                  *CustomAutoscaling                    `json:"customAutoscaling"`
	Ports                              ServiceStackPostgreSqlPorts           `json:"ports"`
	RequestedPorts                     RequestedPortsJsonObject              `json:"requestedPorts"`
	Created                            types.DateTime                        `json:"created"`
	LastUpdate                         types.DateTime                        `json:"lastUpdate"`
	Mode                               enum.ServiceStackModeEnum             `json:"mode"`
	CustomPortsEnabled                 types.Bool                            `json:"customPortsEnabled"`
	SubdomainAccess                    types.Bool                            `json:"subdomainAccess"`
	ReloadAvailable                    types.Bool                            `json:"reloadAvailable"`
	VersionNumber                      types.EmptyString                     `json:"versionNumber"`
	ProjectId                          uuid.ProjectId                        `json:"projectId"`
	Project                            ProjectLight                          `json:"project"`
	ConnectedStacks                    ServiceStackPostgreSqlConnectedStacks `json:"connectedStacks"`
	UserData                           ServiceStackPostgreSqlUserData        `json:"userData"`
	ActiveAppVersion                   *AppVersionLight                      `json:"activeAppVersion"`
	HasUnsyncedUserDataRecord          types.Bool                            `json:"hasUnsyncedUserDataRecord"`
	HasUnsyncedPublicHttpRoutingRecord types.Bool                            `json:"hasUnsyncedPublicHttpRoutingRecord"`
	HasUnsyncedPublicPortRecord        types.Bool                            `json:"hasUnsyncedPublicPortRecord"`
	HasPublicPortRoutingAccess         types.Bool                            `json:"hasPublicPortRoutingAccess"`
	HasPublicHttpRoutingAccess         types.Bool                            `json:"hasPublicHttpRoutingAccess"`
	ActivePublicHttpRoutingCount       types.IntNull                         `json:"activePublicHttpRoutingCount"`
	ActivePublicPortRoutingCount       types.IntNull                         `json:"activePublicPortRoutingCount"`
	User                               types.String                          `json:"user"`
	Password                           types.String                          `json:"password"`
}

type ServiceStackPostgreSqlPorts []ServicePort

func (dto ServiceStackPostgreSqlPorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServicePort(dto))
}

type ServiceStackPostgreSqlConnectedStacks []ServiceStackConnectedServiceStack

func (dto ServiceStackPostgreSqlConnectedStacks) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackConnectedServiceStack(dto))
}

type ServiceStackPostgreSqlUserData []UserDataLight

func (dto ServiceStackPostgreSqlUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataLight(dto))
}
