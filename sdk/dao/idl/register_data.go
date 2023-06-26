package idl

import "github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"

type RegisterDataRequest struct {
	Doi                      string                           `json:"doi,omitempty" binding:"required"`
	DwDoi                    string                           `json:"dw_doi" binding:"required"`    //数据所有者身份标识
	Signature                string                           `json:"signature" binding:"required"` //DW的私钥签名
	PublicKey                string                           `json:"public_key" binding:"required"`
	Content                  []byte                           `json:"content" binding:"required"`                //数据内容
	FilePath                 string                           `json:"file_path" binding:"required"`              //保存的文件路径
	Digest                   *idl.ApiDigest                   `json:"digest,omitempty" binding:"required"`       // 数据内容摘要
	Confirmation             string                           `json:"confirmation,omitempty" binding:"required"` // 确权信息。DW私钥对数据摘要的签名
	SecretKey                string                           `json:"secret_key"`                                //文件内容的加密秘钥，当为空时，代表文件内容没有加密
	ClassificationAndGrading *idl.ApiClassificationAndGrading `json:"classification_grading,omitempty"`          // 数据分类分级信息
}

type RegisterDataResponseData struct{}
