package errorx

// codeMsgMap 预设错误, 业务错误请在业务代码中定义
var codeMsgMap map[ResCode]string

func init() {
	codeMsgMap = map[ResCode]string{
		CodeSuccess:       "操作成功",
		CodeInvalidParams: "请求参数错误",
		CodeUnauthorized:  "请先登陆",
		CodeInvalidToken:  "无效的Token",
		CodeInternalErr:   "服务繁忙，请稍后再试",
	}
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeInternalErr]
	}
	return msg
}
