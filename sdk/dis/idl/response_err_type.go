package idl

import (
	"encoding/json"
	"fmt"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

const (
	Success DisRespErrno = 0 //所有者
)

func (s DisRespErrno) IsSuccess() bool {
	return s == Success
}

type DisRespErrno int64

func (s DisRespErrno) ToInt() int {
	return int(s.ToInt64())
}
func (s DisRespErrno) ToInt8() int8 {
	return int8(s.ToInt64())
}
func (s DisRespErrno) ToInt16() int16 {
	return int16(s.ToInt64())
}
func (s DisRespErrno) ToInt32() int32 {
	return int32(s.ToInt64())
}
func (s DisRespErrno) ToInt64() int64 {
	return int64(s)
}
func (s DisRespErrno) ToString() string {
	return fmt.Sprintf("%d", s.ToInt())
}

func (s *DisRespErrno) UnmarshalJSON(b []byte) error {
	var data interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	*s = DisRespErrno(converter.InterfaceToInt64(data))
	return nil
}
