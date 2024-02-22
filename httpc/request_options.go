package httpc

import (
	"io"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type RequestFunc func(request *resty.Request)

func SetResult(res interface{}) RequestFunc {
	return func(r *resty.Request) {
		r.SetResult(res)
	}
}

func SetBody(body interface{}) RequestFunc {
	return func(r *resty.Request) {
		r.SetBody(body)
	}
}

func SetHeader(header, value string) RequestFunc {
	return func(r *resty.Request) {
		r.SetHeader(header, value)
	}
}

func SetHeaders(headers map[string]string) RequestFunc {
	return func(r *resty.Request) {
		r.SetHeaders(headers)
	}
}

func SetFileReader(param, fileName string, reader io.Reader) RequestFunc {
	return func(r *resty.Request) {
		r.SetFileReader(param, fileName, reader)
	}
}

func SetHeadersFromHTTPHeader(header http.Header) RequestFunc {
	return func(r *resty.Request) {
		r.Header = header
	}
}

func SetQueryParamsFromValues(params url.Values) RequestFunc {
	return func(r *resty.Request) {
		r.SetQueryParamsFromValues(params)
	}
}

func SetQueryParams(params map[string]string) RequestFunc {
	return func(r *resty.Request) {
		r.SetQueryParams(params)
	}
}

func SetQueryParam(param, value string) RequestFunc {
	return func(r *resty.Request) {
		r.SetQueryParam(param, value)
	}
}

func ForceContentType(contentType string) RequestFunc {
	return func(r *resty.Request) {
		r.ForceContentType(contentType)
	}
}

func SetFormData(data map[string]string) RequestFunc {
	return func(r *resty.Request) {
		r.SetFormData(data)
	}
}

func SetPathParam(param, value string) RequestFunc {
	return func(r *resty.Request) {
		r.SetPathParam(param, value)
	}
}

func SetPathParams(params map[string]string) RequestFunc {
	return func(r *resty.Request) {
		r.SetPathParams(params)
	}
}
