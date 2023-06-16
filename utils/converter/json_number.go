package converter

import "encoding/json"

func JsonNumberPointerToInt64(number *json.Number) int64 {
    if number == nil {
        return 0
    }
    ret, _ := number.Int64()
    return ret
}

func JsonNumberPointerToFloat64(number *json.Number) float64 {
    if number == nil {
        return 0
    }
    ret, _ := number.Float64()
    return ret
}

func JsonNumberPointerToString(number *json.Number) string {
    if number == nil {
        return ""
    }
    return number.String()
}

// JsonNumberToInt 将json.Number解析为int，优先用Int64解析，其次试图将float64转换为int,都失败则返回错误
func JsonNumberToInt(num json.Number) (int, error) {
    //入参可能有字符串形式，需要int解析
    data, err := num.Int64()
    if err == nil {
        return int(data), nil
    }

    //android传递的是2.0这个数字，此时无法使用 Int64方法转换，json.Number会直接报错，所以尝试用float转换
    floatData, err := num.Float64()
    if err == nil {
        return int(floatData), nil
    }

    stringData := num.String()
    if stringData != "" {
        return StringToInt(stringData), nil
    }

    return 0, err
}
