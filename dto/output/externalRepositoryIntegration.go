// Generated ZEROPS sdk

package output

import (
	"strconv"
)

var _ strconv.NumError

type ExternalRepositoryIntegration struct {
	GitlabIntegration *GitlabIntegration `json:"gitlabIntegration"`
	GithubIntegration *GithubIntegration `json:"githubIntegration"`
}
