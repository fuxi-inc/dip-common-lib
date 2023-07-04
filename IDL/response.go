package IDL

import (
	"bytes"
	"encoding/json"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

type CommonResponse struct {
	Code    RespCodeType `json:"code"`
	Data    interface{}  `json:"data"`
	Message string       `json:"message"`
}

func NewCommonResponse() *CommonResponse {
	return &CommonResponse{}
}
func (r *CommonResponse) SetCode(code int64) *CommonResponse {
	r.Code = RespCodeType(code)
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

func (r *CommonResponse) ToString() string {
	return converter.ToString(r)
}

func (r *CommonResponse) Unmarshal(data []byte) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(&r)
}
