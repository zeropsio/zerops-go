// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type GetClientAllIntegrationTokenDelegationListResponse struct {
	success            output.ClientIntegrationTokenDelegationList
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetClientAllIntegrationTokenDelegationListResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetClientAllIntegrationTokenDelegationListResponse) Output() (output output.ClientIntegrationTokenDelegationList, err error) {
	return r.success, r.err
}

func (r GetClientAllIntegrationTokenDelegationListResponse) Err() error {
	return r.err
}
func (r GetClientAllIntegrationTokenDelegationListResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetClientAllIntegrationTokenDelegationListResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetClientAllIntegrationTokenDelegationList(ctx context.Context, inputDtoPath path.ClientId) (getClientAllIntegrationTokenDelegationListResponse GetClientAllIntegrationTokenDelegationListResponse, err error) {
	u := "/api/rest/public/client/" + inputDtoPath.Id.Native() + "/integration-token/delegation"

	var response GetClientAllIntegrationTokenDelegationListResponse
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
