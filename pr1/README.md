# Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

Let's start with the naive solution. It's pretty easy to see that in python this can be written pretty elegantly as

```
sum([x for x in range(1000) if x % 3 == 0 or x % 5 == 0])
```

But, this is a slow implementation. There is a faster way to compute this. As can be seen from the benchmarks below, the slow solution is considerably slower than the faster solution.

## Benchmarks

```
BenchmarkFaster-8       200000000            46.1 ns/op (~325000x faster than slow and ~4000x faster than fast)
BenchmarkFast-8             50000        178992   ns/op (~85x faster)
BenchmarkSlow-8               500      14900300   ns/op
```

## Why slow?

The reason why the above solution is slow, is because of branch prediction. There's a great Stackoverflow post [here](https://stackoverflow.com/questions/11227809/why-is-it-faster-to-process-a-sorted-array-than-an-unsorted-array) that explains the phenomenon in detail. The gist of it is that the conditional `if x % 3 == 0 or x % 5 == 0` is hard to predict. So, the processor has to slow down considerably for the branches to resolve. This is called data-dependant branching. So, for faster performance, we should avoid this conditional.

## Fast solution

Notice that it's easy to compute all the multiples of 3 and all the multiples of 5 that are less than 1000 easily *without* branches.

```
[x * k for x in range(1000 / k)]
```

But, we've counted some numbers twice. The numbers that are *both* a multiple of 3 and a multiple of 5 have been counted twice. So, we need to subtract them.

```
solution = sum([x * 3 for x in range(1000 / 3 + 1)]) + sum([x * 5 for x in range(1000 / 5)]) - sum([x * 15 for x in range(1000 / 15 + 1)])
```

This solution is identical to our first but run significantly faster.

## Is this a slower solution in Big-O if we ignore branch prediction?

The slow solution goes through every number between 1 and n. So it's O(n).
But the fast solution has 3 loops of complexity: O(n/3), O(n/5) and O(n/15) = O(8n/15). So indeed from a Big-O complexity standpoint, the "fast" algorithm is still faster.

## Can we get even faster?

Notice that `sum([x * 3 for x in ...])` is the `sum([3, 6, 9, 12, ...])`. The sequence `[3, 6, 9, 12, ...]` is an arithmetic progression. We know that the sum of an arithmetic progression is `sum(n) = n * (a1 + an) / 2`. So, we can write `sum([3, 6, ..., 999])` as `333 * (999 + 3) / 2 = 166833`. By extension the other two sums can also be represented the same way. So, we get to a better solution by doing:

```
def getIdx(n, d):
    i = n / d
    if n % d == 0:
        i = i - 1 # we want strictly less than n
    return i

""" Sum of arithmetic progression """
def getSum(n, f):
    return n * (f(1) + f(n)) / 2

def faster(x, y, n):
    xi = getIdx(n, x)
    yi = getIdx(n, y)
    qi = getIdx(n, x * y)
    return getSum(xi, lambda i: i * x) + getSum(yi, lambda i: i * y) - getSum(qi, lambda i: i * x * y)
```

## How fast is this?

For this benchmark, I raised the value of n under test from 1 million to 100 million.

```
BenchmarkFaster-8       200000000            49.5 ns/op (~30 million x faster than slow and ~375000x faster than fast)
BenchmarkFast-8               500        18521831 ns/op (~81x faster)
BenchmarkSlow-8                 5      1502128589 ns/op
```

Constant time with respect to n!!
