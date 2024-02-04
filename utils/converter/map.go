package converter

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

func MapFieldFormatToString(value map[string]interface{}) map[string]interface{} {
	for k, v := range value {
		switch vv := v.(type) {
		case int:
			value[k] = IntToString(vv)
		case int32:
			value[k] = Int32ToString(vv)
		case int64:
			value[k] = Int64ToString(vv)
		case float32:
			value[k] = Float32ToString(vv)
		case float64:
			value[k] = Float64ToString(vv)
		case bool:
			value[k] = BoolToString(vv)
		case map[string]interface{}:
			value[k] = MapFieldFormatToString(vv)
		}
	}
	return value
}

// MapStringToStruct 将map[string]string转换为结构体
// 注意result传递的必须是指针，数据将直接解析到result里，请确保result的json tag配置正确
//
// Example :
//
//	 //原始数据
//	 originalData:= map[string]string{
//	    "name" : "bob",
//	    "age" : "100"
//	 }
//	//所需结构体
//	type Person struct{
//	   Name string `json:"name"`
//	   Age  int `json:"age,string"`
//	}
//
//	//调用本方法进行解析
//	bob := &Person{}
//	MapStringToStruct(originalData,bob)
//
//	//output ,bob解析完毕：
//	Bob(Person) { Name: "bob", Age : 100}
func MapStringToStruct(originalData map[string]string, result interface{}) error {
	v := reflect.ValueOf(result)
	if v.Kind() != reflect.Struct && v.Kind() != reflect.Ptr {
		return fmt.Errorf("should pass a struct pointer to this func")
	}

	data, err := json.Marshal(originalData)

	if err != nil {
		return errors.Wrap(err, "failed to marshal dfs response")
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal dfs response")
	}
	return nil
}
