package main

import "testing"

func BenchmarkRunLab(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunLab()
	}
}
