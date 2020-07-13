package goalg

// Primes returns all the prime numbers in [2, n) in ascending order.
//
// Time complexity: O(nloglogn).
//
// Space complexity: O(n * sizeof(bool)).
//
// Ref: https://cp-algorithms.com/algebra/sieve-of-eratosthenes.html.
func Primes(n int) []int {
	slot := make([]bool, n)
	var res []int
	for i := 2; i*i < n; i++ {
		if !slot[i] {
			for j := i * i; j < n; j += i {
				slot[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !slot[i] {
			res = append(res, i)
		}
	}
	return res
}

// FastPrimes returns all the prime numbers in [2, n) in ascending order.
//
// Time complexity: O(n).
//
// Space complexity: O(n * sizeof(int)).
//
// Ref: https://cp-algorithms.com/algebra/sieve-of-eratosthenes.html.
func FastPrimes(n int) []int {
	slot := make([]int, n)
	var res []int
	for i := 2; i < n; i++ {
		if slot[i] == 0 {
			slot[i] = i
			res = append(res, i)
		}
		for j := 0; j < len(res) && res[j] <= slot[i] && i*res[j] < n; j++ {
			slot[i*res[j]] = res[j]
		}
	}
	return res
}

// FastPow can compute base**pow in logrithmic time. But it may be overflow.
// If pow is negative, the absolute of the result will less than 1, so 0 will be returned.
//
// If the number is too big, use FastModPow instead.
//
// Time complexity: O(log(pow)).
func FastPow(base, pow int) int {
	if pow < 0 {
		return 0
	}
	res := 1
	for pow > 0 {
		if pow%2 == 1 {
			res *= base
		}
		base *= base
		pow >>= 1
	}
	return res
}

// FastFloatPow can compute base**pow in logrithmic time, and return in float64.
//
// Time complexity: O(log(pow)).
func FastFloatPow(base float64, pow int) float64 {
	if pow < 0 {
		return 1 / FastFloatPow(base, -pow)
	}
	res := 1.0
	for pow > 0 {
		if pow%2 == 1 {
			res *= base
		}
		base *= base
		pow >>= 1
	}
	return res
}

// FastModPow can compute base**pow % mod in logarithmic time. mod should be positive and less than 2^32.
//
// If pow is negative, the absolute of the result will less than 1, so 0 will be returned.
//
// Time complexity: O(log(pow)).
func FastModPow(base, pow, mod int) int {
	if pow < 0 {
		return 0
	}
	res := 1
	for pow > 0 {
		if pow%2 == 1 {
			res *= base
			res %= mod
		}
		base *= base
		base %= mod
		pow >>= 1
	}
	return res
}
