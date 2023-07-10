package idl

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/biu"

	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

const (
	OwnerAuthType AuthorizationType = 0   //所有者
	UserAuthType  AuthorizationType = 255 //使用者
)

func (s AuthorizationType) IsOwner() bool {
	var op uint8
	biu.ReadBinaryString("10000000", &op)
	num := s.ToUInt8() & op
	return num <= 0
}

func (s AuthorizationType) IsUser() bool {
	return !s.IsOwner()
}

type AuthorizationType uint8

func (s AuthorizationType) ToInt() int {
	return int(s.ToInt64())
}
func (s AuthorizationType) ToUInt8() uint8 {
	return uint8(s)
}
func (s AuthorizationType) ToInt16() int16 {
	return int16(s.ToInt64())
}
func (s AuthorizationType) ToInt32() int32 {
	return int32(s.ToInt64())
}
func (s AuthorizationType) ToInt64() int64 {
	return int64(s)
}
func (s AuthorizationType) ToString() string {
	return fmt.Sprintf("%d", s.ToInt())
}

func (s *AuthorizationType) UnmarshalJSON(b []byte) error {
	var data interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	var authType uint8
	switch data.(type) {
	case string:
		authType = uint8(converter.StringToInt8(data.(string)))
	case int, int16, int32, int64:
		authType = uint8(data.(int))
	case float32, float64:
		authType = uint8(data.(float64))
	}

	*s = AuthorizationType(authType)
	return nil
}
