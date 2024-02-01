package dirpc

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"testing"
)

func TestHttpClient_Get(t *testing.T) {
	type fields struct {
		Logger *zap.Logger
	}
	type args struct {
		ctx         *gin.Context
		uri         string
		queryString []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test Get",
			fields: fields{
				Logger: zap.NewExample(),
			},
			args: args{
				ctx:         &gin.Context{},
				uri:         "https://mbd.baidu.com/newspage/data/landingsuper",
				queryString: []byte("context=%7B%22nid%22%3A%22news_10177135762036746704%22%7D&n_type=1&p_from=3"),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HttpClient{
				Logger: tt.fields.Logger,
			}
			got, err := c.Get(tt.args.ctx, tt.args.uri, tt.args.queryString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println(string(got))

		})
	}
}
