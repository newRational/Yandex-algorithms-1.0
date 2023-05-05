package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	f := make([]float64, n)
	s := make([]string, n)

	fmt.Scanf("%f", &f[0])
	for i := 1; i < n; i++ {
		fmt.Scanf("%f%s", &f[i], &s[i])
	}

	var fmin float64 = 30
	var fmax float64 = 4000
	var favr float64

	for i := 1; i < n; i++ {
		favr = (f[i] + f[i-1]) / float64(2)
		if f[i] <= f[i-1] {
			if s[i] == "closer" {
				fmax = math.Min(fmax, favr)
			} else {
				fmin = math.Max(fmin, favr)
			}
		} else {
			if s[i] == "closer" {
				fmin = math.Max(fmin, favr)
			} else {
				fmax = math.Min(fmax, favr)
			}
		}
	}

	fmt.Println(fmin, fmax)
}
