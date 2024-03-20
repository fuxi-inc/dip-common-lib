package idl

import "github.com/golang/protobuf/ptypes/timestamp"

// 数据对象的最新Transaction信息
type TransactionInfo struct {
	Doi           string               `json:"data_doi"`  //数据对象
	TransactionID string               `json:"tx_id"`     // 最新交易ID
	CreatedAt     *timestamp.Timestamp `json:"create_at"` // 创建时间
	UpdatedAt     *timestamp.Timestamp `json:"update_at"` // 更新时间
}
