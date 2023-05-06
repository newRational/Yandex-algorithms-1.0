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
		if i, err := strconv.Atoi(p); err == nil {
			ints = append(ints, i)
		}
	}
	return
}

func main() {
	ints := ScanInts()
	here := make(map[int]struct{})

	for _, i := range ints {
		if _, ok := here[i]; !ok {
			here[i] = struct{}{}
		}
	}

	fmt.Print(len(here))
}
