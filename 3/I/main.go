package main

import "fmt"

func main() {
	var n, m int
	var s string
	l := map[string]int{}

	fmt.Scanf("%d", &n)

	a := make([]string, 0, n)

	for i := 1; i < n; i++ {
		fmt.Scanf("%d", &m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%s", &s)
			l[s]++
		}
	}

	fmt.Scanf("%d", &m)
	for j := 0; j < m; j++ {
		fmt.Scanf("%s", &s)
		l[s]++
		if l[s] == n {
			a = append(a, s)
		}
	}

	fmt.Println(len(a))
	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println(len(l))
	for k := range l {
		fmt.Println(k)
	}
}
