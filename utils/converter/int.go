package converter

import "strconv"

type IntConverter interface {
	ToInt() int
}

type Int8Converter interface {
	ToInt8() int8
}

type Int16Converter interface {
	ToInt16() int16
}

type Int32Converter interface {
	ToInt32() int32
}

type Int64Converter interface {
	ToInt64() int64
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// IntToString int转换为string，比如1 => "1",
// 特别注意： 0 => "0"
func IntToString(i int) string {
	return strconv.Itoa(i)
}

func IntToBool(i int) bool {
	if i > 0 {
		return true
	}
	return false
}

// IntMeanToString 特殊处理了0值的转换，主要是为了兼容某些下游的特殊需要，0值会转换为空字符串
//其他逻辑，参考 IntToString
func IntMeanToString(i int) string {
	if i > 0 {
		return IntToString(i)
	}
	return ""
}
