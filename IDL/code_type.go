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

	var codeType int64
	switch data.(type) {
	case string:
		codeType = converter.StringToInt64(data.(string))
	case int, int16, int32, int64:
		codeType = int64(data.(int))
	case float32, float64:
		codeType = int64(data.(float64))
	}
	*s = RespCodeType(codeType)
	return nil
}
