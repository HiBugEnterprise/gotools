package errorx

type ResCode int

const (
	CodeSuccess       ResCode = 200
	CodeInvalidParams ResCode = 400
	CodeNeedLogin     ResCode = 401
	CodeInvalidToken  ResCode = 403
	CodeInternalErr   ResCode = 500
)
