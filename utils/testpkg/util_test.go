package testpkg

import "testing"

func TestBase64DecodeFile(t *testing.T) {
	type args struct {
		sourceFile string
		destFile   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				sourceFile: "./mock_data/data/ISCDTD20230906402800.dipx",
				destFile:   "./mock_data/data/ISCDTD20230906402800.png",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Base64DecodeFile(tt.args.sourceFile, tt.args.destFile); (err != nil) != tt.wantErr {
				t.Errorf("Base64DecodeFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
