package main

import "fmt"

func fast(x, y, z int) int {
  sum := 0
  prod := x
  for prod < z {
    sum = sum + prod
    prod = prod + x
  }

  prod = y
  for prod < z {
    sum = sum + prod
    prod = prod + y
  }

  prod = x * y
  for prod < z {
    sum = sum - prod
    prod = prod + (x * y)
  }
  return sum
}

func slow(x, y, z int) int {
  sum := 0
  for i := 1; i < z; i++ {
    if i % x == 0 || i % y == 0 {
      sum += i
    }
  }
  return sum
}

func faster(x, y, z int) int {
  xi := getIdx(x, z)
  yi := getIdx(y, z)
  qi := getIdx(x*y, z)

  xf := func(n int) int { return x * n }
  yf := func(n int) int { return y * n }
  qf := func(n int) int { return x * y * n }
  return getSum(xi, xf) + getSum(yi, yf) - getSum(qi, qf)
}

func getSum(n int, f func(j int) int) int {
  return n * (f(1) + f(n)) / 2
}

func getIdx(divisor, num int) int {
  idx := num / divisor
  if num % divisor == 0 {
    idx = idx - 1
  }
  return idx
}

func main() {
  fmt.Println(faster(3, 5, 1000))
  fmt.Println(faster(3, 5, 1000000000))
}
