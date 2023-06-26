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

type PolicyOperation uint16

func NewPolicyOperation(num uint16) *PolicyOperation {
	p := PolicyOperation(num)
	return &p
}
func (p *PolicyOperation) ToUInt16() uint16 { return uint16(*p) }
func (p *PolicyOperation) ToBinaryString() string {
	return biu.ToBinaryString(p.ToUInt16())
}

func (p *PolicyOperation) AddAttributeCreateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("10000000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddAttributeReadAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("01000000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddAttributeUpdateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("00100000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddAttributeDeleteAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("00010000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}

func (p *PolicyOperation) RemoveAttributeCreateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("01111111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveAttributeReadAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("10111111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveAttributeUpdateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("11011111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveAttributeDeleteAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("11101111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}

func (p *PolicyOperation) AddContentCreateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("00001000", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddContentReadAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("00000100", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddContentUpdateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("00000010", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) AddContentDeleteAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("00000001", &op)
	num := p.ToUInt16() | op
	return NewPolicyOperation(num)
}

func (p *PolicyOperation) RemoveContentCreateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("11110111", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveContentReadAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("11111011", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveContentUpdateAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("11111101", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}
func (p *PolicyOperation) RemoveContentDeleteAbility() *PolicyOperation {
	var op uint16
	biu.ReadBinaryString("11111110", &op)
	num := p.ToUInt16() & op
	return NewPolicyOperation(num)
}

/*
判断是否有某项能力
*/

func (p *PolicyOperation) HasAttributeCreateAbility() bool {
	var op uint16
	biu.ReadBinaryString("10000000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasAttributeReadAbility() bool {
	var op uint16
	biu.ReadBinaryString("01000000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasAttributeUpdateAbility() bool {
	var op uint16
	biu.ReadBinaryString("00100000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasAttributeDeleteAbility() bool {
	var op uint16
	biu.ReadBinaryString("00010000", &op)
	num := p.ToUInt16() & op
	return num > 0
}

func (p *PolicyOperation) HasContentCreateAbility() bool {
	var op uint16
	biu.ReadBinaryString("00001000", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasContentReadAbility() bool {
	var op uint16
	biu.ReadBinaryString("00000100", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasContentUpdateAbility() bool {
	var op uint16
	biu.ReadBinaryString("00000010", &op)
	num := p.ToUInt16() & op
	return num > 0
}
func (p *PolicyOperation) HasContentDeleteAbility() bool {
	var op uint16
	biu.ReadBinaryString("00000001", &op)
	num := p.ToUInt16() & op
	return num > 0
}
