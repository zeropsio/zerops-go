// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type S3BucketItems struct {
	Items S3BucketItemsItems `json:"items"`
}

type S3BucketItemsItems []S3BucketItem

func (dto S3BucketItemsItems) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]S3BucketItem(dto))
}
