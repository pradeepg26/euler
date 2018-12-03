package primes

import (
	"math"
	// "sort"
)

type Primes []int64

func ApproximatePrimeCount(n int64) int {
	return int(math.Floor(float64(n) / (math.Log(float64(n)) - 1)))
}

func PrimesLessThan(n int64) []int64 {
	ch := primes(n)
	sl := make([]int64, 0, ApproximatePrimeCount(n))

	for p := range ch {
		sl = append(sl, p)
	}
	return sl
}

func LargestPrimeLessThan(n int64) int64 {
	v := PrimesLessThan(n)
	return v[len(v)-1]
}

func NthPrime(n int) int64 {
	return FirstNPrimes(n)[n-1]
}

func FirstNPrimes(n int) []int64 {
	upper := int64(100)
	if n > 6 {
		// Upper bound for nth prime is n(ln n + ln ln n)
		upper = int64(n) * int64(math.Log(float64(n))+math.Log(math.Log(float64(n))))
	}
	return PrimesLessThan(upper)[:n]
}

func primes(n int64) <-chan int64 {
	ch := make(chan int64)
	go func() {
		sieve := make([]bool, n)
		for i := int64(0); i < n; i++ {
			sieve[i] = true
		}
		// mark all even numbers as composite
		for i := range iterate(4, n, 2) {
			sieve[i] = false
		}

		// start with the first prime
		i := int64(2)
		for i < n {
			if sieve[i] {
				// this is a prime number
				// send the prime to the channel
				ch <- i

				// mark all the multiples of i as composite
				// optimization: step by 2*i starting at 3*i
				// because we've already marked evens as composite
				for m := range iterate(i, n, 2*i) {
					sieve[m] = false
				}
			}
			i = i + 1
		}
		close(ch)
	}()
	return ch
}

func iterate(start, stop, step int64) <-chan int64 {
	if step == 0 {
		panic("cannot step by 0")
	}
	ch := make(chan int64)
	go func() {
		x := start
		if step > 0 {
			for x < stop {
				ch <- x
				x = x + step
			}
		} else {
			for x > stop {
				ch <- x
				x = x + step
			}
		}
		close(ch)
	}()
	return ch
}

func Factorize(n int64) []int64 {
	plt := PrimesLessThan(n)
	factors := make([]int64, 0)
	for _, p := range plt {
		if n == 1 {
			break
		}
		for n%p == 0 {
			n = n / p
			factors = append(factors, p)
		}
	}
	return factors
}
