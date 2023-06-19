package IDL

import (
    "encoding/json"
    "github.com/fuxi-inc/dip-common-lib/utils/converter"
    "strconv"
)

// SafeJsonNumber 安全的数字，比json.Number 更安全,允许解析一些不合法数据
// 常见于一些坑爹的case:
//       1.未赋值情况： { "number" : "" }, 此时解析为:0
//       2.赋值空对象： { "number" : "{}" }, 此时解析为:0
//       3.赋值空数组： { "number" : "[]" }, 此时解析为:0
//
// 其他场景都将按照: json.Number 本身去解析
type SafeJsonNumber json.Number

func NewJsonNumberFromInt(value int) *SafeJsonNumber {
    tmp := SafeJsonNumber(converter.IntToString(value))
    return &tmp
}

func NewJsonNumberFromInt32(value int32) *SafeJsonNumber {
    tmp := SafeJsonNumber(converter.Int32ToString(value))
    return &tmp
}

func NewJsonNumberFromInt64(value int64) *SafeJsonNumber {
    tmp := SafeJsonNumber(converter.Int64ToString(value))
    return &tmp
}

func (d *SafeJsonNumber) ToInt() int {
    value, err := converter.JsonNumberToInt(json.Number(*d))
    if err != nil {
        return 0
    }
    return value
}

func (d *SafeJsonNumber) ToFloat64() float64 {
    value, err := strconv.ParseFloat(string(*d), 64)
    if err != nil {
        return 0
    }
    return value
}

func (d *SafeJsonNumber) UnmarshalJSON(b []byte) error {
    str := string(b)
    //特殊情况，表示没有值
    if str == `{}` || str == `"{}"` || str == `[]` || str == `"[]"` || str == `""` {
        return nil
    }
    var data json.Number
    err := json.Unmarshal(b, &data)
    if err != nil {
        return err
    }
    *d = SafeJsonNumber(data)
    return nil
}
