// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"net/url"
	"strconv"
	"strings"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/query"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

var _ strconv.NumError

type GetGitlabRepositoryBranchResponse struct {
	success            output.VCSRepositoryBranchList
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetGitlabRepositoryBranchResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetGitlabRepositoryBranchResponse) Output() (output output.VCSRepositoryBranchList, err error) {
	return r.success, r.err
}

func (r GetGitlabRepositoryBranchResponse) Err() error {
	return r.err
}
func (r GetGitlabRepositoryBranchResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetGitlabRepositoryBranchResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetGitlabRepositoryBranch(ctx context.Context, inputDtoQuery query.RepositoryFullName) (getGitlabRepositoryBranchResponse GetGitlabRepositoryBranchResponse, err error) {
	u := "/api/rest/public/gitlab/repository-branch"

	var queryParams []string
	{
		param := inputDtoQuery.RepositoryFullName.Native()
		queryParams = append(queryParams, "repositoryFullName="+url.QueryEscape(param))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetGitlabRepositoryBranchResponse
	sdkResponse := sdkBase.Get(
		ctx,
		h.environment,
		u,
	)
	if sdkResponse.Err != nil {
		return response, sdkResponse.Err
	}
	response.responseHeaders = sdkResponse.HttpResponse.Header
	response.responseStatusCode = sdkResponse.HttpResponse.StatusCode

	decoder := json.NewDecoder(sdkResponse.ResponseData)
	if sdkResponse.HttpResponse.StatusCode < http.StatusMultipleChoices {
		if err := decoder.Decode(&response.success); err != nil {
			return response, err
		}
	} else {
		responseString := sdkResponse.ResponseData.String()
		apiErrorResponse := struct {
			Error apiError.Error `json:"error"`
		}{}
		err := decoder.Decode(&apiErrorResponse)
		if err != nil {
			return response, errors.New(sdkResponse.HttpResponse.Status + ": " + responseString)
		}
		apiErrorResponse.Error.HttpStatusCode = sdkResponse.HttpResponse.StatusCode
		response.err = apiErrorResponse.Error
	}

	return response, nil
}
