package dao

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dao/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"testing"
)

func TestClient_GetData(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		ctx     *gin.Context
		request *idl.GetDataRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.GetDataResponse
		wantErr bool
	}{
		{
			name: "[应用测试] alice_create_by_lyl 读取test_pic",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.GetDataRequest{
					DuDoi:         "alice_create_by_lyl.viv.cn.",
					DataDoi:       "test_pic_pm.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
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
			got, err := c.GetData(tt.args.ctx, tt.args.request)
			log.Println("--->test_name", tt.name)
			log.Println("-->request:", converter.ToString(tt.args.request))
			log.Println("-->response:", converter.ToString(got))
			log.Println("-->err:", err)
		})
	}
}
