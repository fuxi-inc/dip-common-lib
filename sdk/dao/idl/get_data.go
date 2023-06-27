package idl

import (
	"bytes"
	"encoding/json"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"

	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

type GetDataRequest struct {
	DuDoi        string `json:"du_doi" binding:"required"`        //数据所有者身份标识
	DataDoi      string `json:"data_doi" binding:"required"`      //数据对象标识
	PermisionDoi string `json:"permision_doi" binding:"required"` //权限
	idl.SignatureData
}

// type GetDataResponseData struct {
// 	DataContent string `json:"data_content" binding:"required"`
// }

type GetDataResponse struct {
	Errno       int64  `json:"errno"`
	Errmsg      string ` son:"errmsg"`
	DataContent string `json:"data_content" binding:"required"`
}

func (r GetDataRequest) ToString() string {
	return converter.ToString(r)
}

func (r *GetDataResponse) Unmarshal(data []byte) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(&r)
}
