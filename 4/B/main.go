package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	m := map[string]int{}

	for s.Scan() {
		if v, ok := m[s.Text()]; !ok {
			fmt.Fprintf(w, "%d ", 0)
		} else {
			fmt.Fprintf(w, "%d ", v)
		}
		m[s.Text()]++
	}

	w.Flush()
}
