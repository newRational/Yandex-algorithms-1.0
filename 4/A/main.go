package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var parts []string

	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	n = i

	l := make(map[string]string, n)
	r := make(map[string]string, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		parts = strings.Split(sc.Text(), " ")
		l[parts[0]] = parts[1]
		r[parts[1]] = parts[0]
	}

	sc.Scan()

	if v, ok := l[sc.Text()]; ok {
		fmt.Print(v)
	} else {
		fmt.Print(r[sc.Text()])
	}
}
