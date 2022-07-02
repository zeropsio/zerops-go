// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type EsStatsHistory struct {
	From           types.DateTime   `json:"from"`
	Till           types.DateTime   `json:"till"`
	ProjectId      types.String     `json:"projectId"`
	ServiceStackId types.String     `json:"serviceStackId"`
	ServiceId      types.StringNull `json:"serviceId"`
	ContainerId    types.StringNull `json:"containerId"`
	ContainerCount types.IntNull    `json:"containerCount"`
	CpuLimit       types.Float      `json:"cpuLimit"`
	CpuUsed        types.Float      `json:"cpuUsed"`
	RamLimit       types.Float      `json:"ramLimit"`
	RamUsed        types.Float      `json:"ramUsed"`
	DiskLimit      types.Float      `json:"diskLimit"`
	DiskUsed       types.Float      `json:"diskUsed"`
}
