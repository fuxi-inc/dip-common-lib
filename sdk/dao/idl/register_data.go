package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

type RegisterDataRequest struct {
	Doi                      string                        `json:"doi,omitempty" binding:"required"`
	DwDoi                    string                        `json:"dw_doi" binding:"required"` //数据所有者身份标识
	PubKey                   string                        `json:"pub_key" binding:"required"`
	Content                  string                        `json:"content"`                                   //数据内容
	FilePath                 string                        `json:"file_path" binding:"required"`              //保存的文件路径
	Digest                   *idl.DataDigest               `json:"digest,omitempty" binding:"required"`       // 数据内容摘要
	Confirmation             string                        `json:"confirmation,omitempty" binding:"required"` // 确权信息。DW私钥对数据摘要的签名
	SecretKey                string                        `json:"secret_key"`                                //文件内容的加密秘钥，当为空时，代表文件内容没有加密
	ClassificationAndGrading *idl.ClassificationAndGrading `json:"classification_grading,omitempty"`          // 数据分类分级信息
	IDL.SignatureData
}

func (s *RegisterDataRequest) ToString() string {
	return converter.ToString(s)
}

type RegisterDataResponseData struct{}
