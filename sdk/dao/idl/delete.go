package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

type DeleteDataRequest struct {
	Doi string `json:"doi,omitempty" binding:"required"`

	IDL.SignatureData
}

func (s *DeleteDataRequest) ToString() string {
	return converter.ToString(s)
}
