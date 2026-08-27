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
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/input/query"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

var _ strconv.NumError

type GetProjectEnvFileForEditResponse struct {
	success            output.EnvFile
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetProjectEnvFileForEditResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetProjectEnvFileForEditResponse) Output() (output output.EnvFile, err error) {
	return r.success, r.err
}

func (r GetProjectEnvFileForEditResponse) Err() error {
	return r.err
}
func (r GetProjectEnvFileForEditResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetProjectEnvFileForEditResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetProjectEnvFileForEdit(ctx context.Context, inputDtoPath path.ProjectId, inputDtoQuery query.GetProjectEnvFileForEdit) (getProjectEnvFileForEditResponse GetProjectEnvFileForEditResponse, err error) {
	u := "/api/rest/public/project/" + inputDtoPath.Id.Native() + "/project-env/env-file"

	var queryParams []string
	{
		param := inputDtoQuery.Sensitive.Native()
		queryParams = append(queryParams, "sensitive="+url.QueryEscape(strconv.FormatBool(param)))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetProjectEnvFileForEditResponse
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
