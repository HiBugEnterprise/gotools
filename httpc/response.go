package httpc

import (
	"context"
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
	body.Msg = "success"
	body.Data = resp
	httpx.OkJsonCtx(ctx, w, body)
}

func RespError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	var (
		code     = http.StatusInternalServerError
		res      = Response{Code: code, Msg: "服务器开小差啦，稍后再来试一试"}
		metadata any
		appType  string
	)
	switch err.(type) {
	case *errorx.Error:
		customErr := errorx.From(err)
		res.Code = customErr.Code
		res.Msg = customErr.Msg
		code = customErr.Code
		appType = customErr.BizType
		metadata = customErr.Metadata
	}

	logc.Errorw(ctx, res.Msg,
		logc.Field("err", err),
		logc.Field("code", code),
		logc.Field("type", appType),
		logc.Field("metadata", metadata),
		logc.Field("method", r.Method),
		logc.Field("path", r.URL.Path),
	)

	httpx.OkJsonCtx(ctx, w, res)
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	logc.Errorw(r.Context(), "Auth failed",
		logc.Field("err", err),
		logc.Field("method", r.Method),
		logc.Field("path", r.URL.Path),
	)

	httpx.WriteJson(w, http.StatusUnauthorized, &Response{http.StatusUnauthorized, "Auth failed", nil})
}
