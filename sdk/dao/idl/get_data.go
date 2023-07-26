package idl

import (
	"bytes"
	"encoding/json"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"

	"github.com/fuxi-inc/dip-common-lib/IDL"
)

type GetDataRequest struct {
	DuDoi   string `json:"du_doi" binding:"required"`   //数据所有者身份标识
	DataDoi string `json:"data_doi" binding:"required"` //数据对象标识
	//PermissionDoi string `json:"permission_doi" binding:"required"` //权限
	IDL.SignatureData
}

func (r *GetDataRequest) ToString() string {
	return converter.ToString(r)
}

// type GetDataResponseData struct {
// 	DataContent string `json:"data_content" binding:"required"`
// }

type GetDataResponse struct {
	Code    int64                `json:"code"`
	Data    *GetDataResponseData `json:"data"`
	Message string               `json:"message"`
}

type GetDataResponseData struct {
	IsAccessible    bool            `json:"is_accessible"`
	DataContent     string          `json:"data_content"`
	EncryptedKey    string          `json:"encrypted_key"`
	TransactionInfo interface{}     `json:"transaction_info"`
	Digest          *idl.DataDigest `json:"digest"`
}

func (r *GetDataResponse) ToString() string {
	return converter.ToString(r)
}

func (r *GetDataResponse) Unmarshal(data []byte) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(&r)
}
