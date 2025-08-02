// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type PublicHttpRoutingLocationConfig struct {
	Redirect     *PublicHttpRoutingLocationRedirect     `json:"redirect"`
	AccessPolicy *PublicHttpRoutingLocationAccessPolicy `json:"accessPolicy"`
	RateLimiting *PublicHttpRoutingLocationRateLimiting `json:"rateLimiting"`
	BasicAuth    *PublicHttpRoutingLocationBasicAuth    `json:"basicAuth"`
	Content      *PublicHttpRoutingLocationContent      `json:"content"`
}
