package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/imroc/biu"
)

type PermissionCondition struct {
	TotalWeight         float64 `json:"total_weight"`
	TotalDecisionNumber float64 `json:"total_decision_number"`
}

type Permission struct {
	Operations   PermissionOperation  `json:"operations,omitempty"`    //基本操作权限，uint16，前4位分别表示对数据对象属性的CRUD （创建、读取、更新、删除）权限，后4位分别表示对数据内容的CRUD权限
	AlgorithmDOI string               `json:"algorithm_doi,omitempty"` //处理算法DOI，string，符合数据对象标识规范
	Weight       float64              `json:"weight,omitempty"`        //权重，uint16，取值表示百分比
	StartAt      *IDL.Time            `json:"start_at,omitempty"`      //生效时间 ,时间戳
	ExpiredAt    *IDL.Time            `json:"expired_at,omitempty"`    //生效时间，时间戳
	Condition    *PermissionCondition `json:"condition,omitempty"`
}

/**
PermissionOperation, 基本的操作权限
目前已经在定义中的有效作用位由低位到高位共有9位：
	4-1位，分别代表对数据内容的CRUD权限;
	8-5位，分别代表对数据属性的CRUD权限；
	9位, 代表是否复用数据owner的权限，当此位置为1时，其他低位均失效，即 100000000和111111111 代表的权限相同
*/
type PermissionOperation uint16

func NewPermissionOperation(num uint16) *PermissionOperation {
	p := PermissionOperation(num)
	return &p
}
func (p *PermissionOperation) ToUInt16() uint16 { return uint16(*p) }
func (p *PermissionOperation) ToBinaryString() string {
	return biu.ToBinaryString(p.ToUInt16())
}

func (p *PermissionOperation) AddParentAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("100000000", &op)
	num := p.ToUInt16() | op
	return NewPermissionOperation(num)
}

func (p *PermissionOperation) RemoveParentAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("011111111", &op)
	num := p.ToUInt16() & op
	return NewPermissionOperation(num)
}

//8-5位，对数据属性的CRUD权限 的操作

func (p *PermissionOperation) AddAttributeCreateAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("010000000", &op)
	num := p.ToUInt16() | op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) AddAttributeReadAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("001000000", &op)
	num := p.ToUInt16() | op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) AddAttributeUpdateAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("000100000", &op)
	num := p.ToUInt16() | op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) AddAttributeDeleteAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("000010000", &op)
	num := p.ToUInt16() | op
	return NewPermissionOperation(num)
}

func (p *PermissionOperation) RemoveAttributeCreateAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("101111111", &op)
	num := p.ToUInt16() & op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) RemoveAttributeReadAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("110111111", &op)
	num := p.ToUInt16() & op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) RemoveAttributeUpdateAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("111011111", &op)
	num := p.ToUInt16() & op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) RemoveAttributeDeleteAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("111101111", &op)
	num := p.ToUInt16() & op
	return NewPermissionOperation(num)
}

//4-1位，分别代表对数据内容的CRUD权限；

func (p *PermissionOperation) AddContentCreateAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("000001000", &op)
	num := p.ToUInt16() | op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) AddContentReadAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("000000100", &op)
	num := p.ToUInt16() | op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) AddContentUpdateAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("000000010", &op)
	num := p.ToUInt16() | op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) AddContentDeleteAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("000000001", &op)
	num := p.ToUInt16() | op
	return NewPermissionOperation(num)
}

func (p *PermissionOperation) RemoveContentCreateAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("111110111", &op)
	num := p.ToUInt16() & op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) RemoveContentReadAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("111111011", &op)
	num := p.ToUInt16() & op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) RemoveContentUpdateAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("111111101", &op)
	num := p.ToUInt16() & op
	return NewPermissionOperation(num)
}
func (p *PermissionOperation) RemoveContentDeleteAbility() *PermissionOperation {
	var op uint16
	biu.ReadBinaryString("111111110", &op)
	num := p.ToUInt16() & op
	return NewPermissionOperation(num)
}

/*
判断是否有某项能力
*/

func (p *PermissionOperation) HasAttributeCreateAbility() bool {
	var op uint16
	biu.ReadBinaryString("010000000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PermissionOperation) HasAttributeReadAbility() bool {
	var op uint16
	biu.ReadBinaryString("001000000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PermissionOperation) HasAttributeUpdateAbility() bool {
	var op uint16
	biu.ReadBinaryString("000100000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PermissionOperation) HasAttributeDeleteAbility() bool {
	var op uint16
	biu.ReadBinaryString("000010000", &op)
	num := p.ToUInt16() & op
	return num > 0
}

func (p *PermissionOperation) HasContentCreateAbility() bool {
	var op uint16
	biu.ReadBinaryString("000001000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PermissionOperation) HasContentReadAbility() bool {
	var op uint16
	biu.ReadBinaryString("000000100", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PermissionOperation) HasContentUpdateAbility() bool {
	var op uint16
	biu.ReadBinaryString("000000010", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PermissionOperation) HasContentDeleteAbility() bool {
	var op uint16
	biu.ReadBinaryString("000000001", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PermissionOperation) HasParentAbility() bool {
	var op uint16
	biu.ReadBinaryString("100000000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
