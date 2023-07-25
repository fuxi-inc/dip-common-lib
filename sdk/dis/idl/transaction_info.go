package idl

import "github.com/golang/protobuf/ptypes/timestamp"

// 数据对象的最新Transaction信息
type TransactionInfo struct {
	Doi           string               `json:"data_doi"`     //权属对象
	TransactionID string               `json:"tx_id"`        // 最新交易ID
	Timestamp     *timestamp.Timestamp `json:"tx_timestamp"` // 最新交易时间戳
}
