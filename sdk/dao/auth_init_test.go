package dao

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"testing"
)

func TestClient_AuthInit(t *testing.T) {
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
		wantErr bool
	}{
		{
			name: "[应用测试] dw alice_create_by_lyl 授权专题给bob",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiAuthInitRequest{
					DataDoi: "subject_create_by_lyl2.viv.cn.",
					Authorization: idl.DataAuthorization{
						Doi:  "bob.viv.cn.",
						Type: idl.UserAuthType,
						Description: &idl.PermissionDescription{
							PermissionDoi: "alice_create_by_lyl_default_permission.viv.cn",
							CreatorDoi:    "alice_create_by_lyl.viv.cn",
							Key:           "",
						},
					},
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
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
			err := c.AuthInit(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("-->err:", err)
			log.Println("-->request:", converter.ToString(tt.args.request))
		})
	}
}
