package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

type UpdateDataContentRequest struct {
	Doi          string          `json:"doi,omitempty" binding:"required"`
	DwDoi        string          `json:"dw_doi" binding:"required"`                 //数据所有者身份标识
	Content      string          `json:"content" binding:"required"`                //数据内容
	Digest       *idl.DataDigest `json:"digest,omitempty" binding:"required"`       // 数据内容摘要
	Confirmation string          `json:"confirmation,omitempty" binding:"required"` // 确权信息。DW私钥对数据摘要的签名
	SecretKey    string          `json:"secret_key"`                                //文件内容的加密秘钥，当为空时，代表文件内容没有加密
	IDL.SignatureData
}

func (s *UpdateDataContentRequest) ToString() string {
	return converter.ToString(s)
}

type UpdateDataContentResponseData struct{}
