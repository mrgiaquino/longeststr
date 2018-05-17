package main

import (
	"fmt"
	"testing"
)

func BenchmarkLongestStr(b *testing.B) {
	// run the longeststr function b.N times
	fmt.Println(b.N)
	for n := 0; n < b.N; n++ {
		longestStr("ddd", "uriqfhkfjlalguoaui4yte8uoihdddkjfladjfakhgakjvnakdjakldddjldfkajdkshgakjsddddklfjald")
	}
}
