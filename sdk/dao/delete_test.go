package dao

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/sdk/dao/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestClient_Delete(t *testing.T) {

	defaultPermission := &idl.Permission{
		Operations:   4,
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
		request *idl.DeleteDataRequest
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
				request: &idl.DeleteDataRequest{

					Doi: "0720test2_data.viv.cn.",

					SignatureData: *IDL.NewSignatureDataWithSign("0720test2.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			wantErr: false,
		},
	}
	fmt.Println("1")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:   tt.fields.Logger,
				DisHost:  tt.fields.DisHost,
				DisQHost: tt.fields.DisQHost,
				DaoHost:  tt.fields.DaoHost,
			}
			fmt.Println("1")
			err := c.Delete(tt.args.ctx, tt.args.request)
			log.Println("--->test_name:", tt.name)
			log.Println("-->err:", err)
			log.Println("--->register_content:", tt.args.permObj.ToString())
			log.Println("-->request:", converter.ToString(tt.args.request))
		})
	}
}
