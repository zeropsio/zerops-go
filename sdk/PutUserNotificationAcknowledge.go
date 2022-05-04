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

type PutUserNotificationAcknowledgeResponse struct {
	success            output.Success
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r PutUserNotificationAcknowledgeResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r PutUserNotificationAcknowledgeResponse) Output() (output output.Success, err error) {
	return r.success, r.err
}

func (r PutUserNotificationAcknowledgeResponse) Err() error {
	return r.err
}
func (r PutUserNotificationAcknowledgeResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r PutUserNotificationAcknowledgeResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) PutUserNotificationAcknowledge(ctx context.Context, inputDtoPath path.UserNotificationId) (putUserNotificationAcknowledgeResponse PutUserNotificationAcknowledgeResponse, err error) {
	u := "/api/rest/public/user-notification/" + inputDtoPath.Id.Native() + "/acknowledge"

	var response PutUserNotificationAcknowledgeResponse
	sdkResponse := sdkBase.Put(
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
