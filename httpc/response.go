package httpc

import (
	"context"
	"errors"
	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func RespSuccess(ctx context.Context, w http.ResponseWriter, resp interface{}) {
	var body Response
	body.Code = http.StatusOK
	body.Msg = errorx.CodeSuccess.Msg()
	body.Data = resp
	httpx.OkJsonCtx(ctx, w, body)
}

func RespError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		code      = http.StatusInternalServerError
		res       = Response{Code: code, Msg: errorx.CodeInternalErr.Msg()}
		metadata  any
		bizType   string
		errDetail string
	)

	var customErr *errorx.Error
	switch {
	case errors.As(err, &customErr):
		res.Code = customErr.Code
		errDetail = customErr.Msg
		if customErr.IsShow {
			res.Msg = customErr.Msg
		}
		code = customErr.Code
		bizType = customErr.BizType
		metadata = customErr.Metadata
	}

	logc.Errorw(r.Context(), errDetail,
		logc.Field("err", err),
		logc.Field("code", code),
		logc.Field("type", bizType),
		logc.Field("metadata", metadata),
		logc.Field("method", r.Method),
		logc.Field("path", r.URL.Path),
	)

	httpx.OkJsonCtx(r.Context(), w, res)
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	logc.Errorw(r.Context(), "Auth failed",
		logc.Field("err", err),
		logc.Field("method", r.Method),
		logc.Field("path", r.URL.Path),
	)

	httpx.WriteJson(w, http.StatusUnauthorized, &Response{http.StatusUnauthorized, "Auth failed", nil})
}
