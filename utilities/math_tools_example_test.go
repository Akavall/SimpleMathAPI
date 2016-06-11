package math_tools

import (
	"fmt"
)

func ExampleIsPrime() {
	fmt.Println(IsPrime(10))
	fmt.Println(IsPrime(19))
	fmt.Println(IsPrime(1))
	// Output:
	// false
	// true
	// false

}
