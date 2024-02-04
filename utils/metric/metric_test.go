package metric

import (
	"fmt"
	"testing"
	"time"

	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"go.uber.org/zap"
)

func TestClient_Push(t *testing.T) {
	type fields struct {
		Logger       *zap.Logger
		Host         string
		AccessKey    string
		AccessSecret string
	}
	type args struct {
		m []Metric
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test push",
			fields: fields{
				Logger:       nil,
				Host:         "https://metrichub-cms-cn-beijing.aliyuncs.com",
				AccessKey:    "LTAI5t7mFktJw3yAp2NowLEj",
				AccessSecret: "IYvhhBZOs5tOnTHcv4mlQwq8Iqmayh",
			},
			args: args{
				m: []Metric{
					{
						Dimensions: map[string]string{
							"tagC": "tagD",
						},
						GroupId:    236442042,
						MetricName: "lyl-test-metric",
						Period:     60,
						Time:       time.Now().UnixMilli(),
						Type:       0,
						Values:     &MValue{Value: 10},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Logger:       tt.fields.Logger,
				Host:         tt.fields.Host,
				AccessKey:    tt.fields.AccessKey,
				AccessSecret: tt.fields.AccessSecret,
			}
			fmt.Println("----->", converter.ToString(tt.args.m))
			if err := c.Push(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
