package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(60.5))
	fmt.Println(math.Round(60.2))
	fmt.Println(math.Floor(60.5))
	fmt.Println(math.Ceil(60.2))

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}
