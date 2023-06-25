package dis

import (
	"testing"

	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
)

func Test_DOCreate(t *testing.T) {

	request := &idl.ApiDOCreateRequest{
		Doi:       "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
		DwDoi:     "usera.viv.cn",
		PubKey:    "XXX",
		Type:      0, // 0表示所有者，1表示使用者
		WhoisData: nil,
		Sign:      "XXX",
	}

	println(request)

}

func Test_DOUpdate(t *testing.T) {

	// 所有未赋值的字段均为空

	// 更新数据标识
	request := &idl.ApiDOUpdateRequest{
		Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
		NewDoi: "XXX.viv.cn",
		DwDoi:  "usera.viv.cn",
		Sign:   "XXX",
	}

	// 更新公钥
	request = &idl.ApiDOUpdateRequest{
		Doi:    "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
		PubKey: "XXX",
		DwDoi:  "usera.viv.cn",
		Sign:   "XXX",
	}

	// 更新数据地址及摘要

	digest := &idl.ApiDigest{
		Algorithm: "SHA256",
		Result:    "XXX",
	}

	classgrade := &idl.ApiClassificationAndGrading{
		Class: 1024,
		Grade: 2048,
	}

	auth := &idl.ApiAuthorization{
		Confirmation: "XXX",
	}

	request = &idl.ApiDOUpdateRequest{
		Doi:                      "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
		Dar:                      "resource.example.com/path",
		Digest:                   digest,
		DwDoi:                    "usera.viv.cn",
		Authorization:            &[]idl.ApiAuthorization{*auth}, // 更新其中的Conformation确权信息
		ClassificationAndGrading: classgrade,
		Sign:                     "XXX",
	}

	// 数据所有者更新自己的权益

	desc := &idl.ApiDescription{
		PermissionDoi: "XXX.viv.cn",
		CreatorDoi:    "yyy.viv.cn",
	}

	auth = &idl.ApiAuthorization{
		Type:        0,
		Description: desc,
	}

	request = &idl.ApiDOUpdateRequest{
		Doi:           "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
		Authorization: &[]idl.ApiAuthorization{*auth}, // 更新其中的Type和Description
		Sign:          "XXX",
	}

	// 更新联系方式

	whois := &idl.ApiWhoisData{
		Doi:     "xxx.viv.cn",
		Contact: []string{"xxx", "yyy"},
	}

	request = &idl.ApiDOUpdateRequest{
		Doi:       "2d8a2384-a705-4c16-a927-1a1b16345b67.viv.cn",
		WhoisData: whois,
		Sign:      "XXX",
	}

	println(request)

}
