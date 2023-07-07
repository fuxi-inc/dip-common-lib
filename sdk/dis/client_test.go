package dis

import (
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"
	"os"
	"testing"

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
	sign.OperatorDoi = "alice.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	if err != nil {
		print(err.Error())
	}
	assert.Nil(t, err)
	sign.Signature = Signature
	whois := &idl.RegistrationData{
		Doi:     "XXX.viv.cn",
		Contact: []string{"xxx", "yyy"},
	}
	request := &idl.ApiDOCreateRequest{
		Doi:           "example_alice.viv.cn.",
		DwDoi:         "alice.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
		WhoisData:     whois,
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

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
	sign.OperatorDoi = "bob.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	if err != nil {
		print(err.Error())
	}
	assert.Nil(t, err)
	sign.Signature = Signature
	whois := &idl.RegistrationData{
		Doi:     "XXX.viv.cn",
		Contact: []string{"xxx", "yyy"},
	}
	request := &idl.ApiDOCreateRequest{
		Doi:           "bob.viv.cn.",
		DwDoi:         "bob.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
		WhoisData:     whois,
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

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
	sign.OperatorDoi = "alice.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	if err != nil {
		print(err.Error())
	}
	assert.Nil(t, err)
	sign.Signature = Signature
	whois := &idl.RegistrationData{
		Doi:     "XXX.viv.cn",
		Contact: []string{"xxx", "yyy"},
	}
	request := &idl.ApiDOCreateRequest{
		Doi:           "张三.viv.cn.",
		DwDoi:         "alice.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
		WhoisData:     whois,
		SignatureData: sign,
	}
	client := NewClient().
		InitLogger(zap.NewExample()).
		// TODO: 添加disq的host名称
		InitDis("http://39.107.180.231:8991")

	// 执行被测试的函数
	ctx := &gin.Context{}
	response, err := client.ApiDOCreate(ctx, request)
	print(response.Errmsg)
	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

}
func Test_DOUpdate(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "alice.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(GetPrivKeyString())
	assert.Nil(t, err)
	sign.Signature = Signature

	// 更新数据标识
	request := &idl.ApiDOUpdateRequest{
		Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
		NewDoi: "XXX.viv.cn.",
		DwDoi:  "alice.viv.cn.",

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
	request = &idl.ApiDOUpdateRequest{
		Doi:           "XXX.viv.cn",
		PubKey:        GetPubKeyString(),
		DwDoi:         "usera.viv.cn",
		SignatureData: sign,
	}
	response, err = client.ApiDOUpdate(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
	// 更新数据地址及摘要

	digest := &idl.DataDigest{
		Algorithm: "SHA256",
		Result:    "sha256",
	}

	classgrade := &idl.ClassificationAndGrading{
		Class: 1024,
		Grade: 2048,
	}

	auth := &idl.DataAuthorization{
		Confirmation: "XXX",
	}

	request = &idl.ApiDOUpdateRequest{
		Doi:                      "XXX.viv.cn",
		Dar:                      "resource.example.com/path",
		Digest:                   digest,
		DwDoi:                    "usera.viv.cn",
		Authorization:            auth, // 更新其中的Conformation确权信息
		ClassificationAndGrading: classgrade,
		SignatureData:            sign,
	}
	response, err = client.ApiDOUpdate(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
	// 数据所有者更新自己的权益

	desc := &idl.PermissionDescription{
		PermissionDoi: "XXX.viv.cn",
		CreatorDoi:    "yyy.viv.cn",
	}

	auth = &idl.DataAuthorization{
		Type:        0,
		Description: desc,
	}

	request = &idl.ApiDOUpdateRequest{
		Doi:           "XXX.viv.cn",
		Authorization: auth, // 更新其中的Type和Description
		SignatureData: sign,
	}
	response, err = client.ApiDOUpdate(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)
	// 更新联系方式

	whois := &idl.RegistrationData{
		Doi:     "XXX.viv.cn",
		Contact: []string{"xxx", "yyy"},
	}

	request = &idl.ApiDOUpdateRequest{
		Doi:           "XXX.viv.cn",
		WhoisData:     whois,
		SignatureData: sign,
	}
	response, err = client.ApiDOUpdate(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

}

//更新公钥
func Test_DOUpdate1(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "alice.viv.cn."
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
		Doi:           "example_alice.viv.cn.",
		PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/cindy/public.hex")),
		DwDoi:         "alice.viv.cn.",
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

//dar digest
func Test_DOUpdate2(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "alice.viv.cn."
	sign.SignatureNonce = "123457"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	assert.Nil(t, err)
	sign.Signature = Signature

	digest := &idl.DataDigest{
		Algorithm: "SHA256",
		Result:    "sha256",
	}
	request := &idl.ApiDOUpdateRequest{
		Doi:           "example_alice.viv.cn.",
		Dar:           "resource.example.com/path",
		Digest:        digest,
		DwDoi:         "alice.viv.cn.",
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
	// 更新公钥

	// 断言函数返回的错误为 nil

}

//分类分级
func Test_DOUpdate3(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "alice.viv.cn."
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
		Doi:                      "data.viv.cn.",
		ClassificationAndGrading: classgrade,
		DwDoi:                    "alice.viv.cn.",
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

//更新权属
func Test_DOUpdate4(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "alice.viv.cn."
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
		PermissionDoi: "XXX.viv.cn",
		CreatorDoi:    "yyy.viv.cn",
	}

	auth := &idl.DataAuthorization{
		Doi:         "alice.viv.cn.",
		Type:        0,
		Description: desc,
	}
	request := &idl.ApiDOUpdateRequest{
		Doi:           "data.viv.cn.",
		Authorization: auth,
		DwDoi:         "alice.viv.cn.",
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

func Test_DOUpdate5(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "bob.viv.cn."
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
		Contact: []string{"http://bob.baidu.com", "yyy", "zzz"},
	}
	request := &idl.ApiDOUpdateRequest{
		Doi:           "bob.viv.cn.",
		WhoisData:     whois,
		DwDoi:         "bob.viv.cn.",
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

//new doi
func Test_DOUpdate6(t *testing.T) {

	sign := IDL.SignatureData{}
	sign.OperatorDoi = "alice.viv.cn."
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
		Doi:           "data.viv.cn.",
		NewDoi:        "data2.viv.cn.",
		DwDoi:         "alice.viv.cn.",
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
	sign.OperatorDoi = "bob.viv.cn."
	sign.SignatureNonce = "123456"
	Signature, err := sign.CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
	assert.Nil(t, err)
	sign.Signature = Signature

	// 删除数据标识
	request := &idl.ApiDODeleteRequest{
		Doi: "bob.viv.cn.",

		SignatureData: sign,
	}
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
		Doi: "example_doi",
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
		InitDisQ("xxxxx")

	// 执行被测试的函数
	response, err := client.ApiDOQuery(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

	// TODO：需要补充具体内容，是否需要测试？
	// 创建一个预期的响应数据
	expectedData := &idl.ApiDOQueryResponseData{
		PubKey: "xxxx",
		Owner:  "xxxx",
		Dar:    "xxx",
		Digest: &idl.DataDigest{
			Algorithm: "xxx",
			Result:    "xxx",
		},
		ClassificationAndGrading: &idl.ClassificationAndGrading{
			Class: 0,
			Grade: 0,
		},
	}

	// 判断预期响应结构是否正确
	assert.Equal(t, expectedData.PubKey, response.Data.PubKey)
	assert.Equal(t, expectedData.Owner, response.Data.Owner)
	assert.Equal(t, expectedData.Dar, response.Data.Dar)
	assert.Equal(t, expectedData.Digest, response.Data.Digest)
	assert.Equal(t, expectedData.ClassificationAndGrading, response.Data.ClassificationAndGrading)

}

func Test_DOAuthQuery(t *testing.T) {

	dudoi := ""

	// 设置测试数据
	request := &idl.ApiDOAuthQueryRequest{
		// TODO: 设置测试doi
		Doi:   "example_doi",
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
		InitDisQ("xxxxx")

	// 执行被测试的函数
	response, err := client.ApiDOAuthQuery(ctx, request)

	// 断言函数返回的错误为 nil
	assert.Nil(t, err)

	// 判断 Errno 是否为 0
	assert.Equal(t, IDL.RespCodeType(0), response.Errno)

	// TODO：需要补充具体内容，是否需要测试？
	// 创建一个预期的响应数据DataAuthorization
	expectedData := &idl.DataAuthorization{
		Doi:          "",
		Type:         idl.AuthorizationType(0),
		Confirmation: "",
		Description: &idl.PermissionDescription{
			PermissionDoi: "",
			CreatorDoi:    "",
			Key:           "",
		},
	}

	au := response.Data.Auth[dudoi]

	// 判断预期响应结构是否正确
	assert.Equal(t, expectedData.Doi, au.Doi)
	assert.Equal(t, expectedData.Type, au.Type)
	assert.Equal(t, expectedData.Confirmation, au.Confirmation)
	assert.Equal(t, expectedData.Description, au.Description)

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
				DisHost:  "",
				DisQHost: "",
				DaoHost:  "",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiDOCreateRequest{
					Doi:           "alice_create_by_liuyanglong.viv.cn",
					DwDoi:         "alice_create_by_liuyanglong.viv.cn",
					PubKey:        string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					WhoisData:     nil,
					SignatureData: *IDL.NewSignatureDataWithSign("bob.viv.cn", string(testpkg.GetMockDataContent("/mock_data/user/bob/private.hex"))),
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
			assert.Equalf(t, tt.want, got, "ApiDOCreate(%v, %v, %v)", tt.args.ctx, tt.args.request, err)
		})
	}
}
