package jeepay_go_sdk

import (
	_context "context"
	"errors"
	_nethttp "net/http"
)

type RefundApiService Service

func (p *RefundApiService) RefundOrder(ctx _context.Context, query RefundRequest) (Response[RefundResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/refund/refundOrder", PayModel: model}
	execute, response, err := postExecute[Response[RefundResponse]](request)
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

func (p *RefundApiService) QueryRefundOrder(ctx _context.Context, query RefundQueryRequest) (Response[RefundQueryResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/refund/query", PayModel: model}

	execute, response, err := postExecute[Response[RefundQueryResponse]](request)
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

func (p *RefundApiService) RefundNotifyOrder(ctx _context.Context, query RefundRequest) (RefundQueryResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/refund/query", PayModel: model}
	return postExecute[RefundQueryResponse](request)
}
