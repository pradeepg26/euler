package primes

import (
	"testing"
)

func TestPrimesLessThan(t *testing.T) {
	actual := PrimesLessThan(15)
	exp := []int64{2, 3, 5, 7, 11, 13}

	assertEqualsInt64s(t, exp, actual)
}

func TestFirstNPrimes(t *testing.T) {
	actual := FirstNPrimes(25)
	exp := []int64{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41,
		43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
	}

	assertEqualsInt64s(t, exp, actual)
}

func TestNthPrime(t *testing.T) {
	actual := NthPrime(25)

	if actual != 97 {
		t.Fatalf("Unexpected prime %d", actual)
	}
}

func TestLargestPrimeLessThan(t *testing.T) {
	actual := LargestPrimeLessThan(100)

	if actual != 97 {
		t.Fatalf("Unexpected prime %d", actual)
	}
}

func TestFactorize(t *testing.T) {
	factors := Factorize(30)
	exp := []int64{2, 3, 5}
	assertEqualsInt64s(t, exp, factors)

	factors = Factorize(64)
	exp = []int64{2, 2, 2, 2, 2, 2}
	assertEqualsInt64s(t, exp, factors)
}

func assertEqualsInt64s(t *testing.T, expected, actual []int64) {
	if len(actual) != len(expected) {
		t.Fatalf("Incorrect number of elements. Expected %d, Actual %d", len(expected), len(actual))
	}
	for i, p := range expected {
		if p != actual[i] {
			t.Fatalf("Unexpected value @ idx(%d) Expected = %d, Actual = %d", i, p, actual[i])
		}
	}
}
