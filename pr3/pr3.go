package main

import (
	"fmt"

	"github.com/pradeepg26/euler/primes"
)

func main() {
	// 3 * 11 * 83 * 331
	// 11 * 82419
	// 11 * (290 + 41)(290 - 41)
	// 90727
	// 7 * 13 * 997
	factors := primes.Factorize(600851475143)
	fmt.Println(factors[len(factors)-1])
}
