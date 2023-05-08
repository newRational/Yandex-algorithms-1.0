package main

import "fmt"

func main() {
	var x, n int
	m := map[int]struct{}{}

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		m[x] = struct{}{}
		fmt.Scanf("%d", &x)
	}

	fmt.Print(len(m))
}
