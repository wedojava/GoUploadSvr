package server

import (
	"github.com/wedojava/mytools"
	"testing"
)

const n = 16

func BenchmarkBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mytools.RandStringBytesMaskImprSrc(n)
	}
}
