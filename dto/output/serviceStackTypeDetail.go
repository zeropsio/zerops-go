// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/stringId"
)

var _ strconv.NumError

type ServiceStackTypeDetail struct {
	Id                          stringId.ServiceStackTypeId                       `json:"id"`
	Name                        types.String                                      `json:"name"`
	Description                 types.Text                                        `json:"description"`
	Created                     types.DateTime                                    `json:"created"`
	LastUpdate                  types.DateTime                                    `json:"lastUpdate"`
	Category                    enum.ServiceStackTypeCategoryEnum                 `json:"category"`
	Subcategory                 types.String                                      `json:"subcategory"`
	DocsUrl                     types.String                                      `json:"docsUrl"`
	IsBuild                     types.Bool                                        `json:"isBuild"`
	IsRuntime                   types.Bool                                        `json:"isRuntime"`
	IsManaged                   types.Bool                                        `json:"isManaged"`
	HasBackup                   types.Bool                                        `json:"hasBackup"`
	HasAccessDetails            types.Bool                                        `json:"hasAccessDetails"`
	HasConfiguration            types.Bool                                        `json:"hasConfiguration"`
	OsList                      types.StringArray                                 `json:"osList"`
	ModeList                    types.StringArray                                 `json:"modeList"`
	ServiceStackTypeVersionList ServiceStackTypeDetailServiceStackTypeVersionList `json:"serviceStackTypeVersionList"`
}

type ServiceStackTypeDetailServiceStackTypeVersionList []ServiceStackTypeVersion

func (dto ServiceStackTypeDetailServiceStackTypeVersionList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackTypeVersion(dto))
}
