package main

func task(n int, nums []int) (int, []int) {
	if n == 1 {
		return 0, []int{}
	}

	if n == 2 {
		if nums[0] == nums[1] {
			return 0, []int{}
		} else {
			return 1, []int{nums[0]}
		}
	}

	var c, l, r int

	r = n - 1
	c = n / 2
	if n%2 == 0 {
		if nums[c] == nums[c-1] {
			l = 0
		} else {
			l = 1
		}
	} else {
		l = 0
	}

	for l < r {
		if nums[l] == nums[r] {
			l++
			r--
		} else {
			for r < n && nums[l] != nums[r] {
				r++
			}
			r--
			if nums[c-1] == nums[c] {
				l++
			} else {
				c++
				if nums[c-1] == nums[c] {
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
