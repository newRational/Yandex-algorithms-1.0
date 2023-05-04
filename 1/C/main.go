package main

import (
	"fmt"
	"unicode"
)

func main() {
	n := make([]string, 4)
	for i := 0; i < 4; i++ {
		fmt.Scanf("%s", &n[i])
	}
	n[0] = conv(n[0])
	for i := 1; i < 4; i++ {
		fmt.Println(rep(n[0], conv(n[i])))
	}
}

func conv(s string) string {
	var res string
	for _, r := range s {
		if unicode.IsDigit(r) {
			res += string(r)
		}
	}
	if len(res) == 7 {
		res = "495" + res
	} else {
		res = res[1:]
	}
	return res
}

func rep(s1, s2 string) string {
	if s1 == s2 {
		return "YES"
	}
	return "NO"
}
