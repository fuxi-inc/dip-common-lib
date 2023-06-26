package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/imroc/biu"
)

type PolicyCondition struct {
	TotalWeight         float64 `json:"total_weight"`
	TotalDecisionNumber float64 `json:"total_decision_number"`
}

type Policy struct {
	Operations   PolicyOperation  `json:"operations,omitempty"`    //基本操作权限，uint16，前4位分别表示对数据对象属性的CRUD （创建、读取、更新、删除）权限，后4位分别表示对数据内容的CRUD权限
	AlgorithmDOI string           `json:"algorithm_doi,omitempty"` //处理算法DOI，string，符合数据对象标识规范
	Weight       float64          `json:"weight,omitempty"`        //权重，uint16，取值表示百分比
	StartAt      *IDL.Time        `json:"start_at,omitempty"`      //生效时间 ,时间戳
	ExpiredAt    *IDL.Time        `json:"expired_at,omitempty"`    //生效时间，时间戳
	Condition    *PolicyCondition `json:"condition,omitempty"`
}

/**
PolicyOperation, 基本的操作权限
目前已经在定义中的有效作用位由低位到高位共有9位：
	4-1位，分别代表对数据内容的CRUD权限;
	8-5位，分别代表对数据属性的CRUD权限；
	9位, 代表是否复用数据owner的权限，当此位置为1时，其他低位均失效，即 100000000和111111111 代表的权限相同
*/
type PolicyOperation uint16

func NewPolicyOperation(num uint16) *PolicyOperation {
	p := PolicyOperation(num)
	return &p
}
func (p *PolicyOperation) ToUInt16() uint16 { return uint16(*p) }
func (p *PolicyOperation) ToBinaryString() string {
	return biu.ToBinaryString(p.ToUInt16())
}

func (p *PolicyOperation) AddParentAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("100000000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}

func (p *PolicyOperation) RemoveParentAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("011111111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}

//8-5位，对数据属性的CRUD权限 的操作

func (p *PolicyOperation) AddAttributeCreateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("010000000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddAttributeReadAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("001000000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddAttributeUpdateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("000100000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddAttributeDeleteAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("000010000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}

func (p *PolicyOperation) RemoveAttributeCreateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("101111111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveAttributeReadAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("110111111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveAttributeUpdateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("111011111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveAttributeDeleteAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("111101111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}

//4-1位，分别代表对数据内容的CRUD权限；

func (p *PolicyOperation) AddContentCreateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("000001000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddContentReadAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("000000100", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddContentUpdateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("000000010", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddContentDeleteAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("000000001", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}

func (p *PolicyOperation) RemoveContentCreateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("111110111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveContentReadAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("111111011", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveContentUpdateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("111111101", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveContentDeleteAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("111111110", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}

/*
判断是否有某项能力
*/

func (p *PolicyOperation) HasAttributeCreateAbility() bool {
	var op uint16
	biu.ReadBinaryString("010000000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasAttributeReadAbility() bool {
	var op uint16
	biu.ReadBinaryString("001000000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasAttributeUpdateAbility() bool {
	var op uint16
	biu.ReadBinaryString("000100000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasAttributeDeleteAbility() bool {
	var op uint16
	biu.ReadBinaryString("000010000", &op)
	num := p.ToUInt16() & op
	return num > 0
}

func (p *PolicyOperation) HasContentCreateAbility() bool {
	var op uint16
	biu.ReadBinaryString("000001000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasContentReadAbility() bool {
	var op uint16
	biu.ReadBinaryString("000000100", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasContentUpdateAbility() bool {
	var op uint16
	biu.ReadBinaryString("000000010", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasContentDeleteAbility() bool {
	var op uint16
	biu.ReadBinaryString("000000001", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasParentAbility() bool {
	var op uint16
	biu.ReadBinaryString("100000000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
