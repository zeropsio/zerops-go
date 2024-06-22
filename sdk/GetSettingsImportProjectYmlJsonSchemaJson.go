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

type GetSettingsImportProjectYmlJsonSchemaJsonResponse struct {
	success            output.FileDownload
	err                error
	responseHeaders    http.Header
	responseStatusCode int
}

func (r GetSettingsImportProjectYmlJsonSchemaJsonResponse) OutputInterface() (output interface{}, err error) {
	return r.success, r.err
}

func (r GetSettingsImportProjectYmlJsonSchemaJsonResponse) Output() (output output.FileDownload, err error) {
	return r.success, r.err
}

func (r GetSettingsImportProjectYmlJsonSchemaJsonResponse) Err() error {
	return r.err
}
func (r GetSettingsImportProjectYmlJsonSchemaJsonResponse) Headers() http.Header {
	return r.responseHeaders
}

func (r GetSettingsImportProjectYmlJsonSchemaJsonResponse) StatusCode() int {
	return r.responseStatusCode
}

func (h Handler) GetSettingsImportProjectYmlJsonSchemaJson(ctx context.Context) (getSettingsImportProjectYmlJsonSchemaJsonResponse GetSettingsImportProjectYmlJsonSchemaJsonResponse, err error) {
	u := "/api/rest/public/settings/import-project-yml-json-schema.json"

	var response GetSettingsImportProjectYmlJsonSchemaJsonResponse
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
