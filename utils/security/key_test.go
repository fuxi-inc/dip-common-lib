package security

import "testing"

func TestHmacSHA1(t *testing.T) {
	type args struct {
		data string
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				data: "POST" + "\n" + "0B9BE351E56C90FED853B32524253E8B" + "\n" + "application/json" + "\n" + "Tue, 11 Dec 2018 21:05:51 +0800" + "\n" + "x-cms-api-version:1.0" + "\n" + "x-cms-ip:127.0.0.1" + "\n" + "x-cms-signature:hmac-sha1" + "\n" + "/metric/custom/upload",
				key:  "testsecret",
			},
			want: "1DC19ED63F755ACDE203614C8A1157EB1097E922",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HmacSHA1(tt.args.data, tt.args.key); got != tt.want {
				t.Errorf("HmacSHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}
