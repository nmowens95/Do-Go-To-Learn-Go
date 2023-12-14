package main

import (
	"testing"
)

func BenchmarkCountLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readFile("file1.txt")
	}
}

func BenchmarkCountLines2(b *testing.B) {
	for i := 0; i < b.N; i++ { // N specifies number of iterations
		readFile("file2.txt")
	}
}

// go test -bench=.
