package math_tools

import (
	"math"
)

// Chech whether the integer is prime
// returns true if integer is prime
// returns false if integer is not prime 
func IsPrime(x int) (int, bool) {
	if x < 2 { return 0, false }
	if x == 2 { return 0, true }
	if x % 2 == 0 { 
		return 2, false
	} else {
		limit := int(math.Sqrt(float64(x))) + 1
		for i := 3; i < limit; i += 2 {
			if x % i == 0 {
				return i, false
			}
		}
	}
	return 0, true
}
