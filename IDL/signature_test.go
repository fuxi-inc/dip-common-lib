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
				OperatorDoi:    "bob.viv.cn.",
				SignatureNonce: "123456",
				Signature:      "308204be020100300d06092a864886f70d0101010500048204a8308204a40201000282010100ce2e5733fc3f8e8dc25ec500b3fc712b624c862c28244c65a051ccc5d8f814e89b2b90afa8d5b2b49be24a3d17b0dc8303d63406d92dcf33bfada82c6bb05cf644590be6eb5b1441cfc8bd6245617f856bd8dfdfc3bc2e779139c9476c788675c4ff6baf4ddca191c9b71b4c02522545ab86b3c304aa8295d55a16837e55a3fc269285b908b2660a0239f31d14da0da10ed116a0201c3a3dae675e2ad3ceea6d3dc3445018b19a7c170b611a6de2ef8d44d7fff4d259abcbd9e8925fe84ee356b2e569a03f4bb43646c113e58339dbc845223a66dd9e8a4ac9b63f413db135e03adec65a527b35aba7da1e66b09a2e0e7b6bc7717c488aa94268043d4aa432e90203010001028201004bd01a5c912a459eb693e65885d6133ee29c44d1c3c8e5169146c2c7152ad3755c693e8078d26b2da20c2944218bf4e96fe5b9a7b2fff72a8e16aab9fab714d7b0f6153c49da69ae715adcd85faea417493aabc739cfde3e34f64c9ea8b200af194ada2bd1b388ed748091d6f71b817d06bf37264476f621661c041f41baf7dabc7c6fb4c2a636925dfa1291e7ae9a0b7faadf92b48a3fd5231c2e4f026d2f73140774f61c4f8c76259b4d69e75d5bdd19e8441b04615890abe0e6adcefcfad6e6a7d63b557f8bac21a4d9d24753ea94e7ba54e32ff1fb5ae45fb3e7eec6300ec74872cbabd143a7babbdb725e8f8b514e3bfbdeae9465a30dc2edf99b04fc0102818100f51fcac4d306774d5be888bd4a9cdbb4cb1388075705b37856f1b4c9801dd821ae096dec2b9bc85b16a17a85278c4385caca2024d70a71296d884848e1344fa00deb4df01b099354b5150e336c244d97f70e9d59b65981abe8569385691c581044ef8de75d8cfa25c007698cb14c1cf52396d6b01e15eca621bd0161bfa800c102818100d75437d9900b5114c64f8a3e89534b5a9f549b6448b4bee92cb28f7d3cb368ef972d83b978f9efe60aa90e5298407faec6560c7af80e0e457c716df256a482f8f06a6debce0d3647a514b5b6284546b4dbed4732f6a523adf2731bc552051fe101e78cb63216ece101ba1f4b63e31ce1bedd2eda1d8dde757baad29f46ed142902818100e38c7dc935349d6f5cd0828f664232daa4621f36e11bd3bcf4c9305095f41e7d35785f688c1af3654b9edf83d870a705fe78a05a529dc8eaf2593ef118ce5cd471d76be466d4fec5f5cbf400dc74deeb215799ce7c2e9ee79ca9320cf8c46d23bd3abc7a7927b3d77369ee1bc342aae70e00aa8f977e36cf0a26d0af93213c8102818067e69c585f980654b6e81475de7f91a1b5b5f6912a4004cab0a5ac752ac00b4768b42e9665587cb88cb64c51f06356f8f77cf7e2b224de200b4f1455561765974732a3bd9f9c626fa9a9579100d784a9aa6150f4b76cb1cdb26a76a0fcd5ba2f31631053add1c01546670fc8a9f721ad90125c1425249a8e27d12276769bd91902818100a51f45bbd9e1dc0651e68baa393e32485ee744868c78858074b93128637f1f6a59d62a7717e6a7fb5ad190e450081875f8b2ed57bf48031c88bd4618c7fa58101d6fdf24e9a42516615b1996a719211c1bd913afef61540e5424db50273b93233e7904165feea90af59db6c9290d95a5829d07b5ddb809a2514479d4529acdc2",
			},
			args: args{
				pubKey: "30820122300d06092a864886f70d01010105000382010f003082010a0282010100ce2e5733fc3f8e8dc25ec500b3fc712b624c862c28244c65a051ccc5d8f814e89b2b90afa8d5b2b49be24a3d17b0dc8303d63406d92dcf33bfada82c6bb05cf644590be6eb5b1441cfc8bd6245617f856bd8dfdfc3bc2e779139c9476c788675c4ff6baf4ddca191c9b71b4c02522545ab86b3c304aa8295d55a16837e55a3fc269285b908b2660a0239f31d14da0da10ed116a0201c3a3dae675e2ad3ceea6d3dc3445018b19a7c170b611a6de2ef8d44d7fff4d259abcbd9e8925fe84ee356b2e569a03f4bb43646c113e58339dbc845223a66dd9e8a4ac9b63f413db135e03adec65a527b35aba7da1e66b09a2e0e7b6bc7717c488aa94268043d4aa432e90203010001",
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
