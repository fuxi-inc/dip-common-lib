package idl

import (
	"encoding/json"
	"fmt"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

const (
	OwnerAuthType AuthorizationType = 0 //所有者
	UserAuthType  AuthorizationType = 1 //使用者
)

func (s AuthorizationType) IsOwner() bool {
	return s == OwnerAuthType
}

func (s AuthorizationType) IsUser() bool {
	return s == UserAuthType
}

type AuthorizationType uint8

func (s AuthorizationType) ToInt() int {
	return int(s.ToInt64())
}
func (s AuthorizationType) ToInt8() int8 {
	return int8(s.ToInt64())
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
	*s = AuthorizationType(converter.InterfaceToInt64(data))
	return nil
}
