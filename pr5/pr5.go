package main

import (
	"fmt"
	"math"

	"github.com/pradeepg26/euler/primes"
)

func pow(x int64, n int) int64 {
	y := int64(1)
	for n > 0 {
		if n%2 == 0 {
			x = x * x
			n = n / 2
		} else {
			y = y * x
			x = x * x
			n = (n - 1) / 2
		}
	}
	return y
}

func solve(n int64) int64 {
	output := int64(1)
	check := true
	limit := int64(math.Ceil(math.Sqrt(float64(n))))
	for _, p := range primes.PrimesLessThan(n) {
		a := 1
		if check {
			if p < limit {
				a = int(math.Floor(math.Log(float64(n)) / math.Log(float64(p))))
			} else {
				check = false
			}
		}
		output = output * pow(p, a)
	}
	return output
}

func main() {
	fmt.Println(solve(20))
}
