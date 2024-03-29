package IDL

import (
	"encoding/json"

	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

const (
	RequestAuthNotifyType              NotifyType = 1 //请求授权消息类型
	ConfirmAuthNotifyType              NotifyType = 2 //确认并接受授权的消息类型
	DenyAuthNotifyType                 NotifyType = 3 //拒绝授权的消息类型
	RevokeAuthNotifyType               NotifyType = 4 //撤销授权的消息类型
	RequestExchangeOwnershipNotifyType NotifyType = 5 //请求交换所有权消息类型
)

type NotifyType int64

func (s NotifyType) ToInt() int {
	return int(s.ToInt64())
}
func (s NotifyType) ToInt8() int8 {
	return int8(s.ToInt64())
}
func (s NotifyType) ToInt16() int16 {
	return int16(s.ToInt64())
}
func (s NotifyType) ToInt32() int32 {
	return int32(s.ToInt64())
}
func (s NotifyType) ToInt64() int64 {
	return int64(s)
}
func (s NotifyType) ToString() string {
	return converter.Int64ToString(s.ToInt64())
}

func (s *NotifyType) UnmarshalJSON(b []byte) error {
	var data interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	var notifyType int64
	switch data.(type) {
	case string:
		notifyType = converter.StringToInt64(data.(string))
	case int, int16, int32, int64:
		notifyType = int64(data.(int))
	case float32, float64:
		notifyType = int64(data.(float64))
	}

	*s = NotifyType(notifyType)
	return nil
}
