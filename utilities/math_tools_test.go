package math_tools

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	first_primes := []int {2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53}
	primes_map := map[int]bool {}
	for _, ele := range first_primes {
		primes_map[ele] = true
	}

	for i := 2; i < 54; i++ {
		div, is_prime := IsPrime(i)
		_, in_map := primes_map[i]
		if is_prime != in_map {
			t.Errorf("%d, is_prime: %t, prime_map: %t", i, is_prime, in_map)
			t.Errorf("%d, has a non-zero divisor for prime number: %d", i, div)
		} 
		if is_prime == false {
			if (i / div) * div != i {
				t.Errorf("%d gets div: %d, which does not divide number evenly", i, div)
			}
		}
	}
}
