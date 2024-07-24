package jeepay

import (
	_context "context"
	"errors"
	_nethttp "net/http"
)

type TransferApiService Service

func (p *TransferApiService) TransferOrder(ctx _context.Context, query TransferExecRequest) (Response[TransferExecResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/transferOrder", PayModel: model}
	execute, response, err := postExecute[Response[TransferExecResponse]](request)
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

func (p *TransferApiService) QueryTransferOrder(ctx _context.Context, query QueryTransferRequest) (Response[QueryTransferResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/transfer/query", PayModel: model}
	execute, response, err := postExecute[Response[QueryTransferResponse]](request)
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

func (p *TransferApiService) NotifyTransferOrder(ctx _context.Context, query RefundRequest, path string) (NotifyTransferResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: path, PayModel: model}
	return postExecute[NotifyTransferResponse](request)
}
