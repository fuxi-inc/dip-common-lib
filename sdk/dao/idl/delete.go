package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

type DeleteDataContentRequest struct {
	Doi string `json:"doi,omitempty" binding:"required"`

	IDL.SignatureData
}

func (s *DeleteDataContentRequest) ToString() string {
	return converter.ToString(s)
}

// type UpdateDataContentResponseData struct{}
