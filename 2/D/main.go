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

func main() {
	ints := ScanInts()
	var c int

	for i := 1; i < len(ints)-1; i++ {
		if ints[i] > ints[i-1] && ints[i] > ints[i+1] {
			c++
			i++
		}
	}

	fmt.Print(c)
}
