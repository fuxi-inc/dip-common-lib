package dis

import (
	"testing"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// func Test_DOCreate(t *testing.T) {

// 	request := &idl.ApiDOCreateRequest{
// 		Doi:       "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
// 		DwDoi:     "usera.viv.cn",
// 		PubKey:    "XXX",
// 		WhoisData: nil,
// 		Sign:      "XXX",
// 	}

// 	println(request)

// }

// func Test_DOUpdate(t *testing.T) {

// 	// 所有未赋值的字段均为空

// 	// 更新数据标识
// 	request := &idl.ApiDOUpdateRequest{
// 		Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
// 		NewDoi: "XXX.viv.cn",
// 		DwDoi:  "usera.viv.cn",
// 		Sign:   "XXX",
// 	}

// 	// 更新公钥
// 	request = &idl.ApiDOUpdateRequest{
// 		Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
// 		PubKey: "XXX",
// 		DwDoi:  "usera.viv.cn",
// 		Sign:   "XXX",
// 	}

// 	// 更新数据地址及摘要

// 	digest := &idl.DataDigest{
// 		Algorithm: "SHA256",
// 		Result:    "XXX",
// 	}

// 	classgrade := &idl.ClassificationAndGrading{
// 		Class: 1024,
// 		Grade: 2048,
// 	}

// 	auth := &idl.DataAuthorization{
// 		Confirmation: "XXX",
// 	}

// 	request = &idl.ApiDOUpdateRequest{
// 		Doi:                      "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
// 		Dar:                      "resource.example.com/path",
// 		Digest:                   digest,
// 		DwDoi:                    "usera.viv.cn",
// 		Authorization:            &[]idl.DataAuthorization{*auth}, // 更新其中的Conformation确权信息
// 		ClassificationAndGrading: classgrade,
// 		Sign:                     "XXX",
// 	}

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
// 		Doi:           "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
// 		Authorization: &[]idl.DataAuthorization{*auth}, // 更新其中的Type和Description
// 		Sign:          "XXX",
// 	}

// 	// 更新联系方式

// 	whois := &idl.RegistrationData{
// 		Doi:     "xxx.viv.cn",
// 		Contact: []string{"xxx", "yyy"},
// 	}

// 	request = &idl.ApiDOUpdateRequest{
// 		Doi:       "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
// 		WhoisData: whois,
// 		Sign:      "XXX",
// 	}

// 	println(request)

// }

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
