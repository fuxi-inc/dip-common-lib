package converter

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func InterfaceToInt(inter interface{}) int {
	res := 0
	switch inter.(type) {
	case string:
		res = StringToInt(inter.(string))
		break
	case int:
		res = inter.(int)
		break
	case float64:
		res = int(inter.(float64))
		break
	}
	return res
}

func InterfaceToInt64(inter interface{}) int64 {
	res := inter.(int64)
	return res
}

func InterfaceToString(inter interface{}) string {
	var res string
	switch inter.(type) {
	case string:
		res = inter.(string)
		break
	case int:
		res = IntToString(inter.(int))
		break
	case float64:
		res = Float64ToString(inter.(float64))
		break
	default:
		resStr, _ := json.Marshal(inter)
		res = string(resStr)
		break
	}
	return res
}

func InterfaceToFloat64(inter interface{}) float64 {
	str := fmt.Sprint(inter)
	ret, _ := strconv.ParseFloat(str, 64)
	return ret
}

// InterfaceToFloat64V2 根据type assertion将interface转换成float
func InterfaceToFloat64V2(inter interface{}) (float float64, err error) {
	switch dataType := inter.(type) {
	case string:
		float, err = strconv.ParseFloat(dataType, 10)
		if err != nil {
			return
		}
	case int:
		float = float64(dataType)
	case float64:
		float = dataType
	default:
		return 0, fmt.Errorf("unsupport converter type=%v", dataType)
	}
	return float, nil
}

func InterfaceToJsonNum(inter interface{}) json.Number {
	tmpNum, ok := inter.(json.Number)
	if ok {
		return tmpNum
	}
	return "0"
}
