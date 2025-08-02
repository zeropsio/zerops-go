// Generated ZEROPS sdk

package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"context"

	"github.com/zeropsio/zerops-go/apiError"
	"github.com/zeropsio/zerops-go/dto/input/body"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/dto/output"
	"github.com/zeropsio/zerops-go/sdkBase"
)

type PutPublicHttpRoutingLocationResponse struct {
	success            output.PublicHttpRouting
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PutPublicHttpRoutingLocationResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PutPublicHttpRoutingLocationResponse) Output() (output output.PublicHttpRouting, err error) {
	return r.success, r.err
}

func (r PutPublicHttpRoutingLocationResponse) Err() error {
	return r.err
}
func (r PutPublicHttpRoutingLocationResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PutPublicHttpRoutingLocationResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PutPublicHttpRoutingLocation(ctx context.Context, inputDtoPath path.PublicHttpRoutingId, inputDtoBody body.PublicHttpRoutingLocationPut) (putPublicHttpRoutingLocationResponse PutPublicHttpRoutingLocationResponse, err error) {
	u := "/api/rest/public/public-http-routing/" + inputDtoPath.Id.Native() + "/location"

	var response PutPublicHttpRoutingLocationResponse
	sdkResponse := sdkBase.Put(
		ctx,
		h.environment,
		u,
		inputDtoBody,
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
