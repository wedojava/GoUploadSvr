package main

import "testing"

const n = 16

func BenchmarkBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrc(n)
	}
}

func BenchmarkRandomBase16String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomBase16String(n)
	}
}