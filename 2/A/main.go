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

	res := "YES"
	for i := 1; i < len(ints); i++ {
		if ints[i-1] >= ints[i] {
			res = "NO"
			break
		}
	}

	fmt.Print(res)
}
