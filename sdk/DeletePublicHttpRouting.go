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

type DeletePublicHttpRoutingResponse struct {
	success            output.DeletePublicHttpRouting
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r DeletePublicHttpRoutingResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r DeletePublicHttpRoutingResponse) Output() (output output.DeletePublicHttpRouting, err error) {
	return r.success, r.err
}

func (r DeletePublicHttpRoutingResponse) Err() error {
	return r.err
}
func (r DeletePublicHttpRoutingResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r DeletePublicHttpRoutingResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) DeletePublicHttpRouting(ctx context.Context, inputDtoPath path.PublicHttpRoutingId) (deletePublicHttpRoutingResponse DeletePublicHttpRoutingResponse, err error) {
	u := "/api/rest/public/public-http-routing/" + inputDtoPath.Id.Native() + ""

	var response DeletePublicHttpRoutingResponse
	sdkResponse := sdkBase.Delete(
		ctx,
		h.environment,
		u,
		struct{}{},
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
