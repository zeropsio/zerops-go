package sdkBase

import (
	"bytes"
	"net/http"
)

type Response struct {
	Err          error
	HttpResponse *http.Response
	ResponseData *bytes.Buffer
	Request      *http.Request
}
