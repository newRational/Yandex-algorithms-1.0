package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var n, m int

	fmt.Scanf("%d", &n)
	max := make([]int, n+1)
	real := make([]int, n+1)

	s.Scan()
	parts := strings.Split(s.Text(), " ")
	for i := 1; i <= n; i++ {
		max[i], _ = strconv.Atoi(parts[i-1])
	}

	fmt.Scanf("%d", &m)

	s.Scan()
	parts = strings.Split(s.Text(), " ")
	for i := 0; i < m; i++ {
		ind, _ := strconv.Atoi(parts[i])
		real[ind]++
	}

	for i := 1; i <= n; i++ {
		if max[i] < real[i] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
