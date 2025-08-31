package jeepay

import "testing"

type PayNotifyReq struct {
	PayOrderId     string `form:"payOrderId,optional,omitempty"`     // 支付订单号
	MchNo          string `form:"mchNo,optional,omitempty"`          // 商户号
	AppId          string `form:"appId,optional,omitempty"`          // 应用ID
	MchOrderNo     string `form:"mchOrderNo,optional,omitempty"`     // 商户订单号
	IfCode         string `form:"ifCode,optional,omitempty"`         // 支付接口编码
	WayCode        string `form:"wayCode,optional,omitempty"`        // 支付方式
	Amount         int    `form:"amount,optional,omitempty"`         // 支付金额, 单位分
	Currency       string `form:"currency,optional,omitempty"`       // 货币代码
	State          int    `form:"state,optional,omitempty"`          // 支付订单状态
	ClientIp       string `form:"clientIp,optional,omitempty"`       // 客户端IPV4地址, 可选
	Subject        string `form:"subject,optional,omitempty"`        // 商品标题
	Body           string `form:"body,optional,omitempty"`           // 商品描述
	ChannelOrderNo string `form:"channelOrderNo,optional,omitempty"` // 对应渠道的订单号, 可选
	ErrCode        string `form:"errCode,optional,omitempty"`        // 渠道下单返回错误码, 可选
	ErrMsg         string `form:"errMsg,optional,omitempty"`         // 渠道下单返回错误描述, 可选
	ExtParam       string `form:"extParam,optional,omitempty"`       // 商户扩展参数, 可选
	CreatedAt      int64  `form:"createdAt,optional,omitempty"`      // 订单创建时间, 13位时间戳
	SuccessTime    int64  `form:"successTime,optional,omitempty"`    // 订单支付成功时间, 13位时间戳, 可选
	ReqTime        int64  `form:"reqTime,optional,omitempty"`        // 通知请求时间, 13位时间戳
	Sign           string `form:"sign,optional,omitempty"`           // 签名值
}

// {PayOrderId:P1962057372443525121 MchNo:M1747033666 AppId:68219e426e70ed6a5e4229b6 MchOrderNo:2025831153825967041216 IfCode:wxpay WayCode:WX_APP Amount:1 Currency:cny State:2 ClientIp:172.17.0.15 Subject:师者 vip 购买(测试) Body:测试使用 ChannelOrderNo:4200002784202508315032772904 ErrCode: ErrMsg: ExtParam:{\"uid\":\"261321239044296704\",\"vip_pkg_id\":7,\"duration\":30} CreatedAt:1756625905510 SuccessTime:1756625913000 ReqTime:1756625913248 Sign:76553A94D38639AF0B60D8F7EBB62CE7}

func TestSecurity(t *testing.T) {
	//amount=10000&clientIp=192.168.0.111&mchOrderNo=P0123456789101&notifyUrl=https://www.baidu.com&platId=1000&reqTime=20190723141000&returnUrl=https://www.baidu.com&version=1.0&key=EWEFD123RGSRETYDFNGFGFGSHDFGH

	payNotify := PayNotifyReq{
		PayOrderId:     "P1962057372443525121",
		MchNo:          "M1747033666",
		AppId:          "68219e426e70ed6a5e4229b6",
		MchOrderNo:     "2025831153825967041216",
		IfCode:         "wxpay",
		WayCode:        "WX_APP",
		Amount:         1,
		Currency:       "cny",
		State:          2,
		ClientIp:       "172.17.0.15",
		Subject:        "师者 vip 购买(测试)",
		Body:           "测试使用",
		ChannelOrderNo: "4200002784202508315032772904",
		ErrCode:        "",
		ErrMsg:         "",
		ExtParam:       "{\"uid\":\"261321239044296704\",\"vip_pkg_id\":7,\"duration\":30}",
		CreatedAt:      1756625905510,
		SuccessTime:    1756625913000,
		ReqTime:        1756625913248,
		//Sign:           "76553A94D38639AF0B60D8F7EBB62CE7",
	}

	key := "FOdgiJfwY8veVkPBCqPSTsCEudrOohrdQz8eCflJAqg1WhiqHeZDWpsgsuCjqMoo1infXjlxDLKnfNPCU1GTl0KOzCLzFwQ4AcqRVQJfOLruoSGHb1xPBhdzjpTXXQT5"
	ok, err := checkSign(payNotify, key, "MD5", payNotify.Sign)

	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Error("签名验证失败")
		return
	}
	t.Log("签名验证成功")
}

func TestMd5(t *testing.T) {
	s := md5encrypt([]byte("amount=10000&clientIp=192.168.0.111&mchOrderNo=P0123456789101&notifyUrl=https://www.baidu.com&platId=1000&reqTime=20190723141000&returnUrl=https://www.baidu.com&version=1.0&key=EWEFD123RGSRETYDFNGFGFGSHDFGH"))
	t.Log(s)
}
