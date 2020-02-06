package server

import (
	"github.com/wedojava/myencrypt"
	"testing"
)

const n = 16

func BenchmarkBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myencrypt.RandStringBytesMaskImprSrc(n)
	}
}
