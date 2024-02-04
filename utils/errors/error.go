package errors

import (
	"fmt"
	"runtime/debug"
)

type DErr struct {
	//返回错误信息
	errmsg string `json:"errmsg"`
}

func NewZError() *DErr {
	return &DErr{}
}

func (z *DErr) SetErrmsg(msg string) *DErr {
	z.errmsg = msg
	return z
}

func (z DErr) Error() string {
	return fmt.Sprintf("errmsg=%s", z.ErrMsg())
}
func (z DErr) ErrMsg() string {
	return z.errmsg
}

func NewError(err error) *DErr {
	if err == nil {
		return NewZError().SetErrmsg(string(debug.Stack()))
	}
	return NewZError().SetErrmsg(err.Error() + " " + string(debug.Stack()))
}
