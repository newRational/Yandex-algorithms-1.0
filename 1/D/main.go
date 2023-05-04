package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)
	fmt.Scanf("%f", &c)

	if c < 0 {
		fmt.Println("NO SOLUTION")
		return
	}

	if a == 0 {
		if math.Pow(1, 2) == b {
			fmt.Println("MANY SOLUTIONS")
		} else {
			fmt.Println("NO SOLUTION")
		}
		return
	}

	res := (c*c - b) / a

	if res-math.Trunc(res) == 0.0 {
		fmt.Println(res)
	} else {
		fmt.Println("NO SOLUTION")
	}
}
