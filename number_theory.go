package goalg

// Primes returns all the prime numbers in [2, n) in ascending order.
// Time complexity: O(nloglogn)
// Space complexity: O(n * sizeof(bool))
// Ref: https://cp-algorithms.com/algebra/sieve-of-eratosthenes.html
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
// Time complexity: O(n)
// Space complexity: O(n * sizeof(int))
// Ref: https://cp-algorithms.com/algebra/sieve-of-eratosthenes.html
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
