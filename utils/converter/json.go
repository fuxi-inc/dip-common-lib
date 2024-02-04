package converter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// SafeUnmarshal 如果是数字，则会转化为jsonNumber
func SafeUnmarshal(str []byte, ret interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(str))
	decoder.UseNumber()
	err := decoder.Decode(ret)
	if err != nil {
		return fmt.Errorf("unmarshal err:%v", err)
	}
	return nil
}

// Unescape 去除json字符串中的转义字符，尤其用于端上传来的某一个参数，是一个json字符串
// Example:
//
//	//1.转义json,不带双引号,无法直接进行 json.Unmarshal, 所以需要进行反转义，返回值为：
//	input := `{\"name\":\"suitingwei\",\"age\":13}`
//	//进行反转义, output = {"name":"suitingwei","age":13}
//	output,_ :=converter.Unescape(input)
//
//	//2.转义json,带双引号
//	input := `"{\"name\":\"suitingwei\",\"age\":13}"`
//	//进行反转义, output = {"name":"suitingwei","age":13}
//	output,_ :=converter.Unescape(input)
//
//	//3.不是转义的json，直接返回，主要是给一些IDL的 UnmarshalJSON方法提供便利，避免合法的json在这个函数调用后包凑
//	input := `{"name":"suitingwei","age":13}"`
//	//进行反转义, output = {"name":"suitingwei","age":13}
//	output,_ :=converter.Unescape(input)
//
// 注意：如果传入的不是转义后的字符串，会返回空
func Unescape(str string) (string, error) {
	//json至少有2个字符, {},如果小了，就不处理
	//如果入参是空字符串,""，那么直接返回
	if len(str) <= 2 {
		return str, nil
	}

	//bugfix： 如果是嵌套json，某一个字段是一个转义字符串，那么整体不需要进行UnQuote
	//所以这里替换为，用 json.Valid判断是否json。 但是json.Valid() 对于字符串都返回true，比如入参是: `"这里面是不合法json，但是被双引号包着"`
	if str[0] != '"' && json.Valid([]byte(str)) {
		return str, nil
	}

	var unquote string
	var err error
	//如果已经是加了双引号，不再加,支持  "{\"name\":\"suitingwei\",\"age\":13}"
	if str[0] == '"' && str[len(str)-1] == '"' {
		unquote, err = strconv.Unquote(str)
	} else {
		unquote, err = strconv.Unquote(`"` + str + `"`)
	}

	if err != nil {
		return "", errors.Wrapf(err, "failed to unescape data: %v", str)
	}
	return unquote, err
}
