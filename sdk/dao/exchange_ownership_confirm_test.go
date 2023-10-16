package dao

import (
	"fmt"
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"testing"
)

func TestClient_ExchangeOwnershipConfirm(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.ApiExchangeOwnershipRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "[应用测试] du bob 确认 alice_create_by_lyl 的授权",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiExchangeOwnershipRequest{
					DataDoi: "subject_create_by_lyl3.viv.cn.",
					Authorization: idl.DataAuthorization{
						Doi:  "bob.viv.cn.",
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
					SignatureData: *IDL.NewSignatureDataWithSign("bob.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
		},

		{
			name: "[数岛递归读取测试] dw 确认dale的请求授权 data—a",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiExchangeOwnershipRequest{
					DataDoi: "dao_data_aaa.viv.cn.",
					Authorization: idl.DataAuthorization{
						Doi:  "dao_dale_by_lyl.viv.cn.",
						Type: idl.UserAuthType,
					},
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: *IDL.NewSignatureDataWithSign("dao_alice_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
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
			err := c.ExchangeOwnershipConfirm(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("-->err:", err)
			log.Println("-->request:", converter.ToString(tt.args.request))
		})
	}
}
