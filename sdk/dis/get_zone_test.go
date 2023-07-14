package dis

import (
	"fmt"
	"testing"

	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestClient_getZone(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		zone string
		sk   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *idl.ZoneResponse
		wantErr assert.ErrorAssertionFunc
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
			got, err := c.getZone(tt.args.zone)
			if !tt.wantErr(t, err, fmt.Sprintf("getZone(%v, %v)", tt.args.zone, tt.args.sk)) {
				return
			}
			assert.Equalf(t, tt.want, got, "getZone(%v, %v)", tt.args.zone, tt.args.sk)
		})
	}
}
