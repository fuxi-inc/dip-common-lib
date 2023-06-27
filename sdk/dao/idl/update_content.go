package idl

import "github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"

type UpdateDataContentRequest struct {
	Doi          string          `json:"doi,omitempty" binding:"required"`
	DwDoi        string          `json:"dw_doi" binding:"required"`                 //数据所有者身份标识
	Content      []byte          `json:"content" binding:"required"`                //数据内容
	Digest       *idl.DataDigest `json:"digest,omitempty" binding:"required"`       // 数据内容摘要
	Confirmation string          `json:"confirmation,omitempty" binding:"required"` // 确权信息。DW私钥对数据摘要的签名
	SecretKey    string          `json:"secret_key"`                                //文件内容的加密秘钥，当为空时，代表文件内容没有加密
	idl.SignatureData
}

type UpdateDataContentResponseData struct{}
