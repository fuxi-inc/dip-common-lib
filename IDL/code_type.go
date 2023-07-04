package IDL

import (
	"encoding/json"
	"fmt"
	"github.com/fuxi-inc/dip-common-lib/constants"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

func (s RespCodeType) IsSuccess() bool {
	return s.ToInt64() == constants.Success
}

type RespCodeType int64

func (s RespCodeType) ToInt() int {
	return int(s.ToInt64())
}
func (s RespCodeType) ToInt8() int8 {
	return int8(s.ToInt64())
}
func (s RespCodeType) ToInt16() int16 {
	return int16(s.ToInt64())
}
func (s RespCodeType) ToInt32() int32 {
	return int32(s.ToInt64())
}
func (s RespCodeType) ToInt64() int64 {
	return int64(s)
}
func (s RespCodeType) ToString() string {
	return fmt.Sprintf("%d", s.ToInt())
}

func (s *RespCodeType) UnmarshalJSON(b []byte) error {
	var data interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	*s = RespCodeType(converter.InterfaceToInt64(data))
	return nil
}
