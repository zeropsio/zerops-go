// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
)

var _ strconv.NumError

type PublicHttpRoutingDomain struct {
	DomainName     types.String                                   `json:"domainName"`
	DnsCheckStatus enum.PublicHttpRoutingDomainDnsCheckStatusEnum `json:"dnsCheckStatus"`
	SslStatus      enum.PublicHttpRoutingDomainSslStatusEnum      `json:"sslStatus"`
}
