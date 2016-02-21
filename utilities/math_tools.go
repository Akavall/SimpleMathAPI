package math_tools

import (
	"math"
)

func IsPrime(x int) bool {
	if x < 2 { return false }
	if x == 2 { return true }
	if x % 2 == 0 { 
		return false
	} else {
		limit := int(math.Sqrt(float64(x))) + 2
		for i := 3; i < limit; i += 2 {
			if x % i == 0 {
				return false
			}
		}
	}
	return true 
}
