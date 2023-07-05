package IDL

import (
	"io/ioutil"
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
				prvKey: func() string {
					data, _ := ioutil.ReadFile("./test_data/slashes_request.json")
					return string(data)
				}(),
			},
			want:    "",
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
