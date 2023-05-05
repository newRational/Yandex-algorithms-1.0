package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mine string

func (m *mine) addOne(mat [][]int) {
	pair := strings.Split(string(*m), " ")
	x, _ := strconv.Atoi(pair[0])
	y, _ := strconv.Atoi(pair[1])
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			mat[i][j]++
		}
	}
}

func main() {
	var n, m, k int
	s := bufio.NewScanner(os.Stdin)

	fmt.Scanf("%d%d%d", &n, &m, &k)
	mines := make(map[mine]*struct{}, k)
	n += 2
	m += 2

	for i := 0; i < k; i++ {
		s.Scan()
		mines[mine(s.Text())] = nil
	}

	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)
	}

	for m := range mines {
		m.addOne(mat)
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-2; j++ {
			if _, ok := mines[mine(fmt.Sprintf("%d %d", i, j))]; ok {
				fmt.Printf("* ")
			} else {
				fmt.Printf("%d ", mat[i][j])
			}
		}
		if _, ok := mines[mine(fmt.Sprintf("%d %d", i, m-2))]; ok {
			fmt.Printf("*\n")
		} else {
			fmt.Printf("%d\n", mat[i][m-2])
		}
	}
}
