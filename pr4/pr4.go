package main

import (
  "fmt"
  //"os"
  //"strconv"
)

func rev(n int64) int64 {
  num := n
  rev := int64(0)
  for num > 0 {
    dig := num % 10
    rev = rev * 10 + dig
    num = num / 10
  }
  return rev
}

func isPalindrome(n int64) bool {
  return rev(n) == n
}

func largestPalindrome(n int64) (int64, int64, int64) {
  lower := pow(10, n - 1)
  upper := pow(10, n) - 1

  for i := upper; i >= lower; i-- {
    for j := upper; j >= i; j-- {
      p := i * j
      if isPalindrome(p) {
        return i, j, p
      }
    }
  }
  return 0, 0, 0
}

func max(a, b int64) int64 {
  if a > b {
    return a
  } else {
    return b
  }
}

func pow(x, n int64) int64 {
  if n == 0 {
    return 1
  }
  y := int64(1)
  for n > 1 {
    if n % 2 == 0 {
      x = x * x
      n = n / 2
    } else {
      y = y * x
      x = x * x
      n = (n - 1) / 2
    }
  }
  return x * y
}

func primes(n int) <-chan int {
  ch := make(chan int)
  go func() {
    sieve := make([]bool, n)
    for i := 0; i < n; i++ {
      sieve[i] = true
    }
    // mark all even numbers as composite
    for i := range(iterate(4, n, 2)) {
      sieve[i] = false
    }

    // start with the first prime
    i := 2
    for i < n {
      if sieve[i] {
        // this is a prime number
        // send the prime to the channel
        ch <- i

        // mark all the multiples of i as composite
        // optimization: step by 2*i starting at 3*i
        // because we've already marked evens as composite
        for m := range(iterate(i, n, 2 * i)) {
          sieve[m] = false
        }
      }
      i = i + 1
    }
    close(ch)
  }()
  return ch
}

func iterate(start, stop, step int) <-chan int {
  ch := make(chan int)
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

func factorize(n int) <-chan int {
  primes := primes(n)
  factors := make(chan int)
  go func() {
    outer:
    for p := range(primes) {
      if n == 1 {
        break outer
      }
      for n % p == 0 {
        n = n / p
        factors <- p
      }
    }
    close(factors)
  }()
  return factors
}

func main() {
  for i := range iterate(9, -1, -1) {
    fmt.Println(i)
  }
  //i, _ := strconv.Atoi(os.Args[1])
  //for factor := range factorize(i) {
  //  fmt.Printf("%v\n", factor)
  //}
}
