package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
)

// 授权发起Request
type InitiateAuthRequest struct {
	DataDoi       string                `json:"data_doi" binding:"required"`      //数据DOI
	Authorization idl.DataAuthorization `json:"authorization" binding:"required"` // 授权信息
	IDL.SignatureData
}

// 授权发起Response
type InitiateAuthResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
