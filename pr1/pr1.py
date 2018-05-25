#!/usr/bin/env python

def getIdx(n, d):
    i = n / d
    if n % d == 0:
        i = i - 1
    return i

def getSum(n, f):
    return n * (f(1) + f(n)) / 2

def faster(x, y, n):
    xi = getIdx(n, x)
    yi = getIdx(n, y)
    qi = getIdx(n, x * y)

    return getSum(xi, lambda i: i * x) + getSum(yi, lambda i: i * y) - getSum(qi, lambda i: i * x * y)

if __name__ == "__main__":
    print(faster(3, 5, 1000))

