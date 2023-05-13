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
	w := bufio.NewWriter(os.Stdout)
	m := map[string]map[string]int{}

	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		pr, _ := strconv.Atoi(parts[2])

		if _, ok := m[parts[0]]; !ok {
			m[parts[0]] = map[string]int{}
		}
		m[parts[0]][parts[1]] += pr
	}

	for n, pm := range m {
		fmt.Fprintf(w, "%s:\n", n)
		for p, pr := range pm {
			fmt.Fprintf(w, "%s %d\n", p, pr)
		}
	}

	w.Flush()
}
