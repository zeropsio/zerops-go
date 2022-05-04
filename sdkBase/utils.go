package sdkBase

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func Get(ctx context.Context, e Environment, url string) *Response {
	return Method(ctx, e, http.MethodGet, url, nil)
}

func Put(ctx context.Context, e Environment, url string, in interface{}) *Response {
	return MethodBody(ctx, e, http.MethodPut, url, in)
}

func Post(ctx context.Context, e Environment, url string, in interface{}) *Response {
	return MethodBody(ctx, e, http.MethodPost, url, in)
}

func Delete(ctx context.Context, e Environment, url string, in interface{}) *Response {
	return MethodBody(ctx, e, http.MethodDelete, url, in)
}

func MethodBody(ctx context.Context, e Environment, method string, url string, in interface{}) (r *Response) {
	r = &Response{}
	buf := bytes.NewBuffer(nil)
	if r.Err = json.NewEncoder(buf).Encode(in); r.Err != nil {
		return
	}
	return Method(ctx, e, method, url, buf)
}

func Method(ctx context.Context, e Environment, method string, url string, body io.Reader) (r *Response) {
	r = &Response{}
	if r.Request, r.Err = e.Request(ctx, method, url, body); r.Err != nil {
		return
	}

	resp, err := e.Do(r.Request)
	if err != nil {
		r.Err = err
		return
	}
	defer resp.Body.Close()

	// TODO - odstranit httpResponse https://redmine.vshosting.cz/issues/17052
	r.HttpResponse = resp

	r.ResponseData = bytes.NewBuffer(nil)
	_, r.Err = io.Copy(r.ResponseData, r.HttpResponse.Body)
	return
}
