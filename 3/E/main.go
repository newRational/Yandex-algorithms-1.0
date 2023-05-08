package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	m := map[rune]struct{}{}
	var x, y, z, n string
	var k int
	fmt.Scanf("%s%s%s", &x, &y, &z)
	fmt.Scanf("%s", &n)
	a, _ := utf8.DecodeRuneInString(x)
	b, _ := utf8.DecodeRuneInString(y)
	c, _ := utf8.DecodeRuneInString(z)

	for _, r := range n {
		m[r] = struct{}{}
	}
	for s := range m {
		if s != a && s != b && s != c {
			k++
		}
	}

	fmt.Print(k)
}
