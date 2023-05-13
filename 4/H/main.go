package main

import (
	"fmt"
	"os"
)

func cmp(s1, s2 []int) bool {
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	f, _ := os.Open("input.txt")
	var n, m int
	var w, s string
	fmt.Fscanf(f, "%d%d\n%s\n%s", &n, &m, &w, &s)

	sw := make([]int, 150)
	sr := make([]int, 150)

	for _, r := range w {
		sw[r]++
	}

	l := 0
	r := 0
	res := 0
	fit := false
	for r < m {
		r++

		if sw[s[r-1]] == 0 {
			l = r
			sr = make([]int, 150)
			fit = false
			continue
		}

		sr[s[r-1]]++

		if r-l == n {
			fit = cmp(sr, sw)
		} else if r-l > n {
			sr[s[l]]--
			l++
			if s[r-1] != s[l-1] {
				if fit {
					fit = false
				} else {
					fit = cmp(sr, sw)
				}
			}
		}

		if fit {
			res++
		}
	}

	fmt.Println(res)
}
