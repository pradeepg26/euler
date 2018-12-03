package main

import (
	"fmt"
)

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

func largestPalindrome(n int) (int64, int64, int64) {
	lower := int64(100)
	upper := int64(999)

	// start iteration at nearest multiple of 11
	start := (upper / 11) * 11

	largest := int64(0)
	x := int64(0)
	y := int64(0)
	for i := start; i >= lower; i = i - 11 {
		for j := upper; j >= i; j-- {
			p := i * j
			if isPalindrome(p) && p > largest {
				largest = p
				x = i
				y = j
			} else if p < largest {
				break
			}
		}
	}
	return x, y, largest
}

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

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

type pod struct {
	U int64
	V int64
	Z int64
}

func main() {
	// x, y, z := largestPalindrome(3)
	// fmt.Printf("%d * %d = %d\n", x, y, z)
	uniq := make(map[int64]pod)
	for a := int64(999); a > 99; a = a - 1 {
		// at each step of a, travese the diagonal up and
		// to the right
		// x decrements from a down to 100
		// y increments from a up to 999
		for x, y := a, a; x >= 100 && y <= 999; x, y = x-1, y+1 {
			product := x * y
			u := (x + y) / 2
			v := u - x
			if _, ok := uniq[product]; product%11 == 0 && !ok {
				uniq[product] = pod{u, v, product}
			}
		}
	}
	for _, v := range uniq {
		fmt.Printf("%d %d %d %d %d\n", v.Z, v.U, v.V, (v.U + v.V), (v.U - v.V))
	}
	// for a := int64(999); a > 99; a-- {
	// 	for b := int64(0); b <= 999-a && b <= a; b++ {
	// 		p := a*a - b*b
	// 		fmt.Printf("%d %d %d\n", p, a, b)
	// 	}
	// }
}
