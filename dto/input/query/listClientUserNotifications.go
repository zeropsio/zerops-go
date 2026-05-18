// Generated ZEROPS sdk

package query

import (
	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

type ListClientUserNotifications struct {
	Limit          types.IntNull
	Offset         types.IntNull
	ProjectId      uuid.ProjectIdNull
	ServiceStackId uuid.ServiceStackIdNull
	Types          types.StringArrayNull
}
