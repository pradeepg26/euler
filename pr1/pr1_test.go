package main

import (
  "testing"
)

func TestFast(t *testing.T) {
  sol := fast(3, 5, 10)
  if sol != 23 {
    t.Error("Expected 23, got ", sol)
  }
}

func TestSlow(t *testing.T) {
  sol := slow(3, 5, 10)
  if sol != 23 {
    t.Error("Expected 23, got ", sol)
  }
}

func BenchmarkFaster(b *testing.B) {
  for i := 0; i < b.N; i++ {
    faster(3, 5, 1000000)
  }
}

func BenchmarkFast(b *testing.B) {
  for i := 0; i < b.N; i++ {
    fast(3, 5, 1000000)
  }
}

func BenchmarkSlow(b *testing.B) {
  for i := 0; i < b.N; i++ {
    slow(3, 5, 1000000)
  }
}
