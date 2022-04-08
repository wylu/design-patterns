package main

import (
	"testing"
)

// race 竞态检测
// go test -race -bench=. -benchmem -run=none
func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go getInstance()
	}
}
