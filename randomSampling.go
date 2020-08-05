package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	infiniteStream := [10]int{23, 34, 56, 88, 12, 39, 100, 55, 38, 76}

	var randomSample int

	rand.Seed(time.Now().UnixNano())

	for i, e := range infiniteStream {
		if i == 0 {
			randomSample = e
		} else {
			if rand.Intn(i+1) < 1 {
				randomSample = e
			}
		}
	}

	fmt.Println(randomSample)
}
