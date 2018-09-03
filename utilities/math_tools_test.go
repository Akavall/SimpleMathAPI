package math_tools

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	first_primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53}
	primes_map := map[int]bool{}
	for _, ele := range first_primes {
		primes_map[ele] = true
	}

	fmt.Println("FAIL")

	for i := 0; i < 54; i++ {
		is_prime := IsPrime(i)
		_, in_map := primes_map[i]
		if is_prime != in_map {
			t.Errorf("%d, is_prime: %t, prime_map: %t", i, is_prime, in_map)
		}
	}
}
