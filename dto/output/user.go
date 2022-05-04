// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type User struct {
	Id                 uuid.UserId         `json:"id"`
	Email              types.Email         `json:"email"`
	FullName           types.String        `json:"fullName"`
	FirstName          types.String        `json:"firstName"`
	LastName           types.String        `json:"lastName"`
	Avatar             *UserAvatar         `json:"avatar"`
	CountryCallingCode types.IntNull       `json:"countryCallingCode"`
	PhoneNumber        types.IntNull       `json:"phoneNumber"`
	Language           Language            `json:"language"`
	Created            types.DateTime      `json:"created"`
	LastUpdate         types.DateTime      `json:"lastUpdate"`
	Status             enum.UserStatusEnum `json:"status"`
	ClientUserList     UserClientUserList  `json:"clientUserList"`
}

type UserClientUserList []ClientUserExtra

func (dto UserClientUserList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ClientUserExtra(dto))
}
