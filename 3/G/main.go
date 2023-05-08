package main

import (
	"fmt"
)

func main() {
	var a, b, n, k int
	var s string
	m := map[string]struct{}{}
	fmt.Scanf("%d", &n)
	k = n
	for i := 0; i < n; i++ {
		fmt.Scanf("%d%d", &a, &b)
		s = fmt.Sprintf("%d %d", a, b)
		if _, ok := m[s]; a+b != n-1 || ok || a < 0 || b < 0 {
			k--
		} else {
			m[s] = struct{}{}
		}
	}
	fmt.Print(k)
}
