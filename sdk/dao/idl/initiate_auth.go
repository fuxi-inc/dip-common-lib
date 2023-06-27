package idl

import "github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"

// 授权发起Request
type InitiateAuthRequest struct {
	DataDoi       string                `json:"data_doi" binding:"required"`      //数据DOI
	Authorization idl.DataAuthorization `json:"authorization" binding:"required"` // 授权信息
	InitiatorDOI  string                `json:"initiator_doi" binding:"required"` //发起者（DW或DU）的DOI
	Sign          string                `json:"sign" binding:"required"`          // 发起者私钥对其自身DOI的签名
}

// 授权发起Response
type InitiateAuthResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
