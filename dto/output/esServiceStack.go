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

type EsServiceStack struct {
	Id                                 uuid.ServiceStackId                `json:"id"`
	ProjectId                          uuid.ProjectId                     `json:"projectId"`
	ClientId                           uuid.ClientId                      `json:"clientId"`
	ServiceStackTypeId                 stringId.ServiceStackTypeId        `json:"serviceStackTypeId"`
	ServiceStackTypeVersionId          stringId.ServiceStackTypeVersionId `json:"serviceStackTypeVersionId"`
	DriverId                           stringId.DriverIdNull              `json:"driverId"`
	ServiceStackTypeInfo               ServiceStackInfoJsonObject         `json:"serviceStackTypeInfo"`
	Status                             enum.ServiceStackStatusEnum        `json:"status"`
	Name                               types.String                       `json:"name"`
	Created                            types.DateTime                     `json:"created"`
	LastUpdate                         types.DateTime                     `json:"lastUpdate"`
	Ports                              EsServiceStackPorts                `json:"ports"`
	RequestedPorts                     RequestedPortsJsonObject           `json:"requestedPorts"`
	IsSystem                           types.Bool                         `json:"isSystem"`
	Mode                               enum.ServiceStackModeEnum          `json:"mode"`
	SubdomainAccess                    types.Bool                         `json:"subdomainAccess"`
	VersionNumber                      types.EmptyString                  `json:"versionNumber"`
	ReloadAvailable                    types.Bool                         `json:"reloadAvailable"`
	Project                            ProjectLight                       `json:"project"`
	UserData                           EsServiceStackUserData             `json:"userData"`
	ConnectedStacks                    EsServiceStackConnectedStacks      `json:"connectedStacks"`
	ActiveAppVersion                   *AppVersionLight                   `json:"activeAppVersion"`
	HasUnsyncedUserDataRecord          types.Bool                         `json:"hasUnsyncedUserDataRecord"`
	HasUnsyncedPublicHttpRoutingRecord types.Bool                         `json:"hasUnsyncedPublicHttpRoutingRecord"`
	HasUnsyncedPublicPortRecord        types.Bool                         `json:"hasUnsyncedPublicPortRecord"`
	HasPublicHttpRoutingAccess         types.Bool                         `json:"hasPublicHttpRoutingAccess"`
	HasPublicPortRoutingAccess         types.Bool                         `json:"hasPublicPortRoutingAccess"`
	ActivePublicHttpRoutingCount       types.IntNull                      `json:"activePublicHttpRoutingCount"`
	ActivePublicPortRoutingCount       types.IntNull                      `json:"activePublicPortRoutingCount"`
	StartOnProjectStart                types.Bool                         `json:"startOnProjectStart"`
	GithubIntegration                  *GithubIntegration                 `json:"githubIntegration"`
	CustomAutoscaling                  *CustomAutoscaling                 `json:"customAutoscaling"`
}

type EsServiceStackPorts []ServicePort

func (dto EsServiceStackPorts) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServicePort(dto))
}

type EsServiceStackUserData []UserDataLight

func (dto EsServiceStackUserData) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]UserDataLight(dto))
}

type EsServiceStackConnectedStacks []ServiceStackConnectedServiceStack

func (dto EsServiceStackConnectedStacks) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackConnectedServiceStack(dto))
}
