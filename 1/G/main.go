package main

import "fmt"

func main() {
	var n, k, m, a, b, r1, r2, sum int
	fmt.Scanf("%d%d%d", &n, &k, &m)

	if n < k || k < m {
		fmt.Print(0)
		return
	}

	b = k / m
	r2 = k - b*m

	for n >= k {
		a = n / k
		sum += a * b

		r1 = n - a*k
		n = r1 + a*r2
	}

	fmt.Print(sum)
}
