package main

import (
	"testing"
)

// go test to run
func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 100, 12, 58, 97, 64, 12, 4, 6}
	expected := 363
	ch := make(chan int)

	go sum(nums, ch)
	result := <-ch

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// go test -bench=.
func BenchmarkSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 100, 12, 58, 97, 64, 12, 4, 6}
	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		go sum(nums, ch)
		<-ch
	}
}
