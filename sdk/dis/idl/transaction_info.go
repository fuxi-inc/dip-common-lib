package idl

import "github.com/golang/protobuf/ptypes/timestamp"

// 数据对象的最新Transaction信息
type TransactionInfo struct {
	Doi           string               //权属对象
	TransactionID string               // 最新交易ID
	Timestamp     *timestamp.Timestamp // 最新交易时间戳
}
