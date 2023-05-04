package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Scanf("%d", &nums[i])
	}

	if n == 1 {
		fmt.Print(0)
		return
	}

	if n == 2 {
		if nums[0] == nums[1] {
			fmt.Print(0)
		} else {
			fmt.Println(1)
			fmt.Print(nums[0])
		}
		return
	}

	var c, cl, l, r int

	r = n - 1
	c = n / 2
	if n%2 == 0 {
		if nums[c] == nums[c-1] {
			cl = 2
			l = 0
		} else {
			cl = 1
			l = 1
		}
	} else {
		cl = 1
		l = 0
	}

	for l < r {
		if nums[l] == nums[r] {
			l++
			r--
		} else {
			if cl == 2 {
				l++
				cl = 1
			} else {
				for r < n && nums[l] != nums[r] {
					r++
				}
				r--
				c++
				if nums[c-1] == nums[c] {
					l++
					cl = 2
				} else {
					l = 2*c - r
				}
			}
		}
	}

	l -= n - 1 - r
	fmt.Print(l)
	if l > 0 {
		fmt.Println()
		ans := nums[:l]
		for i := l - 1; i > 0; i-- {
			fmt.Print(ans[i], " ")
		}
		fmt.Print(ans[0])
	}
}
