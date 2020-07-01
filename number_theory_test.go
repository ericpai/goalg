package goalg_test

import (
	"testing"

	"github.com/ericpai/goalg"
	"github.com/stretchr/testify/assert"
)

func TestPrimes(t *testing.T) {
	primes := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
		127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
		179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
		233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
		283, 293, 307, 311, 313, 317, 331, 337, 347, 349,
		353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
		419, 421, 431, 433, 439, 443, 449, 457, 461, 463,
		467, 479, 487, 491, 499, 503, 509, 521, 523, 541,
		547, 557, 563, 569, 571, 577, 587, 593, 599, 601,
		607, 613, 617, 619, 631, 641, 643, 647, 653, 659,
		661, 673, 677, 683, 691, 701, 709, 719, 727, 733,
		739, 743, 751, 757, 761, 769, 773, 787, 797, 809,
		811, 821, 823, 827, 829, 839, 853, 857, 859, 863,
		877, 881, 883, 887, 907, 911, 919, 929, 937, 941,
		947, 953, 967, 971, 977, 983, 991, 997,
	}
	assert.Equal(t, primes, goalg.Primes(1000))
	assert.Equal(t, primes[:len(primes)-1], goalg.Primes(997))
	assert.Equal(t, []int(nil), goalg.Primes(1))
	assert.Equal(t, primes, goalg.FastPrimes(1000))
	assert.Equal(t, primes[:len(primes)-1], goalg.FastPrimes(997))
	assert.Equal(t, []int(nil), goalg.FastPrimes(1))
}

func TestFastPow(t *testing.T) {
	testCases := []struct {
		base, pow, expected int
	}{
		{1, 100, 1},
		{0, 100, 0},
		{2, 10, 1024},
		{2, -10, 0},
		{2, 0, 1},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, goalg.FastPow(tc.base, tc.pow))
	}
}

func TestFastFloatPow(t *testing.T) {
	testCases := []struct {
		base     float64
		pow      int
		expected float64
	}{
		{1, 100, 1},
		{0, 100, 0},
		{2, 10, 1024},
		{2, -10, 0.000976525},
		{2, 0, 1},
	}
	for _, tc := range testCases {
		assert.InDelta(t, tc.expected, goalg.FastFloatPow(tc.base, tc.pow), 0.0000001)
	}
}

func TestFastModPow(t *testing.T) {
	mod := 97
	testCases := []struct {
		base, pow, expected int
	}{
		{1, 100, 1},
		{0, 100, 0},
		{2, 10, 54},
		{2, -10, 0},
		{2, 0, 1},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, goalg.FastModPow(tc.base, tc.pow, mod))
	}
}
