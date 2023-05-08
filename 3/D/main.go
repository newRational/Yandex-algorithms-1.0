package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := map[string]struct{}{}
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		m[s.Text()] = struct{}{}
	}
	fmt.Println(len(m))
}
