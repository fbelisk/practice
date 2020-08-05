package ce

import (
	"fmt"
	"testing"
)

func TestProf(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func BenchmarkSprintf(b *testing.B){
	num:=10
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		fmt.Sprintf("%d",num)
	}
}