package dao

import (
	"github.com/fuxi-inc/dip-common-lib/sdk/dao/idl"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"testing"
)

func TestClient_UpdateContent(t *testing.T) {
	//subjectNewContent := `{"type":"subject","title":"测试专题1_newcontent","describe":"这是一个测试专题_newcontent","content":{"cover_image":"dip://test_pic_pm3.viv.cn","article_list":[]}}`

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
				DisQHost: "",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args:    args{},
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
			if err := c.UpdateContent(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("UpdateContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
