package pr4

import (
	// "fmt"
	"math"
	"testing"
)

func LargestMultipleLessThan(x, n int64) int64 {
	return (n / x) * x
}

func TestLargestMultiple(t *testing.T) {
	// t.Errorf("%d", LargestMultipleLessThan(11, 989))
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if int64(906609) != solution2() {
			b.Error("did not match")
		}
	}
}

func BenchmarkSolution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if int64(906609) != solution3() {
			b.Error("did not match")
		}
	}
}

func BenchmarkSolution4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if int64(906609) != solution4() {
			b.Error("did not match")
		}
	}
}

func IsPerfectSquare(n int64) (bool, int64) {
	intRoot := int64(math.Sqrt(float64(n)))
	return (intRoot * intRoot) == n, intRoot
}

func solution2() int64 {
	res := int64(0)
	for a := int64(999); a >= 100; a = a - 1 {
		var b, db int64
		if a%11 == 0 {
			b = 999
			db = 1
		} else {
			b = 990
			db = 11
		}
		for ; b >= a; b = b - db {
			p := a * b
			if p <= res {
				break
			}
			if isPalindrome(p) {
				res = p
			}
		}
	}
	return res
}

func solution3() int64 {
	lower := int64(100)
	upper := int64(999)
	start := (upper / 11) * 11

	largest := int64(0)
	for i := start; i >= lower; i = i - 11 {
		for j := upper; j >= i; j-- {
			p := i * j
			if isPalindrome(p) && p > largest {
				largest = p
			} else if p < largest {
				// this product is less than the
				// current known largest palindrome
				// can break the inner loop since all
				// subsequent products will be smaller
				break
			}
		}
	}
	return largest
}

func solution4() int64 {
	start := int64((999 * 999 / 11) * 11)
	for c := start; c > int64(10000); c = c - 11 {
		if !isPalindrome(c) {
			continue
		}
		lower := int64(math.Sqrt(float64(c)))
		for j := int64(999); j >= lower; j = j - 1 {
			if c%j == 0 {
				return c
			}
		}
	}
	return int64(-1)
}

func rev(n int64) int64 {
	num := n
	rev := int64(0)
	for num > 0 {
		dig := num % 10
		rev = rev*10 + dig
		num = num / 10
	}
	return rev
}

func isPalindrome(n int64) bool {
	return rev(n) == n
}
