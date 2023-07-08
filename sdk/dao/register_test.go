package dao

import (
	"encoding/base64"
	"fmt"
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dao/idl"
	idl2 "github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/fuxi-inc/dip-common-lib/utils/security"
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"testing"
	"time"
)

func TestClient_Register(t *testing.T) {
	defaultPermission := &idl.Permission{
		Operations:   4,
		AlgorithmDOI: "",
		Weight:       100,
		StartAt:      IDL.NowTime(),
		ExpiredAt:    IDL.NewTime(IDL.NowTime().Add(time.Hour * 24 * 365 * 10)),
	}

	defaultSubjectArticlePermission := &idl.Permission{
		Operations:   256,
		AlgorithmDOI: "",
		Weight:       100,
		StartAt:      IDL.NowTime(),
		ExpiredAt:    IDL.NewTime(IDL.NowTime().Add(time.Hour * 24 * 365 * 10)),
	}
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.RegisterDataRequest
		permObj *idl.Permission
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "[应用测试用户] 构建默认数据读取权限",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx:     &gin.Context{},
				permObj: defaultPermission,
				request: &idl.RegisterDataRequest{
					Doi:       "alice_create_by_lyl_default_permission.viv.cn.",
					DwDoi:     "alice_create_by_lyl.viv.cn.",
					PublicKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					Content:   []byte(defaultPermission.ToString()),
					FilePath:  "/permission/default_data_permission.data",
					Digest: &idl2.DataDigest{
						Algorithm: "SHA256",
						Result:    base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte(defaultPermission.ToString()))),
					},
					Confirmation: func() string {
						sign, err := IDL.NewSignatureData().SetOperator("").SetNonce(base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte(defaultPermission.ToString())))).CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
						fmt.Println("SignByPK-->:", sign, err)
						return sign
					}(),
					SecretKey: "",
					ClassificationAndGrading: &idl2.ClassificationAndGrading{
						Class: 4097,
						Grade: 0,
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			wantErr: false,
		},

		{
			name: "[应用测试用户] 构建默认专题文章的读取权限",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx:     &gin.Context{},
				permObj: defaultSubjectArticlePermission,
				request: &idl.RegisterDataRequest{
					Doi:       "alice_create_by_lyl_default_subject_article_permission.viv.cn.",
					DwDoi:     "alice_create_by_lyl.viv.cn.",
					PublicKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					Content:   []byte(defaultSubjectArticlePermission.ToString()),
					FilePath:  "/permission/default_subject_article_permission.data",
					Digest: &idl2.DataDigest{
						Algorithm: "SHA256",
						Result:    base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte(defaultSubjectArticlePermission.ToString()))),
					},
					Confirmation: func() string {
						sign, err := IDL.NewSignatureData().SetOperator("").SetNonce(base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte(defaultSubjectArticlePermission.ToString())))).CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
						fmt.Println("SignByPK-->:", sign, err)
						return sign
					}(),
					SecretKey: "",
					ClassificationAndGrading: &idl2.ClassificationAndGrading{
						Class: 4097,
						Grade: 0,
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			wantErr: false,
		},

		{
			name: "[应用测试用户] 创建封面图片",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx:     &gin.Context{},
				permObj: defaultSubjectArticlePermission,
				request: &idl.RegisterDataRequest{
					Doi:       "test_pic.viv.cn.",
					DwDoi:     "alice_create_by_lyl.viv.cn.",
					PublicKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					Content:   testpkg.GetMockDataContent("/mock_data/data/test_pic.jpeg"),
					FilePath:  "/picture/test_pic.jpeg",
					Digest: &idl2.DataDigest{
						Algorithm: "SHA256",
						Result:    base64.StdEncoding.EncodeToString(security.Sha256Hash(testpkg.GetMockDataContent("/mock_data/data/test_pic.jpeg"))),
					},
					Confirmation: func() string {
						sign, err := IDL.NewSignatureData().SetOperator("").SetNonce(base64.StdEncoding.EncodeToString(security.Sha256Hash(testpkg.GetMockDataContent("/mock_data/data/test_pic.jpeg")))).CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
						fmt.Println("SignByPK-->:", sign, err)
						return sign
					}(),
					SecretKey: "",
					ClassificationAndGrading: &idl2.ClassificationAndGrading{
						Class: 0,
						Grade: 0,
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			wantErr: false,
		},

		{
			name: "[应用测试用户] 创建封面图片无匿名权限",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx:     &gin.Context{},
				permObj: defaultSubjectArticlePermission,
				request: &idl.RegisterDataRequest{
					Doi:       "test_pic_no_permission.viv.cn.",
					DwDoi:     "alice_create_by_lyl.viv.cn.",
					PublicKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
					Content:   testpkg.GetMockDataContent("/mock_data/data/test_pic.jpeg"),
					FilePath:  "/picture/test_pic_no_permission.jpeg",
					Digest: &idl2.DataDigest{
						Algorithm: "SHA256",
						Result:    base64.StdEncoding.EncodeToString(security.Sha256Hash(testpkg.GetMockDataContent("/mock_data/data/test_pic.jpeg"))),
					},
					Confirmation: func() string {
						sign, err := IDL.NewSignatureData().SetOperator("").SetNonce(base64.StdEncoding.EncodeToString(security.Sha256Hash(testpkg.GetMockDataContent("/mock_data/data/test_pic.jpeg")))).CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
						fmt.Println("SignByPK-->:", sign, err)
						return sign
					}(),
					SecretKey: "",
					ClassificationAndGrading: &idl2.ClassificationAndGrading{
						Class: 0,
						Grade: 0,
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			wantErr: false,
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
			err := c.Register(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("--->register_content:", tt.args.permObj.ToString())
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->err:", err)
		})
	}
}
