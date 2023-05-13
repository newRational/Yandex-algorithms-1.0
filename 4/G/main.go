package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	DEPOSIT  = "DEPOSIT"
	WITHDRAW = "WITHDRAW"
	BALANCE  = "BALANCE"
	TRANSFER = "TRANSFER"
	INCOME   = "INCOME"
)

type Bank struct {
	Clients map[string]int
	W       io.Writer
}

func (b *Bank) Do(op string, details ...string) {
	switch op {
	case DEPOSIT:
		name := details[0]
		sum, _ := strconv.Atoi(details[1])
		b.Clients[name] += sum
	case WITHDRAW:
		name := details[0]
		sum, _ := strconv.Atoi(details[1])
		b.Clients[name] -= sum
	case TRANSFER:
		from := details[0]
		to := details[1]
		sum, _ := strconv.Atoi(details[2])
		b.Clients[from] -= sum
		b.Clients[to] += sum
	case INCOME:
		p, _ := strconv.Atoi(details[0])
		for c := range b.Clients {
			if b.Clients[c] > 0 {
				b.Clients[c] += b.Clients[c] * p / 100
			}
		}
	case BALANCE:
		name := details[0]
		if sum, ok := b.Clients[name]; ok {
			fmt.Fprintln(b.W, sum)
		} else {
			fmt.Fprintln(b.W, "ERROR")
		}
	}
}

func main() {
	f, _ := os.Open("input.txt")
	s := bufio.NewScanner(f)
	w := bufio.NewWriter(os.Stdout)
	b := Bank{
		Clients: map[string]int{},
		W:       w,
	}

	var parts []string
	for s.Scan() {
		parts = strings.Split(s.Text(), " ")
		b.Do(parts[0], parts[1:]...)
	}

	w.Flush()
}
