package idl

const (
    DemoConstA = "demoA"
    DemoConstB = "demoB"
)

type ApiDemoRequest struct {
    IsPrepay   bool        `json:"is_prepay,omitempty"`
    PrepayType int         `json:"prepay_type,omitempty"`
    TraceId    string      `json:"trace_id,omitempty"`
    PrePayment int         `json:"pre_payment,omitempty"`
    ApplePay   interface{} `json:"apple_pay,omitempty"`
    OutTradeId string      `json:"out_trade_id,omitempty"`
}

type ApiDemoResponseData struct {
    TraceId    string      `json:"trace_id,omitempty"`
    PrePayment int         `json:"pre_payment,omitempty"`
    ApplePay   interface{} `json:"apple_pay,omitempty"`
}

type ApiDemoResponse struct {
    Errno  int64                `json:"errno"`
    Errmsg string               ` son:"errmsg"`
    Data   *ApiDemoResponseData `json:"data,omitempty"`
}
