package main

import (
	"testing"
)

func BenchmarkStringCompair(b *testing.B) {
	a := ""
	for i := 0; i < b.N; i++ {
		if a == "" {

		}
	}
}

func BenchmarkStringLengthCompair(b *testing.B) {
	a := ""
	for i := 0; i < b.N; i++ {
		if len(a) == 0 {

		}
	}
}
