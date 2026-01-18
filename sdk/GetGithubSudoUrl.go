// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type GetGithubSudoUrlResponse struct {
	success            output.GithubAuth
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetGithubSudoUrlResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetGithubSudoUrlResponse) Output() (output output.GithubAuth, err error) {
	return r.success, r.err
}

func (r GetGithubSudoUrlResponse) Err() error {
	return r.err
}
func (r GetGithubSudoUrlResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetGithubSudoUrlResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetGithubSudoUrl(ctx context.Context) (getGithubSudoUrlResponse GetGithubSudoUrlResponse, err error) {
	u := "/api/rest/public/github/sudo-url"

	var response GetGithubSudoUrlResponse
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
