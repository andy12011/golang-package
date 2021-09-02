package utils

import "testing"

const (
	BENCHMARK_TIMES = 1000000
)
// command line : go test -v -bench=. ./utils/ -run none -benchmem -count=5
// ns/op     每一個執行(op)跑多少納秒(ns)
// allocs/op 表示每個操作發生了多少不同的內存分配。
// B/op      是每個操作分配了多少字節。

func BenchmarkStringConnectByBuffer(b *testing.B) {
	for i := 0; i < BENCHMARK_TIMES; i++ {
		StringConnectByBuffer("")
		StringConnectByBuffer("~", "1")
		StringConnectByBuffer("~", "1", "2")
		StringConnectByBuffer("~", "abel test", "3")
	}
	// goos: windows
	// goarch: amd64
	// pkg: golang-package/utils
	// BenchmarkStringConnectByBuffer-4        1000000000               0.0310 ns/op          0 B/op          0 allocs/op
	// BenchmarkStringConnectByBuilder-4       1000000000               0.0170 ns/op          0 B/op          0 allocs/op
	// ok      golang-package/utils    0.568s
}

func BenchmarkStringConnectByBuilder(b *testing.B) {
	for i := 0; i < BENCHMARK_TIMES; i++ {
		StringConnectByBuilder("")
		StringConnectByBuilder("~", "1")
		StringConnectByBuilder("~", "1", "2")
		StringConnectByBuilder("~", "abel test", "3")
	}
}
