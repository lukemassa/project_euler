package main

import (
	"math"
	"math/big"

	"golang.org/x/exp/constraints"
)

func genPrimes(nums chan int64) {
	nums <- 2
	nums <- 3
	for i := int64(6); ; i += 6 {
		if isPrime(i - 1) {
			nums <- i - 1
		}
		if isPrime(i + 1) {
			nums <- i + 1
		}
	}
}

func factor(n int64) []int64 {
	factors := make([]int64, 0)
	for {
		factor := nextFactor(n)
		if factor == -1 {
			factors = append(factors, n)
			return factors
		}
		factors = append(factors, factor)
		n = n / factor
	}
}

func divisors[T constraints.Integer](n T) []T {
	divisors := make([]T, 0)
	for i := T(1); i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}

	}
	return divisors
}

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	factor := nextFactor(n)
	return factor == -1
}
func fasterIsPrime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}

func nextFactor(n int64) int64 {
	//fmt.Printf("Working on %d\n", n)
	max := int64(math.Sqrt(float64(n)))
	for i := int64(2); i <= max; i++ {
		if n%i == 0 {
			//fmt.Printf("   Found %d\n", i)
			return i
		}
	}
	return -1
}

func numDivisors(n int64) int {
	// Special case; 1 doesn't have any prime factors so the below calculation won't work
	if n == 1 {
		return 1
	}
	ret := 1
	factors := factor(n)
	for _, freq := range getFreq(factors) {
		// Each prime contributes 0, 1, ..., p to a given divisor
		ret *= freq + 1
	}
	return ret
}

func sumDivisors[T constraints.Integer](n T) T {
	return sum(divisors(n))
}
func sum[T constraints.Integer](arr []T) T {
	// Special case; 1 doesn't have any prime factors so the below calculation won't work
	var sum T
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func numPrimesPolynomial(a, b int64) int64 {
	for n := int64(0); ; n++ {
		val := n*n + a*n + b
		if val < 2 {
			return n
		}
		if !isPrime(val) {
			return n
		}
	}
}
