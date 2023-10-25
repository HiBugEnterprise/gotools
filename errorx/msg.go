package errorx

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "操作成功",

	// 400 系列
	CodeValidateParamsErr:  "参数校验错误",
	CodeInvalidParams:      "请求参数错误",
	CodeConfirmPasswordErr: "两次输入的密码不一致",
	// 401 系列
	CodeNeedLogin: "请先登陆",

	// 403 系列
	CodeInvalidToken:            "无效的token",
	CodeWrongPassword:           "密码错误",
	CodeWrongUserNameOrPassword: "用户名或密码错误",
	CodeNeedAcceptAgreement:     "未勾选用户协议",
	CodeWrongOldPassword:        "旧密码不正确",
	CodeWrongEmailCode:          "邮箱验证码不正确",
	CodeNotAdminRole:            "操作失败，只有管理员可操作",

	// 404 系列
	CodeUserNotFound: "该用户不存在",

	// 409 系列
	CodeDistributedLock: "请求频繁，请稍后再试",
	CodeUserExist:       "该用户已存在",
	CodeEmailExist:      "该邮箱已注册",
	CodePhoneExist:      "该手机号已注册",
	CodeTenantExist:     "该租户已存在",
	CodeOnlyOneAdmin:    "只能有一个管理员",

	// 500 系列
	CodeInternalErr:            "服务器开小差啦，稍后再来试一试",
	CodeGenTokenErr:            "生成Token异常",
	CodeJSONEmptyErr:           "解析JSON异常,JSON empty",
	CodeOnlyAdminCanDo:         "非管理员不能操作",
	CodeUnSupportPlatform:      "暂不支持该平台",
	CodeUnsupportedLanguage:    "暂不支持该语言",
	CodeCreateScannerConfigErr: "创建扫描配置失败",
}

func (code ResCode) Msg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeInternalErr]
	}
	return msg
}
