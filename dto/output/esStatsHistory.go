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
	CpuUsedMin     types.Float      `json:"cpuUsedMin"`
	CpuUsedMedian  types.Float      `json:"cpuUsedMedian"`
	CpuUsedMax     types.Float      `json:"cpuUsedMax"`
	RamUsedMin     types.Float      `json:"ramUsedMin"`
	RamUsedMedian  types.Float      `json:"ramUsedMedian"`
	RamUsedMax     types.Float      `json:"ramUsedMax"`
	DiskUsedMin    types.Float      `json:"diskUsedMin"`
	DiskUsedMedian types.Float      `json:"diskUsedMedian"`
	DiskUsedMax    types.Float      `json:"diskUsedMax"`
}
