package main

import (
	"fmt"
	"math"
)

type pair struct {
	sum int
	max int
}

func main() {
	var a, b, c, d int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)

	s := math.MaxInt
	var m, n, m1, n1 int

	m = a + d
	if b > c {
		n = b
	} else {
		n = c
	}
	if m*n < s {
		s = m * n
		m1 = m
		n1 = n
	}

	m = b + c
	if a > d {
		n = a
	} else {
		n = d
	}
	if m*n < s {
		s = m * n
		m1 = m
		n1 = n
	}

	m = a + c
	if b > d {
		n = b
	} else {
		n = d
	}
	if m*n < s {
		s = m * n
		m1 = m
		n1 = n
	}

	m = b + d
	if a > c {
		n = a
	} else {
		n = c
	}
	if m*n < s {
		s = m * n
		m1 = m
		n1 = n
	}

	fmt.Println(m1, n1)
}
