package main

import (
	"fmt"
	"math"
)

func slow(n int64) int64 {
  var sum int64 = 2 // start the sum at 2 since we won't see it in the loop
  var fn2 int64 = 1
  var fn1 int64 = 2
  f := fn1 + fn2
  for f < n {
    if f % 2 == 0 {
      sum = sum + f
    }
    fn2 = fn1
    fn1 = f
    f = fn1 + fn2
  }
  return sum
}

func fast(n int64) int64 {
  var sum int64 = 2 // start the sum at 2
  var fnm1 int64 = 1
  var fn int64 = 2
  for {
    fnp2 := 2 * fn + fnm1
    fnp3 := 3 * fn + 2 * fnm1
    if (fnp3 > n) {
      return sum
    }
    sum = sum + fnp3 // this is the next even number
    fnm1 = fnp2 // set next f(n - 1) to current f(n + 2)
    fn = fnp3 // set next f(n) to current f(n + 3)
  }
}

func mucho(n int64) int64 {
  var sum int64 = 2
  var g0 int64 = 0
  var g1 int64 = 2
  var g int64 = g0 + 4 * g1
  for g < n {
    sum = sum + g
    g0 = g1
    g1 = g
    g = g0 + 4 * g1
  }
  return sum
}

var phi = (1 + math.Sqrt(5)) / 2

func faster(n int64) int64 {
	idx := int64(math.Floor(math.Log(float64(n) * math.Sqrt(5)) / math.Log(phi)))
	idx = idx - (idx % 3) // go down to nearest multiple of 3
	//return (fastfib(idx) + fastfib(idx + 3) - 2) / 4
	return (fib(idx) + fib(idx + 3) - 2) / 4
}

// (define (fib n)
//    (fib-iter 1 0 0 1 n))
// 
// (define (fib-iter a b p q count)
//    (cond ((= count 0) b)
//          ((even? count)
//           (fib-iter a
//                     b
//                     (+ (* p p) (* q q))     ; compute p'
//                     (+ (* 2 p q) (* q q))   ; compute q'
//                     (/ count 2)))
//          (else (fib-iter (+ (* b q) (* a q) (* a p))
//                          (+ (* b p) (* a q))
//                          p
//                          q
//                          (- count 1)))))

func fastfib(n int64) int64 {
	var a, b, p, q, count int64
	a = 1
	b = 0
	p = 0
	q = 1
	count = n
	for count > 0 {
		if count % 2 == 0 {
			pp := p * p + q * q // compute p'
      qp := 2 * p * q + q * q // compute q'
			p = pp // update p to p'
			q = qp // update q to q'
      count = count / 2
		} else {
			ap := b * q + a * q + a * p // compute a'
			bp := b * p + a * q // compute b'
			a = ap // update a to a'
			b = bp // update b to b'
      count = count - 1
		}
  }
	return b
}

func fib(n int64) int64 {
	var a, b int64 = 0, 1
  for b < n {
		next := a + b
		a = b
    b = next
	}
	return b
}

func main() {
  fmt.Println(slow(4000000))
  fmt.Println(fast(4000000))
	fmt.Println(mucho(4000000))
}
