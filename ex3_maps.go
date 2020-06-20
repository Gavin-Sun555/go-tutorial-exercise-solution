package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// initialize the counts variable
    counts := make(map[string]int)

    // separate the strings into slices of words
    words := strings.Fields(s)

    // iterate over the slice
    for _, word := range words {
        counts[word]++
    }

    return counts
}

func main() {
	wc.Test(WordCount)
}
