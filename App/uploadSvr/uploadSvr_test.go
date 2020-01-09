package main

import (
	"github.com/wedojava/MyTools"
	"testing"
)

const n = 16

func BenchmarkBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyTools.RandStringBytesMaskImprSrc(n)
	}
}
