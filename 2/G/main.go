package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScanInts() (ints []int) {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	parts := strings.Split(s.Text(), " ")
	for _, p := range parts {
		i, _ := strconv.Atoi(p)
		ints = append(ints, i)
	}
	return
}

func Min(arg1, arg2 int) int {
	if arg1 < arg2 {
		return arg1
	}
	return arg2
}

func Max(arg1, arg2 int) int {
	if arg1 > arg2 {
		return arg1
	}
	return arg2
}

func task(ints []int) {
	max1 := Max(ints[0], ints[1])
	min2 := max1
	min1 := Min(ints[0], ints[1])
	max2 := min1

	for i := 2; i < len(ints); i++ {
		if ints[i] > max1 {
			max2 = max1
			max1 = ints[i]
		} else if ints[i] > max2 {
			max2 = ints[i]
		}

		if ints[i] < min1 {
			min2 = min1
			min1 = ints[i]
		} else if ints[i] < min2 {
			min2 = ints[i]
		}
	}

	if max1*max2 > min1*min2 {
		fmt.Print(max2, max1)
	} else {
		fmt.Print(min1, min2)
	}
}

func main() {
	task(ScanInts())
}
