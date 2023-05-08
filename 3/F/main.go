package main

import "fmt"

func main() {
	var a, b string
	var n int
	ma := map[string]int{}
	mb := map[string]int{}
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)

	for i := 1; i < len(a); i++ {
		ma[a[i-1:i+1]]++
	}
	for i := 1; i < len(b); i++ {
		mb[b[i-1:i+1]]++
	}
	for k := range mb {
		if l, ok := ma[k]; ok {
			n += l
		}
	}
	fmt.Print(n)
}
