package errorx

type ResCode int

const (
	CodeSuccess ResCode = 200

	CodeInvalidParams      ResCode = 400
	CodeValidateParamsErr  ResCode = 400001
	CodeConfirmPasswordErr ResCode = 400003

	CodeNeedLogin ResCode = 401

	CodeInvalidToken            ResCode = 403
	CodeWrongPassword           ResCode = 403001
	CodeWrongUserNameOrPassword ResCode = 403002
	CodeNeedAcceptAgreement     ResCode = 403005
	CodeWrongOldPassword        ResCode = 403006
	CodeWrongEmailCode          ResCode = 403007
	CodeNotAdminRole            ResCode = 403008

	CodeUserNotFound ResCode = 404001

	CodeDistributedLock ResCode = 409001
	CodeUserExist       ResCode = 409002
	CodeEmailExist      ResCode = 409003
	CodePhoneExist      ResCode = 409004
	CodeTenantExist     ResCode = 409006
	CodeOnlyOneAdmin    ResCode = 409007

	CodeInternalErr            ResCode = 500
	CodeGenTokenErr            ResCode = 500001
	CodeJSONEmptyErr           ResCode = 500009
	CodeOnlyAdminCanDo         ResCode = 500015
	CodeUnSupportPlatform      ResCode = 500016
	CodeUnsupportedLanguage    ResCode = 500017
	CodeCreateScannerConfigErr ResCode = 500021
)
