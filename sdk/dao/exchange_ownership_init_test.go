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

func TestClient_ExchangeOwnershipInit(t *testing.T) {
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
			name: "[应用测试] dw alice_create_by_lyl 授权专题给bob",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://39.107.180.231:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.ApiExchangeOwnershipRequest{
					DataDoi:  "ox1a5bq4v7.viv.cn.",
					DwDoi:    "et6vbyylo2.viv.cn.",
					DuDoi:    "lob22xqwup.viv.cn.",
					FilePath: "path",
					Fields: map[string]string{
						"testkey1": "testkeya",
						"testkey2": "testkeyb",
					},
					Authorization: idl.DataAuthorization{
						Doi:  "lob22xqwup.viv.cn.",
						Type: idl.OwnerAuthType,
						Description: &idl.PermissionDescription{
							Key: "",
						},
					},
					SignatureData: *IDL.NewSignatureDataWithSign("et6vbyylo2.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
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

			err := c.ExchangeOwnershipInit(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("-->err:", err)
			log.Println("-->request:", converter.ToString(tt.args.request))
		})
	}
}
