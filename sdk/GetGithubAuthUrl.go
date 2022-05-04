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

type GetGithubAuthUrlResponse struct {
	success            output.GithubAuth
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetGithubAuthUrlResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetGithubAuthUrlResponse) Output() (output output.GithubAuth, err error) {
	return r.success, r.err
}

func (r GetGithubAuthUrlResponse) Err() error {
	return r.err
}
func (r GetGithubAuthUrlResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetGithubAuthUrlResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetGithubAuthUrl(ctx context.Context, inputDtoQuery query.GithubUrl) (getGithubAuthUrlResponse GetGithubAuthUrlResponse, err error) {
	u := "/api/rest/public/github/auth-url"

	var queryParams []string
	{
		param := inputDtoQuery.RedirectUrl.Native()
		queryParams = append(queryParams, "redirectUrl="+url.QueryEscape(param))
	}
	{
		param := inputDtoQuery.Action.Native()
		queryParams = append(queryParams, "action="+url.QueryEscape(param))
	}
	if param, ok := inputDtoQuery.HaRecipe.Get(); ok {
		queryParams = append(queryParams, "haRecipe="+url.QueryEscape(param.Native()))
	}
	if param, ok := inputDtoQuery.NonHaRecipe.Get(); ok {
		queryParams = append(queryParams, "nonHaRecipe="+url.QueryEscape(param.Native()))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetGithubAuthUrlResponse
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
