// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ExternalRepositoryIntegration)(nil)

type ExternalRepositoryIntegration struct {
	GithubIntegration *GithubIntegration `json:"githubIntegration"`
	GitlabIntegration *GitlabIntegration `json:"gitlabIntegration"`
}

func (dto ExternalRepositoryIntegration) GetGithubIntegration() *GithubIntegration {
	return dto.GithubIntegration
}
func (dto ExternalRepositoryIntegration) GetGitlabIntegration() *GitlabIntegration {
	return dto.GitlabIntegration
}

func (dto *ExternalRepositoryIntegration) UnmarshalJSON(b []byte) error {
	var aux = struct {
		GithubIntegration *GithubIntegration
		GitlabIntegration *GitlabIntegration
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ExternalRepositoryIntegration", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.GithubIntegration = aux.GithubIntegration
	dto.GitlabIntegration = aux.GitlabIntegration

	return nil
}
