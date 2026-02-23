// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type DomainDnsCheckChallenge struct {
	ChallengeUrl   types.String   `json:"challengeUrl"`
	ChallengeHost  types.String   `json:"challengeHost"`
	HttpStatusCode types.IntNull  `json:"httpStatusCode"`
	HttpResponse   types.TextNull `json:"httpResponse"`
	Error          types.TextNull `json:"error"`
}
