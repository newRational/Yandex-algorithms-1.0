package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	var a, b, n, sum int
	fmt.Fscanf(f, "%d", &n)
	m := make(map[int]int, n)

	var parts []string
	for s.Scan() {
		parts = strings.Split(s.Text(), " ")
		a, _ = strconv.Atoi(parts[0])
		b, _ = strconv.Atoi(parts[1])

		if v, ok := m[a]; !ok || b > v {
			m[a] = b
			sum += b - v
		}
	}

	fmt.Println(sum)
}
