package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64, round int) {

    //initialize variables z to 1 and round to 0
	z = 1.0
	round = 0

    //Loop until the difference is small enough
	for math.Abs((z*z - x)/(2*z)) > 1e-15*z {
		z -= (z*z - x) / (2*z)
		round++
	}
	return
}

func main() {

    // can change to any number for test.
	fmt.Println(Sqrt(512435425124.1))
	fmt.Println(math.Sqrt(512435425124.1))
}
