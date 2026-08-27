// Generated ZEROPS sdk

package query

import (
	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

type ListServiceStackContainers struct {
	Limit    types.IntNull
	Offset   types.IntNull
	SortDir  enum.GenericListSortDirEnum
	Statuses types.StringArrayNull
}
