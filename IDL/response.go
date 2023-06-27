package IDL

type CommonResponse struct {
    Code    int64       `json:"code"`
    Data    interface{} `json:"data"`
    Message string      `json:"message"`
}

func (r *CommonResponse) SetCode(code int64) *CommonResponse {
    r.Code = code
    return r
}

func (r *CommonResponse) SetData(data interface{}) *CommonResponse {
    r.Data = data
    return r
}

func (r *CommonResponse) SetMessage(message string) *CommonResponse {
    r.Message = message
    return r
}
