package main

import (
	"testing"
)

func BenchmarkSlow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slow(400000000)
	}
}

func BenchmarkFast(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fast(400000000)
	}
}

func BenchmarkMucho(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mucho(400000000)
	}
}

//func BenchmarkFaster(b *testing.B) {
//  for n := 0; n < b.N; n++ {
//    faster(400000000)
//  }
//}
//
//func BenchmarkFib(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		fib(1300000)
//  }
//}
//
//func BenchmarkFastfib(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		fastfib(1300000)
//  }
//}
