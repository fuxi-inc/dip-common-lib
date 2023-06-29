package idl

import (
	"bytes"
	"encoding/json"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

type GetDataRequest struct {
	DuDoi   string `json:"du_doi" binding:"required"`   //数据所有者身份标识
	DataDoi string `json:"data_doi" binding:"required"` //数据对象标识
	//PermissionDoi string `json:"permission_doi" binding:"required"` //权限
	IDL.SignatureData
}

// type GetDataResponseData struct {
// 	DataContent string `json:"data_content" binding:"required"`
// }

type GetDataResponse struct {
	Errno        int64  `json:"errno"`
	Errmsg       string `json:"errmsg"`
	DataContent  []byte `json:"data_content"`
	EncryptedKey string `json:"encrypted_key"`
}

func (r GetDataRequest) ToString() string {
	return converter.ToString(r)
}

func (r *GetDataResponse) Unmarshal(data []byte) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(&r)
}
