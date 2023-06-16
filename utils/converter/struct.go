package converter

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}
	return m
}

func StructToMapIgnoreEmpty(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		keyField := t.Field(i)
		keyType := keyField.Type
		mapKeyName := keyField.Tag.Get("json")

		valueField := v.Field(i)
		if valueField.Kind() == reflect.Struct {
			structValue := StructToMapIgnoreEmpty(valueField)
			for k, v := range structValue {
				m[k] = v
			}
		}

		if valueField.Kind() != reflect.Ptr || !valueField.IsNil() {
			//避免结构体里有 unexported field
			if valueField.CanInterface() {
				m[mapKeyName] = valueField.Interface()
			}
		} else {
			m[mapKeyName] = fieldToDefaultValue(keyType)
		}
	}
	return m
}

func StructToMapOmitEmpty(obj interface{}) (result map[string]interface{}, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to json encode struct")
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to json decode struct ")
	}

	return result, nil
}

// StructToMapString 将结构体转换为map[string]string
func StructToMapString(obj interface{}) (map[string]string, error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("must be an struct")
	}

	result := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		jsonName := t.Field(i).Tag.Get("json")
		jsonNameSplit := strings.Split(jsonName, ",")
		name := ""
		if len(jsonNameSplit) > 0 {
			name = jsonNameSplit[0]
		}
		field := v.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}
		result[name] = fieldToString(field)
	}
	return result, nil
}

// StructToMapStringOmitEmpty 将结构体转换为map[string]string（有omitempty的情况）
func StructToMapStringOmitEmpty(obj interface{}) (map[string]string, error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("must be an struct")
	}

	result := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		jsonName := t.Field(i).Tag.Get("json")
		jsonNameSplit := strings.Split(jsonName, ",")
		name := ""
		if len(jsonNameSplit) > 0 {
			name = jsonNameSplit[0]
		}
		field := v.Field(i)

		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}
		if len(jsonNameSplit) > 1 && jsonNameSplit[1] == "omitempty" {
			if field.IsZero() {
				continue
			}
		}
		result[name] = fieldToString(field)
	}
	return result, nil
}

func fieldToString(field reflect.Value) string {
	var data string
	switch field.Kind() {
	case reflect.Ptr:
		{
			data = fieldToString(field.Elem())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		{
			data = strconv.FormatInt(field.Int(), 10)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		{
			data = strconv.FormatUint(field.Uint(), 10)
		}
	case reflect.Bool:
		{
			data = strconv.FormatBool(field.Bool())
		}
	case reflect.String:
		{
			data = field.String()
		}
	case reflect.Float64, reflect.Float32:
		{
			data = Float64ToString(field.Float())
		}
	default:
		jsonData, _ := json.Marshal(field.Interface())
		data = string(jsonData)
	}
	return data
}

func fieldToDefaultValue(field reflect.Type) interface{} {
	var data interface{}
	switch field.Kind() {
	case reflect.Ptr:
		{
			data = fieldToDefaultValue(field.Elem())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		{
			data = 0
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		{
			data = 0
		}
	case reflect.Bool:
		{
			data = false
		}
	case reflect.String:
		{
			data = ""
		}
	case reflect.Float64, reflect.Float32:
		{
			data = 0
		}
	default:
	}
	return data
}

// StructToPHPTransform
//主要用途：按照php格式的结构体转换函数,主要用于json格式转换的兼容。
//如果传入的是**结构体指针**不为空，那么返回结构体
//如果传入的**结构体指针**是nil，那么返回空数组，用于json encode的时候，变成 "[]"
//
//  注意!! 入参必须是结构体指针，其他类型目前不支持，都将返回原本数据
//
//  Example:
//
//  //创建一个结构体
//  testData = person{name: "bob", age: 20 }
//
//  //传入结构体指针
//  result := StructToPHPTransform(&testData)
//
//  //输出结果还是这个结构体指针
//  //output result = *person{name: "bob", age: 20}
//
//  //但是如果传入一个空指针
//  result := StructToPHPTransform(nil)
//
//  //返回值是一个空结构体数组,此时进行json encode，会变成[]
//  //result = []struct{}
func StructToPHPTransform(obj interface{}) interface{} {
	t := reflect.ValueOf(obj)

	//这里处理的是传入普通的 nil,参考: https://golang.org/doc/faq#nil_error
	if obj == nil {
		return []struct{}{}
	}

	//如果不是指针，不管
	if t.Kind() != reflect.Ptr {
		return obj
	}

	//如果是空结构体，返回空数组,用于php格式兼容(php的空对象会被json encode为 [])
	if t.IsNil() {
		return []struct{}{}
	}

	//如果不是结构体不管
	if t.Elem().Kind() != reflect.Struct {
		return obj
	}

	//其他情况不管，返回本身
	return obj
}
