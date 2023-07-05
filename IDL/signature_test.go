package IDL

import (
	"github.com/fuxi-inc/dip-common-lib/utils/testpkg"
	"testing"
)

func TestSignatureData_CreateSignature(t *testing.T) {
	type fields struct {
		OperatorDoi    string
		SignatureNonce string
		Signature      string
	}
	type args struct {
		prvKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				OperatorDoi:    "alice.viv.cn",
				SignatureNonce: "asdf-qwer-dfgh-cvbx",
				Signature:      "",
			},
			args: args{
				prvKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/private.hex")),
			},
			want:    "7dc4da3ce903f278e0eed8bcfcb485a3efdab44e6f9b9803abdc9e873d4087bfda3cb04385b80c38a5e28556cd884ab5a1adc6431c091d6c7bb3efbc5df36f26a6f0f846b3aa4d6b29502c77d51863f503906f2241a91631e1c84c45971ec2f316fd3ae2df5b7c07c98fd8a2c374a8c38d9b061b8dc31872baf505b4c10eacddc116131613a4d0d99bf39e87cb73a2ca007a45084d2dd099df7339bcb953bf26aa3323d6e991a90402702d523057f47a8df8b00620d6f5d3832478596e5f832957fff4a1aef753986b3b9897211988f6a102542bd83a8442d2e32af3c6208248270731c87b561510ddde90e78bae6f42adfe24a389c9f8bb0c79077258c10daf",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignatureData{
				OperatorDoi:    tt.fields.OperatorDoi,
				SignatureNonce: tt.fields.SignatureNonce,
				Signature:      tt.fields.Signature,
			}
			got, err := s.CreateSignature(tt.args.prvKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateSignature() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignatureData_VerifySignature(t *testing.T) {
	type fields struct {
		OperatorDoi    string
		SignatureNonce string
		Signature      string
	}
	type args struct {
		pubKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				OperatorDoi:    "alice.viv.cn",
				SignatureNonce: "asdf-qwer-dfgh-cvbx",
				Signature:      "7dc4da3ce903f278e0eed8bcfcb485a3efdab44e6f9b9803abdc9e873d4087bfda3cb04385b80c38a5e28556cd884ab5a1adc6431c091d6c7bb3efbc5df36f26a6f0f846b3aa4d6b29502c77d51863f503906f2241a91631e1c84c45971ec2f316fd3ae2df5b7c07c98fd8a2c374a8c38d9b061b8dc31872baf505b4c10eacddc116131613a4d0d99bf39e87cb73a2ca007a45084d2dd099df7339bcb953bf26aa3323d6e991a90402702d523057f47a8df8b00620d6f5d3832478596e5f832957fff4a1aef753986b3b9897211988f6a102542bd83a8442d2e32af3c6208248270731c87b561510ddde90e78bae6f42adfe24a389c9f8bb0c79077258c10daf",
			},
			args: args{
				pubKey: string(testpkg.GetMockDataContent("/mock_data/user/alice/public.hex")),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SignatureData{
				OperatorDoi:    tt.fields.OperatorDoi,
				SignatureNonce: tt.fields.SignatureNonce,
				Signature:      tt.fields.Signature,
			}
			if err := s.VerifySignature(tt.args.pubKey); (err != nil) != tt.wantErr {
				t.Errorf("VerifySignature() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
