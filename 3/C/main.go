package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func Min(arg1, arg2 int) int {
	if arg1 < arg2 {
		return arg1
	}
	return arg2
}

func printSlc(w io.Writer, s []int) {
	for _, i := range s {
		fmt.Fprintf(w, "%d ", i)
	}
	fmt.Fprintln(w)
}

func readMap(r io.Reader, m map[int]struct{}, n int) {
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d\n", &x)
		m[x] = struct{}{}
	}
}

func main() {
	var n, m int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanf(r, "%d %d\n", &n, &m)

	a := make(map[int]struct{}, n)
	b := make(map[int]struct{}, m)
	g := make(map[int]struct{}, Min(m, n))

	readMap(r, a, n)
	readMap(r, b, m)

	aSlc := make([]int, 0)
	bSlc := make([]int, 0)
	gSlc := make([]int, 0)

	for i := range a {
		if _, ok := b[i]; ok {
			g[i] = struct{}{}
			gSlc = append(gSlc, i)
		}
	}

	for i := range a {
		if _, ok := g[i]; !ok {
			aSlc = append(aSlc, i)
		}
	}

	for i := range b {
		if _, ok := g[i]; !ok {
			bSlc = append(bSlc, i)
		}
	}

	sort.Slice(gSlc, func(i, j int) bool { return gSlc[i] < gSlc[j] })
	sort.Slice(aSlc, func(i, j int) bool { return aSlc[i] < aSlc[j] })
	sort.Slice(bSlc, func(i, j int) bool { return bSlc[i] < bSlc[j] })

	fmt.Fprintln(w, len(gSlc))
	printSlc(w, gSlc)

	fmt.Fprintln(w, len(aSlc))
	printSlc(w, aSlc)

	fmt.Fprintln(w, len(bSlc))
	printSlc(w, bSlc)

	w.Flush()
}
