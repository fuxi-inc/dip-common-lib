package dao

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dao/idl"
	idl2 "github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/fuxi-inc/dip-common-lib/utils/security"
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestClient_UpdateContent(t *testing.T) {
	subjectNewContent := `{"type":"subject","title":"测试专题1_newcontent","describe":"这是一个测试专题_newcontent","content":{"cover_image":"dip://test_pic_pm3.viv.cn","article_list":[]}}`

	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.UpdateDataContentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "[应用测试用户] 更新专题内容",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.UpdateDataContentRequest{
					Doi:     "subject_create_by_lyl.viv.cn.",
					DwDoi:   "alice_create_by_lyl.viv.cn.",
					Content: (subjectNewContent),
					Digest: &idl2.DataDigest{
						Algorithm: "SHA256",
						Result:    base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte(subjectNewContent))),
					},
					Confirmation: func() string {
						sign, err := IDL.NewSignatureData().SetOperator("").SetNonce(base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte(subjectNewContent)))).CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
						fmt.Println("SignByPK-->:", sign, err)
						return sign
					}(),
					SecretKey:     "",
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			wantErr: false,
		},

		{
			name: "[应用测试用户-代理修改权限] 代理更新内容",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.UpdateDataContentRequest{
					Doi:     "update_user_a_file_aa.viv.cn.",
					DwDoi:   "update_user_a.viv.cn.",
					Content: (subjectNewContent),
					Digest: &idl2.DataDigest{
						Algorithm: "SHA256",
						Result:    base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte(subjectNewContent))),
					},
					Confirmation: func() string {
						sign, err := IDL.NewSignatureData().SetOperator("").SetNonce(base64.StdEncoding.EncodeToString(security.Sha256Hash([]byte(subjectNewContent)))).CreateSignature(string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")))
						fmt.Println("SignByPK-->:", sign, err)
						return sign
					}(),
					SecretKey:     "",
					SignatureData: *IDL.NewSignatureDataWithSign("update_user_b.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
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
			err := c.UpdateContent(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("-->err:", err)
			log.Println("-->request:", converter.ToString(tt.args.request))
		})
	}
}
