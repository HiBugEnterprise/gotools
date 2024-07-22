# Jeepay-sdk-go

#### 加密规则: [协议规则](https://docs.jeequan.com/docs/jeepay/api_rule)

# 使用说明

其中的AppId和AppSecret为jeepay应用列表对应的AppId和AppSecret

MchNo是和AppId对应的商户号
```go
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

```