// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type S3BucketItem struct {
	Name    types.String   `json:"name"`
	Created types.DateTime `json:"created"`
}
