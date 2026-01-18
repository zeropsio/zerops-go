// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type PublicHttpRoutingDomain struct {
	DomainName     types.String                  `json:"domainName"`
	DnsCheckStatus enum.DomainDnsCheckStatusEnum `json:"dnsCheckStatus"`
	SslStatus      enum.DomainSslStatusEnum      `json:"sslStatus"`
	CdnStatus      enum.DomainCdnStatusEnum      `json:"cdnStatus"`
}
