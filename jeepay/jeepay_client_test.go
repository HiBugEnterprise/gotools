package jeepay_go_sdk

import (
	"context"
	"fmt"
	"testing"
)

const (
	host   = "xxxxxxxxxxxxxxxxxx"
	schema = "http"
)

// TestCreateClient 测试创建客户端
func TestCreateClient(t *testing.T) {
	newConfiguration := NewConfiguration()
	newConfiguration.AppId = "test"
	newConfiguration.AppSecret = "test"
	newConfiguration.Host = host
	newConfiguration.Scheme = schema
	client := NewApiClient(newConfiguration)
	_ = client.PayApi.client
}

func TestPayApi(t *testing.T) {
	newConfiguration := NewConfiguration()
	newConfiguration.AppId = "xxxxxxxxxxx"
	newConfiguration.AppSecret = "xxxxxxxxxxxxxx"
	newConfiguration.Host = host
	newConfiguration.Scheme = schema
	client := NewApiClient(newConfiguration)
	amount := 4231
	mchno := "xxxxxxxxxxxx"
	mchorderno := "asdasd"
	waycode := WxLite
	currency := "cny"
	subject := "测试"
	body := "测试"
	weixinAppletOpenId := "xxxxxxxxxxxxxxx"

	request := PayCreateRequest{
		Amount:       &amount,
		MchNo:        &mchno,
		MchOrderNo:   &mchorderno,
		WayCode:      &waycode,
		Currency:     &currency,
		Subject:      &subject,
		Body:         &body,
		ChannelExtra: Pointer(fmt.Sprintf(`{"%s": "%s"}`, "openid", weixinAppletOpenId)),
	}

	execute, response, err := client.PayApi.CreateOrder(context.Background(), request)

	if err != nil {
		t.Error(err)
	}
	t.Log(execute)
	t.Log(response)
}

func TestQueryOrder(t *testing.T) {
	newConfiguration := NewConfiguration()
	newConfiguration.AppId = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	newConfiguration.AppSecret = "xxxxxxxxxxxxxxxxxxxxxxxxx"
	newConfiguration.Host = host
	newConfiguration.Scheme = schema
	client := NewApiClient(newConfiguration)

	request := PayQueryRequest{
		MchNo:      Pointer("xxxxxxxxxxx"),
		MchOrderNo: Pointer("xxxxxxxxxxxxxxxxxx"),
	}

	execute, _, err := client.PayApi.QueryOrder(context.Background(), request)

	if err != nil {
		t.Log(err)
		return
	}
	if execute.Code != 0 {
		fmt.Println(execute.Code, execute.Msg)
	}
}
