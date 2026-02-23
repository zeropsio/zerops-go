// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type DomainCertificateInfo struct {
	ValidFrom        types.DateTimeNull `json:"validFrom"`
	ValidTill        types.DateTimeNull `json:"validTill"`
	RegisteredDomain types.String       `json:"registeredDomain"`
	DomainList       types.StringArray  `json:"domainList"`
}
