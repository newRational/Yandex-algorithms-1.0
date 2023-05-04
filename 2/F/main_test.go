package main

import (
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		n    int
		nums []int
		m    int
		ans  []int
	}{
		{
			n:    9,
			nums: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			m:    0,
			ans:  []int{},
		},
		{
			n:    5,
			nums: []int{1, 2, 1, 2, 2},
			m:    3,
			ans:  []int{1, 2, 1},
		},
		{
			n:    5,
			nums: []int{1, 2, 3, 4, 5},
			m:    4,
			ans:  []int{1, 2, 3, 4},
		},
		{
			n:    5,
			nums: []int{1, 2, 3, 3, 2},
			m:    1,
			ans:  []int{1},
		},
		{
			n:    7,
			nums: []int{1, 2, 3, 4, 4, 5, 4},
			m:    4,
			ans:  []int{1, 2, 3, 4},
		},
		{
			n:    7,
			nums: []int{1, 2, 3, 4, 4, 4, 4},
			m:    3,
			ans:  []int{1, 2, 3},
		},
		{
			n:    3,
			nums: []int{1, 2, 2},
			m:    1,
			ans:  []int{1},
		},
		{
			n:    2,
			nums: []int{1, 2},
			m:    1,
			ans:  []int{1},
		},
		{
			n:    2,
			nums: []int{2, 2},
			m:    0,
			ans:  []int{},
		},
		{
			n:    1,
			nums: []int{1},
			m:    0,
			ans:  []int{},
		},
	}

	for _, tt := range tests {
		m, ans := task(tt.n, tt.nums)
		if m != tt.m {
			t.Errorf("m (%d) != tt.m (%d)\n", m, tt.m)
			t.Errorf("ans (%v) != tt.ans (%v)\n", ans, tt.ans)
		}
	}
}
