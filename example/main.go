package main

import (
	"fmt"
	"gitlab.glaske.net/mglaske/distance"
)

func main() {
	x := 5 * distance.Mile

	x = x + (500 * distance.Foot)

	distance.Imperial = true
	fmt.Printf("Distance: %s\n", x)
	fmt.Printf("Distance in feet: %f\n", x.Feet())
	distance.Imperial = false
	fmt.Printf("Distance: %s\n", x)
}
