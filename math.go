package main

import (
	"fmt"
	"math"
)

func main() {
	ceil := math.Ceil(1.40)
	floor := math.Floor(1.60)
	round := math.Round(1.60)
	min := math.Min(10, 20)
	max := math.Max(10, 20)

	fmt.Println(ceil)
	fmt.Println(floor)
	fmt.Println(round)
	fmt.Println(min)
	fmt.Println(max)
}
