package dis

import (
	"testing"

	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
)

func Test_DOUpdate(t *testing.T) {

	// 测试更新DAR
	description := &idl.ApiDescription{
		Key: "XXXX", //	解密密钥
	}

	request := &idl.ApiDOUpdateRequest{
		Doi:         "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
		Dar:         "resource.example.com/path",
		Description: description,
		Sign:        "XXX",
	}

	println(request)

}
