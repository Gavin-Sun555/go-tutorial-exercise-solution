package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

// This is an simple exercise for closure.

func fibonacci() func() int {
    cur, next := 0,1
    return func() int {
        cur, next = next, cur+next
        return cur
    } 
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
