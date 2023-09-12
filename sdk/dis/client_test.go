package dis

import (
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func GetPubKeyString() string {
	file, err := os.Open("../testpkg/user/alice/public.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	keybase64 := base64.StdEncoding.EncodeToString(block.Bytes)
	return keybase64

}
func GetPrivKeyString() string {
	file, err := os.Open("../testpkg/user/alice/private.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	keybase64 := base64.StdEncoding.EncodeToString(block.Bytes)
	return keybase64

}
func Test_DOCreate(t *testing.T) {
	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test1.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	if err != nil {
		print(err.Error())
	}
	assert.Nil(t, err)
	sign.Signature = Signature
	whois := &idl.RegistrationData{
		Doi:     "25test1.viv.cn.",
		Contact: []string{"http://www.baidu.com"},
	}
	request := &idl.ApiDOCreateRequest{
		Doi:           "25test1.viv.cn.",
		DwDoi:         "25test1.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
		WhoisData:     whois,
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://192.168.10.232:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOCreate(ctx, request)
	print(response.Errmsg)
	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

}

func Test_DOCreate2(t *testing.T) {
	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test2.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	if err != nil {
		print(err.Error())
	}
	assert.Nil(t, err)
	sign.Signature = Signature
	whois := &idl.RegistrationData{
		Doi:     "25test2.viv.cn.",
		Contact: []string{"http://www.baidu.com"},
	}
	request := &idl.ApiDOCreateRequest{
		Doi:           "25test2.viv.cn.",
		DwDoi:         "25test2.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
		WhoisData:     whois,
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://192.168.10.232:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOCreate(ctx, request)
	print(response.Errmsg)
	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

}

func Test_DOCreate3(t *testing.T) {
	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test3.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	if err != nil {
		print(err.Error())
	}
	assert.Nil(t, err)
	sign.Signature = Signature
	whois := &idl.RegistrationData{
		Doi:     "25test3.viv.cn.",
		Contact: []string{"http://www.baidu.com"},
	}
	request := &idl.ApiDOCreateRequest{
		Doi:           "25test3.viv.cn.",
		DwDoi:         "25test3.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
		WhoisData:     whois,
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://192.168.10.232:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOCreate(ctx, request)
	print(response.Errmsg)
	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

}

func Test_DOCreate4(t *testing.T) {
	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test1.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	if err != nil {
		print(err.Error())
	}
	assert.Nil(t, err)
	sign.Signature = Signature
	whois := &idl.RegistrationData{
		Doi:     "25test1_data.viv.cn.",
		Contact: []string{"http://www.baidu.com"},
	}
	request := &idl.ApiDOCreateRequest{
		Doi:           "25test1_data.viv.cn.",
		DwDoi:         "25test1.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
		WhoisData:     whois,
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://192.168.10.232:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOCreate(ctx, request)
	print(response.Errmsg)
	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

}

func Test_DOCreate5(t *testing.T) {
	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test1.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	if err != nil {
		print(err.Error())
	}
	assert.Nil(t, err)
	sign.Signature = Signature
	whois := &idl.RegistrationData{
		Doi:     "data.viv.cn.",
		Contact: []string{"http://www.baidu.com"},
	}
	request := &idl.ApiDOCreateRequest{
		Doi:           "data.viv.cn.",
		DwDoi:         "25test1.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
		WhoisData:     whois,
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://192.168.10.232:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOCreate(ctx, request)
	print(response.Errmsg)
	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

}

// func Test_DOUpdate(t *testing.T) {

// 	sign := IDL.SignatureData{}
// 	sign.OperatorDoi = "alice.viv.cn."
// 	sign.SignatureNonce = "123456"
// 	Signature, err := sign.CreateSignature(GetPrivKeyString())
// 	assert.Nil(t, err)
// 	sign.Signature = Signature

// 	// 更新数据标识
// 	request := &idl.ApiDOUpdateRequest{
// 		Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
// 		NewDoi: "XXX.viv.cn.",
// 		DwDoi:  "alice.viv.cn.",

// 		SignatureData: sign,
// 	}
// 	client := NewClient().
// 		InitLogger(zap.NewExample()).
// 		// TODO: 添加disq的host名称
// 		InitDis("http://39.107.180.231:8991")

// 	// 执行被测试的函数
// 	ctx := &gin.Context{}
// 	response, err := client.ApiDOUpdate(ctx, request)

// 	// 断言函数返回的错误为 nil
// 	assert.Nil(t, err)

// 	// 判断 Errno 是否为 0
// 	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
// 	// 更新公钥
// 	request = &idl.ApiDOUpdateRequest{
// 		Doi:           "XXX.viv.cn",
// 		PubKey:        GetPubKeyString(),
// 		DwDoi:         "usera.viv.cn",
// 		SignatureData: sign,
// 	}
// 	response, err = client.ApiDOUpdate(ctx, request)

// 	// 断言函数返回的错误为 nil
// 	assert.Nil(t, err)

// 	// 判断 Errno 是否为 0
// 	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
// 	// 更新数据地址及摘要

// 	digest := &idl.DataDigest{
// 		Algorithm: "SHA256",
// 		Result:    "sha256",
// 	}

// 	classgrade := &idl.ClassificationAndGrading{
// 		Class: 1024,
// 		Grade: 2048,
// 	}

// 	auth := &idl.DataAuthorization{
// 		Confirmation: "XXX",
// 	}

// 	request = &idl.ApiDOUpdateRequest{
// 		Doi:                      "XXX.viv.cn",
// 		Dar:                      "resource.example.com/path",
// 		Digest:                   digest,
// 		DwDoi:                    "usera.viv.cn",
// 		Authorization:            auth, // 更新其中的Conformation确权信息
// 		ClassificationAndGrading: classgrade,
// 		SignatureData:            sign,
// 	}
// 	response, err = client.ApiDOUpdate(ctx, request)

// 	// 断言函数返回的错误为 nil
// 	assert.Nil(t, err)

// 	// 判断 Errno 是否为 0
// 	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
// 	// 数据所有者更新自己的权益

// 	desc := &idl.PermissionDescription{
// 		PermissionDoi: "XXX.viv.cn",
// 		CreatorDoi:    "yyy.viv.cn",
// 	}

// 	auth = &idl.DataAuthorization{
// 		Type:        0,
// 		Description: desc,
// 	}

// 	request = &idl.ApiDOUpdateRequest{
// 		Doi:           "XXX.viv.cn",
// 		Authorization: auth, // 更新其中的Type和Description
// 		SignatureData: sign,
// 	}
// 	response, err = client.ApiDOUpdate(ctx, request)

// 	// 断言函数返回的错误为 nil
// 	assert.Nil(t, err)

// 	// 判断 Errno 是否为 0
// 	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
// 	// 更新联系方式

// 	whois := &idl.RegistrationData{
// 		Doi:     "XXX.viv.cn",
// 		Contact: []string{"xxx", "yyy"},
// 	}

// 	request = &idl.ApiDOUpdateRequest{
// 		Doi:           "XXX.viv.cn",
// 		WhoisData:     whois,
// 		SignatureData: sign,
// 	}
// 	response, err = client.ApiDOUpdate(ctx, request)

// 	// 断言函数返回的错误为 nil
// 	assert.Nil(t, err)

// 	// 判断 Errno 是否为 0
// 	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

// }

// 更新公钥
func Test_DOUpdate1(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test1.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	assert.Nil(t, err)
	sign.Signature = Signature

	// 更新数据标识
	//request := &idl.ApiDOUpdateRequest{
	//	Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
	//	NewDoi: "XXX.viv.cn.",
	//	DwDoi:  "alice.viv.cn.",
	//
	//	SignatureData: sign,
	//}

	request := &idl.ApiDOUpdateRequest{
		Doi:           "25test1_data.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/cindy/public.hex")),
		DwDoi:         "25test1.viv.cn.",
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOUpdate(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
	// 更新公钥

	// 断言函数返回的错误为 nil

}

// dar digest confirmation
func Test_DOUpdate2(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test1.viv.cn."
	sign.SignatureNonce = "123457"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	assert.Nil(t, err)
	sign.Signature = Signature

	digest := &idl.DataDigest{
		Algorithm: "SHA256",
		Result:    "sha256",
	}

	auth := &idl.DataAuthorization{
		Doi:          "25test1.viv.cn.",
		Confirmation: "xxxx",
	}

	request := &idl.ApiDOUpdateRequest{
		Doi:           "25test1_data.viv.cn.",
		Dar:           "resource.example.com/path",
		Digest:        digest,
		DwDoi:         "25test1.viv.cn.",
		Authorization: auth,
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOUpdate(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)
	fmt.Print(response)
	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

}

// 分类分级
func Test_DOUpdate3(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test1.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	assert.Nil(t, err)
	sign.Signature = Signature

	// 更新数据标识
	//request := &idl.ApiDOUpdateRequest{
	//	Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
	//	NewDoi: "XXX.viv.cn.",
	//	DwDoi:  "alice.viv.cn.",
	//
	//	SignatureData: sign,
	//}
	classgrade := &idl.ClassificationAndGrading{
		Class: 1024,
		Grade: 2048,
	}
	request := &idl.ApiDOUpdateRequest{
		Doi:                      "25test1_data.viv.cn.",
		ClassificationAndGrading: classgrade,
		DwDoi:                    "25test1.viv.cn.",
		SignatureData:            sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOUpdate(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
	// 更新公钥

	// 断言函数返回的错误为 nil

}

// 更新所有权信息
func Test_DOUpdate4(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test1.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	assert.Nil(t, err)
	sign.Signature = Signature

	// 更新数据标识
	//request := &idl.ApiDOUpdateRequest{
	//	Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
	//	NewDoi: "XXX.viv.cn.",
	//	DwDoi:  "alice.viv.cn.",
	//
	//	SignatureData: sign,
	//}
	desc := &idl.PermissionDescription{
		PermissionDoi: "XXX_alice3_data3.viv.cn.",
		CreatorDoi:    "25test1.viv.cn.",
	}

	auth := &idl.DataAuthorization{
		Doi:         "25test1.viv.cn.",
		Type:        0,
		Description: desc,
	}
	request := &idl.ApiDOUpdateRequest{
		Doi:           "25test1_data.viv.cn.",
		Authorization: auth,
		DwDoi:         "25test1.viv.cn.",
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOUpdate(ctx, request)
	fmt.Println(response.Errmsg)
	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
	// 更新公钥

	// 断言函数返回的错误为 nil

}

func Test_DOUpdate5(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test1.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	assert.Nil(t, err)
	sign.Signature = Signature

	// 更新数据标识
	//request := &idl.ApiDOUpdateRequest{
	//	Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
	//	NewDoi: "XXX.viv.cn.",
	//	DwDoi:  "alice.viv.cn.",
	//
	//	SignatureData: sign,
	//}
	//desc := &idl.PermissionDescription{
	//	PermissionDoi: "XXX.viv.cn",
	//	CreatorDoi:    "yyy.viv.cn",
	//}

	whois := &idl.RegistrationData{
		Doi:     "XXXxxx.viv.cn",
		Contact: []string{"http://test1.baidu.com", "yyy", "zzz"},
	}
	request := &idl.ApiDOUpdateRequest{
		Doi:           "25test1.viv.cn.",
		WhoisData:     whois,
		DwDoi:         "25test1.viv.cn.",
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOUpdate(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
	// 更新公钥

	// 断言函数返回的错误为 nil

}

// new doi
func Test_DOUpdate6(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "25test1.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	assert.Nil(t, err)
	sign.Signature = Signature

	// 更新数据标识
	//request := &idl.ApiDOUpdateRequest{
	//	Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
	//	NewDoi: "XXX.viv.cn.",
	//	DwDoi:  "alice.viv.cn.",
	//
	//	SignatureData: sign,
	//}
	//desc := &idl.PermissionDescription{
	//	PermissionDoi: "XXX.viv.cn",
	//	CreatorDoi:    "yyy.viv.cn",
	//}

	request := &idl.ApiDOUpdateRequest{
		Doi:           "25test1.viv.cn.",
		NewDoi:        "25test2.viv.cn.",
		DwDoi:         "25test1.viv.cn.",
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOUpdate(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
	// 更新公钥

	// 断言函数返回的错误为 nil

}

func Test_DODelete(t *testing.T) {
	sign := IDL.SignatureData{}
	sign.OperatorDoi = "test_pic_pm2.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	assert.Nil(t, err)
	sign.Signature = Signature

	// 删除数据标识
	request := &idl.ApiDODeleteRequest{
		Doi: "test_pic_pm2.viv.cn.",

		SignatureData: sign,
	}
	log.Println("request:", converter.ToString(request))
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDODelete(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
}
func Test_DOQuery(t *testing.T) {

	// 设置测试数据
	request := &idl.ApiDOQueryRequest{
		// TODO: 设置测试doi
		Doi: "25test1_data.viv.cn.",
		Type: []idl.SearchType{
			idl.Dar,
			idl.Owner,
			idl.PubKey,
			idl.Digest,
			idl.ClassGrade,
		},
	}

	ctx := &gin.Context{}

	// 创建一个 Client 实例
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDisQ("http://39.107.180.231:8053")

	// 执行被测试的函数
	response, err := client.ApiDOQuery(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

	log.Println(response.Data.PubKey)
	log.Println(response.Data.Owner)
	log.Println(response.Data.ClassificationAndGrading.Class)
	log.Println(response.Data.ClassificationAndGrading.Grade)

	log.Println(response.Data.Dar)
	log.Println(response.Data.Digest.Algorithm)
	log.Println(response.Data.Digest.Result)

}

func Test_DOAuthQuery(t *testing.T) {

	dudoi := "25test2.viv.cn."

	// 设置测试数据
	request := &idl.ApiDOAuthQueryRequest{
		// TODO: 设置测试doi
		Doi:   "25test1_data.viv.cn.",
		DuDoi: dudoi,
		Type: []idl.SearchType{
			idl.Auth,
		},
	}

	ctx := &gin.Context{}

	// 创建一个 Client 实例
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDisQ("http://39.107.180.231:8053")

	// 执行被测试的函数
	response, err := client.ApiDOAuthQuery(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
	log.Println("----->auth: ", converter.ToString(response.Data))
	log.Println(response.Data.Auth[dudoi].Doi)
	log.Println(response.Data.Auth[dudoi].Type)
	log.Println(response.Data.Auth[dudoi].Confirmation)
	log.Println(response.Data.Auth[dudoi].Description.CreatorDoi)
	log.Println(response.Data.Auth[dudoi].Description.PermissionDoi)

}

func TestClient_ApiDOCreate(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiDOCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ApiDisResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "[应用测试用户] 注册用户",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOCreateRequest{
					Doi:    "alice_create_by_lyl4.viv.cn.",
					DwDoi:  "alice_create_by_lyl4.viv.cn.",
					PubKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					WhoisData: &idl.RegistrationData{
						Doi: "alice_create_by_lyl2.viv.cn.",
						Contact: []string{
							"https://segmentfault.com/q/1010000043984824",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "[应用测试用户] 注册数岛",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOCreateRequest{
					Doi:           "alice_create_by_lyl_dao.viv.cn.",
					DwDoi:         "alice_create_by_lyl.viv.cn.",
					PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[数岛递归读取测试] 注册用户", //注册dao_alice,dao_bob,dao_cindy,dao_dale
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOCreateRequest{
					Doi:    "dao_dale_by_lyl.viv.cn.",
					DwDoi:  "dao_dale_by_lyl.viv.cn.",
					PubKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					WhoisData: &idl.RegistrationData{
						Doi: "dao_dale_by_lyl.viv.cn.",
						Contact: []string{
							"https://segmentfault.com/q/1010000043984824",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("dao_dale_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiDOCreate(tt.args.ctx, tt.args.request)
			log.Println("--->test_name", tt.name)
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->err:", err)
		})
	}
}

func TestClient_ApiAuthInit(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiAuthInitRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ApiDisResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Du授权测试用例",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://192.168.10.232:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiAuthInitRequest{
					DataDoi: "25test1_data.viv.cn.",
					Authorization: idl.DataAuthorization{
						Doi:  "25test2.viv.cn.",
						Type: idl.UserAuthType,
						Confirmation: func() string {
							sign, err := IDL.NewSignatureData().SetOperator("").SetNonce("sha256").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							fmt.Println("SignByPK-->:", sign, err)
							return sign
						}(),
						Description: &idl.PermissionDescription{
							PermissionDoi: "data.viv.cn.",
							CreatorDoi:    "25test2.viv.cn.",
							Key:           "",
						},
					},
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "25test2.viv.cn.",
						SignatureNonce: "123456",
						Signature: func() string {
							sign, _ := IDL.NewSignatureData().SetOperator("25test2.viv.cn.").SetNonce("123456").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							return sign
						}(),
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "Dw授权测试用例",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://192.168.10.232:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiAuthInitRequest{
					DataDoi: "25test1_data.viv.cn.",
					Authorization: idl.DataAuthorization{
						Doi:  "25test3.viv.cn.",
						Type: idl.UserAuthType,
						Description: &idl.PermissionDescription{
							PermissionDoi: "data.viv.cn.",
							CreatorDoi:    "25test1.viv.cn.",
							Key:           "",
						},
					},
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "25test1.viv.cn.",
						SignatureNonce: "123456",
						Signature: func() string {
							sign, _ := IDL.NewSignatureData().SetOperator("25test1.viv.cn.").SetNonce("123456").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							return sign
						}(),
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[应用测试用户-代理修改权限] 授权",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "http://39.107.180.231:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiAuthInitRequest{
					DataDoi: "update_user_a_file_aa.viv.cn.",
					Authorization: idl.DataAuthorization{
						Doi:  "update_user_b.viv.cn.",
						Type: idl.UserAuthType,
						Description: &idl.PermissionDescription{
							PermissionDoi: "update_user_a_permission_aa.viv.cn.",
							CreatorDoi:    "update_user_a.viv.cn.",
							Key:           "",
						},
					},
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "update_user_a.viv.cn.",
						SignatureNonce: "123456",
						Signature: func() string {
							sign, _ := IDL.NewSignatureData().SetOperator("update_user_a.viv.cn.").SetNonce("123456").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							return sign
						}(),
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiAuthInit(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->err:", err)
		})
	}
}

func TestClient_ApiAuthConf(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiAuthConfRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IDL.CommonResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "TestClient_DwApiAuthConf",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://192.168.10.232:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiAuthConfRequest{
					DataDoi: "25test1_data.viv.cn.",
					Authorization: idl.DataAuthorization{
						Doi: "25test2.viv.cn.",
					},
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "25test1.viv.cn.",
						SignatureNonce: "123456",
						Signature: func() string {
							sign, _ := IDL.NewSignatureData().SetOperator("25test1.viv.cn.").SetNonce("123456").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							return sign
						}(),
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "TestClient_DuApiAuthConf",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://192.168.10.232:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiAuthConfRequest{
					DataDoi: "25test1_data.viv.cn.",
					Authorization: idl.DataAuthorization{
						Doi:  "25test3.viv.cn.",
						Type: idl.UserAuthType,
						Confirmation: func() string {
							sign, err := IDL.NewSignatureData().SetOperator("").SetNonce("sha256").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							fmt.Println("SignByPK-->:", sign, err)
							return sign
						}(),
					},
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "25test3.viv.cn.",
						SignatureNonce: "123456",
						Signature: func() string {
							sign, _ := IDL.NewSignatureData().SetOperator("25test3.viv.cn.").SetNonce("123456").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							return sign
						}(),
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[应用测试用户-代理修改权限]确认授权",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "http://39.107.180.231:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiAuthConfRequest{
					DataDoi: "update_user_a_file_aa.viv.cn.",
					Authorization: idl.DataAuthorization{
						Doi:  "update_user_b.viv.cn.",
						Type: idl.UserAuthType,
						Confirmation: func() string {
							sign, err := IDL.NewSignatureData().SetOperator("").SetNonce("sha256").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							fmt.Println("SignByPK-->:", sign, err)
							return sign
						}(),
					},
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "update_user_b.viv.cn.",
						SignatureNonce: "123456",
						Signature: func() string {
							sign, _ := IDL.NewSignatureData().SetOperator("update_user_b.viv.cn.").SetNonce("123456").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							return sign
						}(),
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiAuthConf(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->err:", err)
		})
	}
}

func TestClient_ApiAuthRev(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiAuthRevRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ApiDisResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Du撤销授权测试用例",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiAuthRevRequest{
					DataDoi: "25test1_data.viv.cn.",
					DuDoi:   "25test2.viv.cn.",
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "25test2.viv.cn.",
						SignatureNonce: "123456",
						Signature: func() string {
							sign, _ := IDL.NewSignatureData().SetOperator("25test2.viv.cn.").SetNonce("123456").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							return sign
						}(),
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "Dw撤销授权测试用例",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiAuthRevRequest{
					DataDoi: "25test1_data.viv.cn.",
					DuDoi:   "25test3.viv.cn.",
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "25test1.viv.cn.",
						SignatureNonce: "123456",
						Signature: func() string {
							sign, _ := IDL.NewSignatureData().SetOperator("25test1.viv.cn.").SetNonce("123456").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							return sign
						}(),
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiAuthRevoke(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->err:", err)
		})
	}
}

func TestClient_ApiTxGet(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiTransactionInfoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ApiDisResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "查询数据区块链交易信息",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiTransactionInfoRequest{
					DataDoi: "25test1_data.viv.cn.",
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "25test2.viv.cn.",
						SignatureNonce: "123456",
						Signature: func() string {
							sign, _ := IDL.NewSignatureData().SetOperator("25test2.viv.cn.").SetNonce("123456").CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
							return sign
						}(),
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiTransactionGet(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->err:", err)
		})
	}
}

func TestClient_ApiDOUpdate(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiDOUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ApiDisResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "[应用测试用户] 更新数岛DAR地址",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOUpdateRequest{
					Doi:           "alice_create_by_lyl_dao.viv.cn.",
					Dar:           "http://alice_create_by_lyl.dao.viv.cn",
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[应用测试用户] 更新图片标识的匿名权限",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOUpdateRequest{
					Doi: "test_pic_pm3.viv.cn.",
					Authorization: &idl.DataAuthorization{
						Doi:  "test_pic_pm3.viv.cn.",
						Type: idl.UserAuthType,
						Description: &idl.PermissionDescription{
							PermissionDoi: "alice_create_by_lyl_default_permission.viv.cn.",
							CreatorDoi:    "alice_create_by_lyl.viv.cn.",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[应用测试用户] 更新专题的匿名权限",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOUpdateRequest{
					Doi: "subject_create_by_lyl.viv.cn.",
					Authorization: &idl.DataAuthorization{
						Doi:  "subject_create_by_lyl.viv.cn.",
						Type: idl.UserAuthType,
						Description: &idl.PermissionDescription{
							PermissionDoi: "alice_create_by_lyl_default_permission.viv.cn.",
							CreatorDoi:    "alice_create_by_lyl.viv.cn.",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[应用测试用户] 更新资讯属性，增加到专题",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOUpdateRequest{
					Doi:           "information_create_by_lyl.viv.cn.",
					DwDoi:         "subject_create_by_lyl2.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[应用测试用户] 更新资讯的继承上级匿名权限",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOUpdateRequest{
					Doi: "information_create_by_lyl.viv.cn.",
					Authorization: &idl.DataAuthorization{
						Doi:  "information_create_by_lyl.viv.cn.",
						Type: idl.UserAuthType,
						Description: &idl.PermissionDescription{
							PermissionDoi: "alice_create_by_lyl_default_subject_article_permission.viv.cn.",
							CreatorDoi:    "alice_create_by_lyl.viv.cn.",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[应用测试用户] 资讯添加到专题，更新资讯的匿名权限",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOUpdateRequest{
					Doi: "information_create_by_lyl2.viv.cn.",
					Authorization: &idl.DataAuthorization{
						Doi:  "information_create_by_lyl2.viv.cn.",
						Type: idl.UserAuthType,
						Description: &idl.PermissionDescription{
							PermissionDoi: "alice_create_by_lyl_default_subject_article_permission.viv.cn.",
							CreatorDoi:    "alice_create_by_lyl.viv.cn.",
							ParentDoi:     "subject_create_by_lyl3.viv.cn.",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[数岛递归读取测试]更新测试数据的匿名权限",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOUpdateRequest{
					Doi: "dao_data_bbb.viv.cn.",
					Authorization: &idl.DataAuthorization{
						Doi:  "dao_data_bbb.viv.cn.",
						Type: idl.UserAuthType,
						Description: &idl.PermissionDescription{
							PermissionDoi: "alice_create_by_lyl_default_subject_article_permission.viv.cn.",
							CreatorDoi:    "dao_bob_by_lyl.viv.cn.",
							ParentDoi:     "dao_data_aaa.viv.cn.",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("dao_bob_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[应用测试用户-代理修改权限]更新测试数据的dar",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "http://39.107.180.231:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOUpdateRequest{
					Doi:           "update_user_a_file_aa.viv.cn.",
					Dar:           "http://www.google.com",
					SignatureData: *IDL.NewSignatureDataWithSign("update_user_c.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiDOUpdate(tt.args.ctx, tt.args.request)
			log.Println("--->test_name", tt.name)
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->err:", err)
		})
	}
}

func TestClient_ApiDOQuery(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiDOQueryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ApiDOQueryResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "[应用测试用户] 查询用户属性",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOQueryRequest{
					Doi: "alice_create_by_lyl.viv.cn.",
					Type: []idl.SearchType{
						idl.ClassGrade,
						idl.Owner,
						idl.PubKey,
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "[应用测试用户] 查询数岛属性",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://39.107.180.231:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOQueryRequest{
					Doi: "update_user_a_file.viv.cn.",
					Type: []idl.SearchType{
						//idl.ClassGrade,
						//idl.Owner,
						//idl.PubKey,
						idl.Dar,
					},
					DirectQuery: true,
				},
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "[Online] 查询数岛",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://registryservice-api.dis.fuxizhiku.org.cn",
				DisQHost: "http://resolverservice-api.dis.fuxizhiku.org.cn",
				DaoHost:  "http://39.107.180.231:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOQueryRequest{
					Doi: "update_user_a_file.viv.cn.",
					Type: []idl.SearchType{
						idl.Dar,
						idl.ClassGrade,
						idl.Owner,
					},
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiDOQuery(tt.args.ctx, tt.args.request)
			log.Println("--->test_name", tt.name)
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->err:", err)
		})
	}
}
func TestClient_ApiRegistrationDataUpdate(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiWhoisUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ApiDisResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "[应用测试用户] 更新whois信息",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiWhoisUpdateRequest{
					WhoisData: &idl.RegistrationData{
						Doi:          "whois.viv.cn.",
						Organization: []string{"organization", "organization2"},
						Contact:      []string{"contact", "contact2"},
						IP:           []string{"ip", "ip2"},
						ASN:          []string{"asn", "asn2"},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("whois.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiRegistrationDataUpdate(tt.args.ctx, tt.args.request)
			log.Println("--->test_name", tt.name)
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->err:", err)
		})

	}
}

func TestClient_ApiGetRegistrationData(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiRegDataRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ApiDisResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "[应用测试用户] 获取whois信息",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://localhost:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiRegDataRequest{
					DataDoi: "whois.viv.cn.",
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiGetRegistrationData(tt.args.ctx, tt.args.request)
			log.Println("--->test_name", tt.name)
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->err:", err)
		})

	}
}

func TestClient_ApiDOCreateforWhoisTest(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiDOCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ApiDisResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "[应用测试用户] 注册用户",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://localhost:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOCreateRequest{
					Doi:    "whois.viv.cn.",
					DwDoi:  "whois.viv.cn.",
					PubKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					WhoisData: &idl.RegistrationData{
						Doi: "whois.viv.cn.",
						Contact: []string{
							"https://segmentfault.com/q/1010000043984824",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("whois.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "[应用测试用户] 注册知库用户",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.177.1:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOCreateRequest{
					Doi:    "zhiku.viv.cn.",
					DwDoi:  "zhiku.viv.cn.",
					PubKey: "30820122300d06092a864886f70d01010105000382010f003082010a028201010099c15ef5017c2c3ee8f9c278f6107e03755b2dc85eeff54e3f5d08cf53a26fe61c05a3dc8b71d7f2929d5c60699e8f9a4f30b12322617eb0682db2658216f1246e3babf70263f457205edbb42f43519cba5ae6e1fff8c4e667538577306a18e1340978dba673efed8f2224f7b9d5310e9aaef3f5095d613c0f42dcfcd7dfb8642b9d5674b0c0e8c0ce33fdfd744e5d593b0ba29ab2d69cd0c2e7ec76746223492755d57c1b3829f4cb5e5f86963f11939ae973c22eacfc193199110ee72345548bf3ce2b92ff9e6396f9b060a8b8c66e9559a64f6066c2f5594a722ca3b4c2205c1f14fef731ef142e1897741abb23308fc733cc966170a68edee435476b71290203010001",
					WhoisData: &idl.RegistrationData{
						Doi: "zhiku.viv.cn.",
						Contact: []string{
							"http://adminapi.fuxizhiku.org.cn/fxkb/platform/dis-agency/callback",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("zhiku.viv.cn.", "308204bf020100300d06092a864886f70d0101010500048204a9308204a5020100028201010099c15ef5017c2c3ee8f9c278f6107e03755b2dc85eeff54e3f5d08cf53a26fe61c05a3dc8b71d7f2929d5c60699e8f9a4f30b12322617eb0682db2658216f1246e3babf70263f457205edbb42f43519cba5ae6e1fff8c4e667538577306a18e1340978dba673efed8f2224f7b9d5310e9aaef3f5095d613c0f42dcfcd7dfb8642b9d5674b0c0e8c0ce33fdfd744e5d593b0ba29ab2d69cd0c2e7ec76746223492755d57c1b3829f4cb5e5f86963f11939ae973c22eacfc193199110ee72345548bf3ce2b92ff9e6396f9b060a8b8c66e9559a64f6066c2f5594a722ca3b4c2205c1f14fef731ef142e1897741abb23308fc733cc966170a68edee435476b712902030100010282010100869db1e8aae32fcde8a67ba8f22f205520d3b2b2e2c40eaef2751ef87e8e1290926a31bc2f6e7e16485d73fa899ea4b602ad6bf989e9784535010557305dabc52caa1bf688e6409063ac0989a973b42997536d1fe3bfc78cbe5c76d712d1617512dd542e637ed04ada6d979638e2ba2fc22879394978b36025307251012443c83c9055a9dc023da6c926c01c74064b884610d4e30b7324979f6aa1be6c82540e930de013bfaced48a00fb90a9ee8fc7ba31f617f7ccb4c5be1f67a7e9d503559bd46948e1cda624cc53e071f3b5f943c31a438a56a5c6a6a36b6f43a0d1d2ba65a4f192ba3329f4d0a97c0bbbc9a72a31b67ac54c7523a57070be39058bd83fd02818100cb5798dc3ba7abc5e87840116e2207caafd948f7ad085987a9f91bd4d93e446d4c069fe842ee0c5d0caba33f219352aae0528b54d0538aa1754ab239d4573709af3a361efc0e18f954cad9f17b6b3be977abc9737598e7eacdd9986177f66aa1713dcef1d9fc0660c29982c4df8c046ec2a7433b9e087553f8e11844e26ab80f02818100c19274c70ac4061c95444600d07bde351d0ede179c58cfdff834c7cd352e5edb658a3d58c65f924973ae0f2d89eb291f6ba33938e5125a1d769d6b19e842e788fa4dbff34700df69bbee54ab27efbd0d7acc8d73069d21c47d189d5c9b7ca0e3fe8ff7ebe8d5cbb85522cf4cf103a4220a9c4c4b43a2ecba49eaebe7b3ba4b4702818069fb7812a18d1cbc8413d8e0bcd443d7629c0fa9f7a7b8723b27395850fa6153ce224dca12c85bed4ba351ec9fa5579af45c517e9d2e4dbd25930f1d910cfc04b22dd6f383501db82677abec6ef54f3eba3ef13a9a7a5db646203989e3aaad9d0396c17bd0afc8eb39c22524539778dff9d88ff44cc3ffd30a8ed7c55f755c05028181008cca6c3b753e1c3fdfe577911212760d65a431af349d781cabd81fd6c6ae8279cb01e01ad8b61c9d66111ca2ffa45615af6159b6630e9512c6fa3a32eeb6f2d6b34fa7a457697015e485b57983a3a07ad46d41187f9ffc3680d24d6a550131b882a7ce27fd02bb98c7fb7891badeee1b80622c2fb5f323815f5009e34ddadaf302818100ad1748dc78df1ede04e23c5ae03cf219b6d105454193c28e25c918e2949fe684adbeab4479a82f445b5c6621d89c475b3f49d004192d1ccd202ca295f4dea30f065d7ee48b449aa3676fdcbff11efa240a81c8693e1e2479dd95d876463635f17f1457436cbf94e7445860f2710d3c7395f46aefd3be2f8aab20ef673f0a438f"),
				},
			},
			want:    nil,
			wantErr: nil,
		},

		{
			name: "[应用测试用户-代理修改权限] 注册用户",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOCreateRequest{
					Doi:    "update_user_c.viv.cn.",
					DwDoi:  "update_user_c.viv.cn.",
					PubKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					WhoisData: &idl.RegistrationData{
						Doi: "update_user_c.viv.cn.",
						Contact: []string{
							"https://segmentfault.com/q/1010000043984824",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("update_user_c.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			got, err := c.ApiDOCreate(tt.args.ctx, tt.args.request)
			log.Println("--->test_name", tt.name)
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->err:", err)
		})
	}
}
