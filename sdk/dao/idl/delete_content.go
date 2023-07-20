package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

type DeleteDataContentRequest struct {
	Doi   string `json:"doi,omitempty" binding:"required"`
	DwDoi string `json:"dw_doi" binding:"required"` //数据所有者身份标识

	IDL.SignatureData
}

func (s *DeleteDataContentRequest) ToString() string {
	return converter.ToString(s)
}

// type UpdateDataContentResponseData struct{}
