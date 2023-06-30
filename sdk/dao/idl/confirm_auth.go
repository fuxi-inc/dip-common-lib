package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

// 授权确认Request
type ConfirmAuthRequest struct {
	DataDoi       string                `json:"data_doi" binding:"required"`      //数据DOI
	Authorization idl.DataAuthorization `json:"authorization" binding:"required"` // 授权信息
	IDL.SignatureData
}

func (s *ConfirmAuthRequest) ToString() string {
	return converter.ToString(s)
}

// 授权确认Response
type ConfirmAuthResponseData struct {
}
