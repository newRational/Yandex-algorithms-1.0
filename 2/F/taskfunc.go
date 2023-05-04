package main

import "fmt"

func task(n int, nums []int) (int, []int) {
	if n == 1 {
		fmt.Print(0)
	}

	if n == 2 {
		if nums[0] == nums[1] {
			fmt.Print(0)
		} else {
			fmt.Println(1)
			fmt.Print(nums[0])
		}
	}

	var c, cl, l, r int

	r = len(nums) - 1
	c = len(nums) / 2
	if len(nums)%2 == 0 {
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

	for c < r && l < r {
		if nums[l] == nums[r] {
			l++
			r--
		} else {
			if cl == 2 {
				l++
				cl = 1
			} else {
				c++
				if nums[c-1] == nums[c] {
					cl = 2
					l++
				} else {
					l = 2*c - r
				}
			}
		}
	}

	l -= n - 1 - r
	ans := nums[:l]
	return l, ans
}
