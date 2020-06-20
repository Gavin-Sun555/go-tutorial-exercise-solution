package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

//error func for the error type
func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}

func Sqrt(x float64) (float64, error) {

    // return an error if the num is negative
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
    
    // otherwise, return the result     
	}else {
		return math.Sqrt(x), nil
	}
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
