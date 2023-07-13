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
					DataDoi: "s0mf7w2u23.viv.cn.",
					//SignatureData: *IDL.NewSignatureDataWithSign("ra22abf7n7.viv.cn.", string(testpkg.GetMockDataContent("/mock_data/user/usera/private.hex"))),

					/*
						SignatureData: IDL.SignatureData{
							OperatorDoi:    "ra22abf7n7.viv.cn.",
							SignatureNonce: `jdnynogYuY\\vhiwd`,
							Signature:      "9f7fd042fac820e16aa10961a60f6e2f4cc6d9fdcd78b1340e0d41a02f4d0e447b5f68b94c482c4f99ff21f84b1146b0d17abb55e78d8223a3f940d4b4825cbd7e0edfb6e0562787d9c788b0a97ee63337856dfb4e01096c8bf9064d3bb9869a05e5877ef3567444832f6104d00bd1623ae91f80b0920c0fbfad12f008c18496f00a6288160987cff313916cfcfe201cebd308a3036c1b11c0406a0a4eb20ab3e819735fba143f8fe394253e93c9a6a0b8ec57901b87c260ddef63cc0aaba435b58e41e46aa83f855c31a8ed0444c1e3229903a6f1c5891681b92dd35c69847f204385082318fcdb6970ccb4deedc9aca2562172dbd70f3dd94c90a8f042b002",
						},
					*/
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
