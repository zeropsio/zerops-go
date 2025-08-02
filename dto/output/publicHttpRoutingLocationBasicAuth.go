// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type PublicHttpRoutingLocationBasicAuth struct {
	Enabled types.Bool                              `json:"enabled"`
	Realm   types.EmptyString                       `json:"realm"`
	Users   PublicHttpRoutingLocationBasicAuthUsers `json:"users"`
}

type PublicHttpRoutingLocationBasicAuthUsers []PublicHttpRoutingLocationBasicAuthUser

func (dto PublicHttpRoutingLocationBasicAuthUsers) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PublicHttpRoutingLocationBasicAuthUser(dto))
}
