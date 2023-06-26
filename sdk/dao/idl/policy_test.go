package idl

import (
	"fmt"
	"github.com/imroc/biu"
	"testing"
)

func TestPolicyOperation_ToBinaryString(t *testing.T) {
	tests := []struct {
		name string
		p    PolicyOperation
		want string
	}{
		{
			name: "Case Normal",
			p:    PolicyOperation(123),
			want: "01111011",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ToBinaryString(); got != tt.want {
				t.Errorf("ToBinaryString() = %v, want %v", got, tt.want)
			}
			var b int16
			biu.ReadBinaryString("100000000", &b)
			fmt.Println("ToBinaryString:", b)
		})
	}
}
