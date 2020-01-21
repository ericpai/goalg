package goalg

// Primes returns all the prime numbers in [2, n) in ascending order.
// Time complexity: O(nloglogn)
// Space complexity: O(n)
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
