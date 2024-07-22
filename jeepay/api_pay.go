package jeepay_go_sdk

import (
	_context "context"
	"errors"
	_nethttp "net/http"
)

type PayApiService Service

func (p *PayApiService) CreateOrder(ctx _context.Context, query PayCreateRequest) (Response[BaseResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/pay/unifiedOrder", PayModel: model}
	// 做安全检测
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

func (p *PayApiService) QueryOrder(ctx _context.Context, query PayQueryRequest) (Response[PayQueryItem], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/pay/query", PayModel: model}
	execute, response, err := postExecute[Response[PayQueryItem]](request)
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

func (p *PayApiService) NotificationOrder(ctx _context.Context, query OrderNotifyRequest, path string) (OrderNotifyResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: path, PayModel: model}
	return postExecute[OrderNotifyResponse](request)
}

func (p *PayApiService) CloseOrder(ctx _context.Context, query OrderCloseRequest) (Response[BaseResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/pay/close", PayModel: model}
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

func (p *PayApiService) ChannelOrder(ctx _context.Context, query OrderCloseRequest) (OrderChannelResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/channelUserId/jump", PayModel: model}
	return postExecute[OrderChannelResponse](request)
}
