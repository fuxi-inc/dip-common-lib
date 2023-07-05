package dis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestClient_getAllZones(t *testing.T) {
	type fields struct {
		Logger   *zap.Logger
		DisHost  string
		DisQHost string
		DaoHost  string
	}
	type args struct {
		sk string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
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
			got, err := c.getAllZones(tt.args.sk)
			if !tt.wantErr(t, err, fmt.Sprintf("getAllZones(%v)", tt.args.sk)) {
				return
			}
			assert.Equalf(t, tt.want, got, "getAllZones(%v)", tt.args.sk)
		})
	}
}
