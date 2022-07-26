// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"
)

var _ strconv.NumError

type ExternalRepositoryIntegration struct {
	GitlabIntegration *GitlabIntegration `json:"gitlabIntegration"`
	GithubIntegration *GithubIntegration `json:"githubIntegration"`
}
