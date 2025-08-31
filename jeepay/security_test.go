package jeepay

import "testing"

// {PayOrderId:P1962057372443525121 MchNo:M1747033666 AppId:68219e426e70ed6a5e4229b6 MchOrderNo:2025831153825967041216 IfCode:wxpay WayCode:WX_APP Amount:1 Currency:cny State:2 ClientIp:172.17.0.15 Subject:师者 vip 购买(测试) Body:测试使用 ChannelOrderNo:4200002784202508315032772904 ErrCode: ErrMsg: ExtParam:{\"uid\":\"261321239044296704\",\"vip_pkg_id\":7,\"duration\":30} CreatedAt:1756625905510 SuccessTime:1756625913000 ReqTime:1756625913248 Sign:76553A94D38639AF0B60D8F7EBB62CE7}

func TestSecurity(t *testing.T) {
	//amount=10000&clientIp=192.168.0.111&mchOrderNo=P0123456789101&notifyUrl=https://www.baidu.com&platId=1000&reqTime=20190723141000&returnUrl=https://www.baidu.com&version=1.0&key=EWEFD123RGSRETYDFNGFGFGSHDFGH
	//{P1962057372443525121 M1747033666 68219e426e70ed6a5e4229b6 2025831153825967041216 wxpay WX_APP 1 cny 2 172.17.0.15 师者 vip 购买(测试) 测试使用 4200002784202508315032772904   {"uid":"261321239044296704","vip_pkg_id":7,"duration":30} 1756625905510 1756625913000 1756625913248}
	payNotify := OrderNotifyCheckSignRequest{
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
	ok, err := checkSign(payNotify, key, "MD5", "76553A94D38639AF0B60D8F7EBB62CE7")

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
