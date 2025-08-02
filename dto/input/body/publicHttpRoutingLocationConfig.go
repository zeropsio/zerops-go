// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingLocationConfig)(nil)

type PublicHttpRoutingLocationConfig struct {
	Redirect     *PublicHttpRoutingLocationRedirect     `json:"redirect"`
	AccessPolicy *PublicHttpRoutingLocationAccessPolicy `json:"accessPolicy"`
	RateLimiting *PublicHttpRoutingLocationRateLimiting `json:"rateLimiting"`
	BasicAuth    *PublicHttpRoutingLocationBasicAuth    `json:"basicAuth"`
	Content      *PublicHttpRoutingLocationContent      `json:"content"`
}

func (dto PublicHttpRoutingLocationConfig) GetRedirect() *PublicHttpRoutingLocationRedirect {
	return dto.Redirect
}
func (dto PublicHttpRoutingLocationConfig) GetAccessPolicy() *PublicHttpRoutingLocationAccessPolicy {
	return dto.AccessPolicy
}
func (dto PublicHttpRoutingLocationConfig) GetRateLimiting() *PublicHttpRoutingLocationRateLimiting {
	return dto.RateLimiting
}
func (dto PublicHttpRoutingLocationConfig) GetBasicAuth() *PublicHttpRoutingLocationBasicAuth {
	return dto.BasicAuth
}
func (dto PublicHttpRoutingLocationConfig) GetContent() *PublicHttpRoutingLocationContent {
	return dto.Content
}

func (dto *PublicHttpRoutingLocationConfig) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Redirect     *PublicHttpRoutingLocationRedirect
		AccessPolicy *PublicHttpRoutingLocationAccessPolicy
		RateLimiting *PublicHttpRoutingLocationRateLimiting
		BasicAuth    *PublicHttpRoutingLocationBasicAuth
		Content      *PublicHttpRoutingLocationContent
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingLocationConfig", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.Redirect = aux.Redirect
	dto.AccessPolicy = aux.AccessPolicy
	dto.RateLimiting = aux.RateLimiting
	dto.BasicAuth = aux.BasicAuth
	dto.Content = aux.Content

	return nil
}
