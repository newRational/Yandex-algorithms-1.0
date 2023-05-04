package main

import "fmt"

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	var n, x int

	fmt.Scanf("%d", &n)
	arr := make([]int, n)

	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	fmt.Scanf("%d", &x)

	diff := abs(x - arr[0])
	mindiff := diff
	res := arr[0]

	for i := 1; i < len(arr); i++ {
		diff = abs(x - arr[i])
		if diff < mindiff {
			mindiff = diff
			res = arr[i]
		}
	}

	fmt.Print(res)
}
