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

		{
			name: "[应用测试] test_pic_pm 读取 test_pic_pm",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.GetDataRequest{
					DuDoi:         "test_pic_pm3.viv.cn.",
					DataDoi:       "test_pic_pm3.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "[应用测试] lyl 读取 subject_create_by_lyl",
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
					DataDoi:       "subject_create_by_lyl.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "[应用测试] subject_create_by_lyl 读取 subject_create_by_lyl",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.GetDataRequest{
					DuDoi:         "subject_create_by_lyl.viv.cn.",
					DataDoi:       "subject_create_by_lyl.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "[应用测试] bob 读取 subject_create_by_lyl",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.GetDataRequest{
					DuDoi:         "bob.viv.cn.",
					DataDoi:       "subject_create_by_lyl.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "[应用测试] bob 读取 subject_create_by_lyl2，无匿名权限",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.GetDataRequest{
					DuDoi:         "bob.viv.cn.",
					DataDoi:       "subject_create_by_lyl2.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("bob.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "[应用测试] bob 读取 information_create_by_lyl2",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.GetDataRequest{
					DuDoi:         "bob.viv.cn.",
					DataDoi:       "information_create_by_lyl2.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("bob.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "手动测试 获取加密文件",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.GetDataRequest{
					DuDoi:   "2igu76pcye.viv.cn.",
					DataDoi: "hzb7ug12cf.viv.cn.",
					SignatureData: IDL.SignatureData{
						OperatorDoi:    "2igu76pcye.viv.cn.",
						SignatureNonce: "Yb\\dvewrsqlor]ih",
						Signature:      "8c552c19c5ec8d7ecf14f46142f775190bafe705129079e1130f513a45b1d48e2d859b2cf21fa9b13779ed3a7fdac3dac3fa9dea7de6bd83e8fdd5a2a61d2c986472fa373e0a62260e3b4379eeeb926d83d8725b4e43c73d03f15236a2a02f42990dfc07a05c7671cf323150d93977351d0057c620c45d39abda9f87db045a941c31398017c29870412111b210fd61cc1848751ac25cb388817e18965d56789bdaaaef566b80a56e5a637750e11f7e03e33165d151cf83c8bdc4f2652840dea187ca84e9532dc51941ca447d51e97081be896ae28693c7e482e190421e5e47dfc5ca03a656d1589c0c4f58430e17a36830f274c76828a3c2570554fb9b213851",
					},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "[应用测试] 获取加密文件",
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
					DataDoi:       "encryption_file4.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("alice_create_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
				},
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "手动测试",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.GetDataRequest{
					DuDoi:   "ra22abf7n7.viv.cn.",
					DataDoi: "f1me52sh39.viv.cn.",
					//SignatureData: *IDL.NewSignatureDataWithSign("ra22abf7n7.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/usera/private.hex"))),

					SignatureData: IDL.SignatureData{
						OperatorDoi:    "ra22abf7n7.viv.cn.",
						SignatureNonce: `rucxuhc806`,
						Signature:      "077bd3f879447a36f92a63f9f07bff5b29b5cf2ca05e545d7e529470cd8080abd7f146f1176235ddc3c7855cb00aef4c89679c66230c0802b1d6cb2bf12d59cd0136ad363154c13a4c3f7162c314b8143d3c94a15ed686264c18d0345bea815df366a42f5d3a44c68e3408d3afe0c5d6cb9b31fbb8987301f6114bf7f15bf92ef107e1a0c1ad9223367bada8fe0f64ac465e8585b936aad16ad9dc2eee13f7653be45f232fd6f394c0da7a20a3b53bd19c0465cf320a320e6b5760177eac7524543e725c0777f73b877baa18eaef8b20d754591aad839934debd4c93f40871bed888f431d45d3fd6776e01259e72edcdc9d80a12a83229307e361d7b5b2580be",
					},
				},
			},
			want:    nil,
			wantErr: false,
		},

		{
			name: "[数岛递归读取测试] 读取数据",
			fields: fields{
				Logger:   zap.NewExample(),
				DisHost:  "http://39.107.180.231:8991",
				DisQHost: "http://39.107.180.231:8053",
				DaoHost:  "http://127.0.0.1:8990",
			},
			args: args{
				ctx: &gin.Context{},
				request: &idl.GetDataRequest{
					DuDoi:         "dao_dale_by_lyl.viv.cn.",
					DataDoi:       "dao_data_ccc.viv.cn.",
					SignatureData: *IDL.NewSignatureDataWithSign("dao_dale_by_lyl.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex"))),
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
