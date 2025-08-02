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

type GetProjectVpnListResponse struct {
	success            output.ProjectVpnList
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetProjectVpnListResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetProjectVpnListResponse) Output() (output output.ProjectVpnList, err error) {
	return r.success, r.err
}

func (r GetProjectVpnListResponse) Err() error {
	return r.err
}
func (r GetProjectVpnListResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetProjectVpnListResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetProjectVpnList(ctx context.Context, inputDtoPath path.ProjectId, inputDtoQuery query.GetProjectVpn) (getProjectVpnListResponse GetProjectVpnListResponse, err error) {
	u := "/api/rest/public/project/" + inputDtoPath.Id.Native() + "/vpn/list"

	var queryParams []string
	if param, ok := inputDtoQuery.InstanceId.Get(); ok {
		queryParams = append(queryParams, "instanceId="+url.QueryEscape(param.Native()))
	}

	if len(queryParams) > 0 {
		u += "?" + strings.Join(queryParams, "&")
	}

	var response GetProjectVpnListResponse
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
