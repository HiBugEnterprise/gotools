package jeepay_go_sdk

import (
	"context"
	"fmt"
	"testing"
)

const (
	host   = "XXXXXXXXXXXXXXXXXXXXXXXX"
	schema = "http"
)

func client() *APIClient {
	newConfiguration := NewConfiguration()
	newConfiguration.AppId = "XXXXXXXXXXXXXXXXXXXXXX"
	newConfiguration.AppSecret = "XXXXXXXXXXXXXXXXX"
	newConfiguration.Host = host
	newConfiguration.Scheme = schema
	cli := NewApiClient(newConfiguration)
	return cli
}

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
	mchorderno := "xxxxxxxxx"
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

func TestCloseOrder(t *testing.T) {
	cli := client()

	request := OrderCloseRequest{
		MchNo:      "XXXXXXXXXXXX",
		MchOrderNo: "xxxxxxxxx",
	}
	execute, _, err := cli.PayApi.CloseOrder(context.Background(), request)
	if err != nil {
		t.Log(err)
		return
	}
	if execute.Code != 0 {
		t.Log(execute.Code, execute.Msg)
	}
	t.Log(execute)
}
func TestQueryOrder(t *testing.T) {
	cli := client()

	request := PayQueryRequest{
		MchNo:      Pointer("XXXXXXXXXXXXXXXX"),
		MchOrderNo: Pointer("xxxxxxxxxxxxxxxx"),
	}

	execute, _, err := cli.PayApi.QueryOrder(context.Background(), request)

	fmt.Printf("%+v\n", execute.Data)
	fmt.Println(*execute.Data.Amount, *execute.Data.CreatedAt, *execute.Data.State)
	if err != nil {
		t.Log(err)
		return
	}
	if execute.Code != 0 {
		fmt.Println(execute.Code, execute.Msg)
	}
}
