package dao

import (
	"github.com/fuxi-inc/dip-common-lib/sdk/dao/idl"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"testing"
)

func TestClient_UpdateContent(t *testing.T) {
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
		// TODO: Add test cases.
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
