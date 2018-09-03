package math_tools

import (
	"errors"
	"fmt"
	"math"
)

// Chech whether the integer is prime
// returns true if integer is prime
// returns false if integer is not prime
func IsPrime(x int) (largest_div int, isPrime bool, err error) {

	defer func() {
		if r := recover(); r != nil {
			largest_div = 0
			isPrime = false
			err = errors.New(fmt.Sprintf("%s", r))
		}
	}()

	if x < 2 {
		return 0, false, nil
	}
	if x == 2 {
		return 0, true, nil
	}
	if x%2 == 0 {
		return 2, false, nil
	} else {
		limit := int(math.Sqrt(float64(x))) + 1
		for i := 3; i < limit; i += 2 {
			if x%i == 0 {
				return i, false, nil
			}
		}
	}
	return 0, true, nil
}
