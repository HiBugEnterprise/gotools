package jeepay

import (
	_context "context"
	"errors"
	_nethttp "net/http"
)

type SubAccountApiService Service

func (p *SubAccountApiService) DivisionBindReceiver(ctx _context.Context, query DivisionBindRequest) (DivisionBindResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/division/receiver/bind", PayModel: model}
	return postExecute[DivisionBindResponse](request)
}

func (p *SubAccountApiService) DivisionExec(ctx _context.Context, query DivisionExecRequest) (Response[BaseResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/division/exec", PayModel: model}
	execute, response, err := postExecute[Response[BaseResponse]](request)
	if err != nil {
		return execute, response, err
	}

	// 对data内数据签名,如data为空则不返回，sign不为空说明data不为空
	equal, err := checkSign(execute.Data, p.Configuration.AppSecret, "MD5", execute.Sign)
	if err != nil {
		return execute, response, err
	}
	if !equal {
		err = errors.New("返回结果验证签名失败")
		return execute, response, err
	}

	return execute, response, nil
}
