// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type DomainDnsCheckDetail struct {
	WildcardCnameCheckError types.TextNull                 `json:"wildcardCnameCheckError"`
	ResolvedNameServers     types.StringArray              `json:"resolvedNameServers"`
	ResolvedIpAddresses     types.StringArray              `json:"resolvedIpAddresses"`
	Challenges              DomainDnsCheckDetailChallenges `json:"challenges"`
}

type DomainDnsCheckDetailChallenges []DomainDnsCheckChallenge

func (dto DomainDnsCheckDetailChallenges) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]DomainDnsCheckChallenge(dto))
}
