package errorx

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "操作成功",
	CodeInvalidParams: "请求参数错误",
	CodeNeedLogin:     "请先登陆",
	CodeInvalidToken:  "无效的token",
	CodeInternalErr:   "服务器开小差啦，稍后再来试一试",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeInternalErr]
	}
	return msg
}
