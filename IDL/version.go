package IDL

import (
	"github.com/hashicorp/go-version"
)

// VersionString 用于表示一个代表版本号的字符串,支持json tag，可以直接从入参里解析
//type Request struct{
//   AppVersion VersionString `json:"app_version"`
//}
//注意：初始化不合法的版本号字符串，会报错。比如传递：空字符串:"",错误字符串:"2.s.3xx"
type VersionString string

// GreaterThan 当前版本号是否"大于"输入版本号
func (v VersionString) GreaterThan(v2Str VersionString) (bool, error) {
	v1, err := version.NewVersion(v.String())
	if err != nil {
		return false, err
	}
	v2, err := version.NewVersion(v2Str.String())
	if err != nil {
		return false, err
	}

	return v1.GreaterThan(v2), nil
}

// GreaterOrEqualThan 当前版本号是否"大于等于"输入版本号
func (v VersionString) GreaterOrEqualThan(v2Str VersionString) (bool, error) {
	v1, err := version.NewVersion(v.String())
	if err != nil {
		return false, err
	}
	v2, err := version.NewVersion(v2Str.String())
	if err != nil {
		return false, err
	}

	return v1.GreaterThanOrEqual(v2), nil
}

// LessOrEqualThan 当前版本号是否"小于等于"输入版本号
func (v VersionString) LessOrEqualThan(v2Str VersionString) (bool, error) {
	v1, err := version.NewVersion(v.String())
	if err != nil {
		return false, err
	}
	v2, err := version.NewVersion(v2Str.String())
	if err != nil {
		return false, err
	}

	return v1.LessThanOrEqual(v2), nil
}

// LessThan 当前版本号是否"小于"输入版本号
func (v VersionString) LessThan(v2Str VersionString) (bool, error) {
	v1, err := version.NewVersion(v.String())
	if err != nil {
		return false, err
	}
	v2, err := version.NewVersion(v2Str.String())
	if err != nil {
		return false, err
	}

	return v1.LessThan(v2), nil
}

// EqualsTo 当前版本号是否"等于"输入版本号
func (v VersionString) EqualsTo(v2Str VersionString) (bool, error) {
	v1, err := version.NewVersion(v.String())
	if err != nil {
		return false, err
	}
	v2, err := version.NewVersion(v2Str.String())
	if err != nil {
		return false, err
	}

	return v1.Equal(v2), nil
}

func (v VersionString) String() string {
	return string(v)
}
