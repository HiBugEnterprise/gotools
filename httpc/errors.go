package httpc

import (
	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func NewErrorFromRestyResponse(res *resty.Response) *errorx.Error {
	switch res.StatusCode() {
	case http.StatusConflict:

	}
	return errorx.New(
		http.StatusText(res.StatusCode()),
		res.StatusCode(),
		res.String())

}
