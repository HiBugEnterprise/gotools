package jeepay

type PayCreateRequest struct {
	MchNo        *string  `json:"mchNo"`        // 商户号
	AppId        *string  `json:"appId"`        // 应用ID
	MchOrderNo   *string  `json:"mchOrderNo"`   // 商户订单号
	WayCode      *WayCode `json:"wayCode"`      // 支付方式
	Amount       *int     `json:"amount"`       // 支付金额
	Currency     *string  `json:"currency"`     // 货币代码
	ClientIp     *string  `json:"clientIp"`     // 客户端IPV4地址
	Subject      *string  `json:"subject"`      // 商品标题
	Body         *string  `json:"body"`         // 商品描述
	NotifyUrl    *string  `json:"notifyUrl"`    // 异步通知地址
	ReturnUrl    *string  `json:"returnUrl"`    // 跳转通知地址
	ExpiredTime  *int     `json:"expiredTime"`  // 失效时间
	ChannelExtra *string  `json:"channelExtra"` // 渠道参数
	DivisionMode *int     `json:"divisionMode"` // 分账模式
	ExtParam     *string  `json:"extParam"`     // 扩展参数
	ReqTime      *int64   `json:"reqTime"`      // 请求时间
	Version      *string  `json:"version"`      // 接口版本
	Sign         *string  `json:"sign"`         // 签名值
	SignType     *string  `json:"signType"`     // 签名类型
}
type PayQueryItem struct {
	Amount         *int    `json:"amount"`
	AppId          string  `json:"appId"`
	Body           string  `json:"body"`
	ChannelOrderNo string  `json:"channelOrderNo"`
	ClientIp       string  `json:"clientIp"`
	CreatedAt      *int64  `json:"createdAt"`
	Currency       string  `json:"currency"`
	ExtParam       string  `json:"extParam"`
	IfCode         string  `json:"ifCode"`
	MchNo          string  `json:"mchNo"`
	MchOrderNo     string  `json:"mchOrderNo"`
	PayOrderId     string  `json:"payOrderId"`
	State          *int    `json:"state"`
	Subject        string  `json:"subject"`
	SuccessTime    int64   `json:"successTime"`
	WayCode        string  `json:"wayCode"`
	ErrCode        *string `json:"errCode"`
	ErrMsg         *string `json:"errMsg"`
}

type BaseResponse struct {
	ErrCode     string `json:"errCode"`
	ErrMsg      string `json:"errMsg"`
	MchOrderNo  string `json:"mchOrderNo"`
	OrderState  *int   `json:"orderState"`
	PayOrderId  string `json:"payOrderId"`
	PayDataType string `json:"payDataType"`
	PayData     string `json:"payData"`
}

type PayChannelExtra struct {
	// 当 wayCode=ALI_JSAPI 时，channelExtra必须传buyerUserId，为支付宝用户ID，channelExtra示例数据如：
	BuyerUserId *string `json:"buyerUserId"`
	// 当 wayCode=AUTO_BAR 或 wayCode=ALI_BAR 或 wayCode=WX_BAR 或 wayCode=YSF_BAR 时，channelExtra必须传auth_code，为用户的付款码值，channelExtra示例数据如：
	AuthCode *string `json:"auth_code"`
	// 当 wayCode=WX_JSAPI 或 wayCode=WX_LITE 时，channelExtra必须传openid，为支付宝用户ID，channelExtra示例数据如：
	Openid *string `json:"openid"`
	// 当 wayCode=ALI_WAP 时，channelExtra可以传payDataType设置返回支付数据支付类型。此时payDataType可以为：form-返回自动跳转的支付表单,codeImgUrl-返回一个二维码图片URL,payUrl-返回支付链接，不传payDataType默认返回payUrl类型, channelExtra示例数据如：
	// 当 wayCode=ALI_PC 时，channelExtra可以传payDataType设置返回支付数据支付类型。此时payDataType可以为：form-返回自动跳转的支付表单,payUrl-返回支付链接，不传payDataType默认返回payUrl类型, channelExtra示例数据如：
	PayDataType *string `json:"payDataType"`
}

type PayQueryRequest struct {
	MchNo      *string `json:"mchNo"`      // 商户号
	AppId      *string `json:"appId"`      // 应用ID
	PayOrderId *string `json:"payOrderId"` // 支付中心生成的订单号，与mchOrderNo二者传一即可
	MchOrderNo *string `json:"mchOrderNo"` // 商户生成的订单号，与payOrderId二者传一即可
	ReqTime    *int64  `json:"reqTime"`    // 请求接口时间，13位时间戳
	Version    *string `json:"version"`    // 接口版本号，固定：1.0
	Sign       *string `json:"sign"`       // 签名值
	SignType   *string `json:"signType"`   // 签名类型，目前只支持MD5方式
}

func NewPayModel() *PayCreateRequest {
	this := PayCreateRequest{}
	return &this
}

type OrderCloseRequest struct {
	MchNo      *string `json:"mchNo"`      // 商户号
	AppId      *string `json:"appId"`      // 应用ID
	PayOrderId *string `json:"payOrderId"` // 支付中心生成的订单号，与mchOrderNo二者传一即可
	MchOrderNo *string `json:"mchOrderNo"` // 商户生成的订单号，与payOrderId二者传一即可
	ReqTime    *int64  `json:"reqTime"`    // 请求接口时间，13位时间戳
	Version    *string `json:"version"`    // 接口版本号，固定：1.0
	Sign       *string `json:"sign"`       // 签名值
	SignType   *string `json:"signType"`   // 签名类型，目前只支持MD5方式
}

type OrderNotifyRequest struct {
	Amount     int    `json:"amount"`
	Body       string `json:"body"`
	ClientIp   string `json:"clientIp"`
	CreatedAt  string `json:"createdAt"`
	Currency   string `json:"currency"`
	ExtParam   string `json:"extParam"`
	IfCode     string `json:"ifCode"`
	MchNo      string `json:"mchNo"`
	AppId      string `json:"appId"`
	MchOrderNo string `json:"mchOrderNo"`
	PayOrderId string `json:"payOrderId"`
	State      int    `json:"state"`
	Subject    string `json:"subject"`
	WayCode    string `json:"wayCode"`
	Sign       string `json:"sign"`
}

type OrderChannelRequest struct {
	// 商户号
	MchNo *string `json:"mchNo"`
	// 应用ID
	AppId *string `json:"appId"`
	// 支付接口，目前只支持传 AUTO
	IfCode *string `json:"ifCode"`
	// 跳转地址，获取到用户ID后，会携带用户ID参数跳转到该地址
	RedirectUrl *string `json:"redirectUrl"`
	// 请求接口时间，13位时间戳
	ReqTime *int64 `json:"reqTime"`
	// 接口版本号，固定：1.0
	Version *string `json:"version"`
	// 签名值
	Sign *string `json:"sign"`
	// 签名类型，目前只支持MD5方式
	SignType *string `json:"signType"`
}

type OrderChannelResponse struct {
	MchNo       string `json:"mchNo"`
	AppId       string `json:"appId"`
	IfCode      string `json:"ifCode"`
	RedirectUrl string `json:"redirectUrl"`
	Sign        string `json:"sign"`
	SignType    string `json:"signType"`
	ReqTime     string `json:"reqTime"`
	Version     string `json:"version"`
}

type OrderNotifyResponse struct {
	Amount     int    `json:"amount"`
	Body       string `json:"body"`
	ClientIp   string `json:"clientIp"`
	CreatedAt  string `json:"createdAt"`
	Currency   string `json:"currency"`
	ExtParam   string `json:"extParam"`
	IfCode     string `json:"ifCode"`
	MchNo      string `json:"mchNo"`
	AppId      string `json:"appId"`
	MchOrderNo string `json:"mchOrderNo"`
	PayOrderId string `json:"payOrderId"`
	State      int    `json:"state"`
	Subject    string `json:"subject"`
	WayCode    string `json:"wayCode"`
	Sign       string `json:"sign"`
}

type OrderNotifyCheckSignRequest struct {
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
}
