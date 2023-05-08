package main

import (
	"fmt"
	"sort"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printSlc(s []int) {
	for _, i := range s {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func readMap(m map[int]struct{}, n int) {
	var x int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		m[x] = struct{}{}
	}
}

func main() {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	min := Min(m, n)
	a := make(map[int]struct{}, n)
	b := make(map[int]struct{}, m)
	g := make(map[int]struct{}, min)

	var x int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		a[x] = struct{}{}
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &x)
		b[x] = struct{}{}
	}

	gSlc := make([]int, 0, min)

	for i := range a {
		if _, ok := b[i]; ok {
			g[i] = struct{}{}
			gSlc = append(gSlc, i)
		}
	}

	aSlc := make([]int, 0, n-len(gSlc))
	bSlc := make([]int, 0, m-len(gSlc))

	for i := range a {
		if _, ok := g[i]; !ok {
			aSlc = append(aSlc, i)
		}
	}

	for i := range b {
		if _, ok := g[i]; !ok {
			bSlc = append(bSlc, i)
		}
	}

	sort.Slice(gSlc, func(i, j int) bool { return gSlc[i] < gSlc[j] })
	sort.Slice(aSlc, func(i, j int) bool { return aSlc[i] < aSlc[j] })
	sort.Slice(bSlc, func(i, j int) bool { return bSlc[i] < bSlc[j] })

	fmt.Println(len(gSlc))
	printSlc(gSlc)

	fmt.Println(len(aSlc))
	printSlc(aSlc)

	fmt.Println(len(bSlc))
	printSlc(bSlc)
}
