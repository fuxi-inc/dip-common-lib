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
		{
			name: "Test Case 2",
			fields: fields{
				OperatorDoi:    "hvp28o6qwb.viv.cn.",
				SignatureNonce: "qt1xpswedo",
				Signature:      "8323f069a002fe92d801ce1809fe739611f5c5ce2bb51bdd3ba80a204fa6cb1674e398b13c1b759aa0a09dbf34d6240245879e9e7f12db1c8d57f5fb8518e0a29dba341a49c0f4133199a88fd36dad64aae77bae65cd7c7066d17db4c9dc12523d948de218d261dafad149a08bb40fede33e7729cd78bd1490b157699b6499914a6c012f7413e7238d18fb29ce7375607cc4769f6987d3009ec254b76c47150755e0658347e10c7e956ca41d00088c6ba0ea4145a6893c9d8149f9548174983bde4cdffeba3df04f4bf9fd558ff797f588a7ed0b41370d74f307f2f3ddf1361ed47c32ca6a51b7fcc4c81d30b0e79e58c18b569bf7133f976632a4314688ba12",
			},
			args: args{
				pubKey: "30820122300d06092a864886f70d01010105000382010f003082010a0282010100e09720a2ff26ff12d3b93a2d75004d2a060215384b62172390e6c70ebdc244e26490805b9152e2a1817812c2d3afba72e6475dc3965b940a861558c208a963dacb09b225e31f31fd688711caab8577352c644742b1789931d64b23bd01136c55bcc3de753d5cb4b80823e64be005be8fc6bf911f03107ade5adc7b6962d811555b49f8aa4d87f228c06f2b5244894732e924a87f8ed2145f9e12aff0d8a6527806a01cb8c5c783e5a639e1d7a81774ea97f2ab43d2d35db6754fe6e4fae6bc752c901a1a366832090d5bb697ac6f669d22cf7476aec295c330c6ad74515617da9b17eed6e687b8d47a1bc9d1eff268d443ffa7bd0714177df4cc470706a8b1870203010001",
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
