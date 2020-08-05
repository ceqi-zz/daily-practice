package main

import (
	"fmt"
	"math/rand"
)

func main() {

	r := 0.5

	inCircle := 0.0
	total := 5000

	fmt.Println("total points: ", total)

	for i := 0; i < total; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= r*r {
			inCircle++
		}
	}

	fmt.Println("in circle points: ", inCircle)

	pi := 4 * inCircle / float64(total)

	fmt.Printf("%.3f", pi)

}
