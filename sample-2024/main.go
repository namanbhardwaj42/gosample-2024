// gosample-2024
package main

import (
	"fmt"
	"sample-2024/logics"
)

func main() {
	fmt.Println("-------------------")

	a := 10
	b := 20
	res := logics.Sum(a, b)

	if res == 30 {
		logics.Success()
		fmt.Printf(" : sum of %d and %d is %d \n\n", a, b, res)
	} else {
		logics.Failed()
		fmt.Printf(" : not a valid response\n\n")
	}

	fmt.Println("-------------------")
}
