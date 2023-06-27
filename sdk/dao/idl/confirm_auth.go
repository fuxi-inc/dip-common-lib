package idl

import dis_idl "github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"

// 授权确认Request
type ConfirmAuthRequest struct {
	DataDoi       string                    `json:"data_doi" binding:"required"`      //数据DOI
	Authorization dis_idl.DataAuthorization `json:"authorization" binding:"required"` // 授权信息
	ConfirmerDOI  string                    `json:"confirmer_doi" binding:"required"` //确认者（DW或DU）的DOI
	Sign          string                    `json:"sign" binding:"required"`          // 使用者私钥对其自身DOI的签名
}

// 授权确认Response
type ConfirmAuthResponse struct {
	Errno  int64  `json:"errno"`
	Errmsg string `json:"errmsg"`
}
