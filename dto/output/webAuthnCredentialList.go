// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type WebAuthnCredentialList struct {
	List WebAuthnCredentialListList `json:"list"`
}

type WebAuthnCredentialListList []WebAuthnCredential

func (dto WebAuthnCredentialListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]WebAuthnCredential(dto))
}
