package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ScanInts() (ints []int) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i, _ := strconv.Atoi(s.Text())
		if i == -2000000000 {
			return
		}
		ints = append(ints, i)
	}
	return
}

func main() {
	s := 0b011
	ints := ScanInts()
	for i := 1; i < len(ints); i++ {
		if ints[i] == ints[i-1] {
			s |= 0b100
		} else if ints[i] > ints[i-1] {
			s &= 0b110
		} else if ints[i] < ints[i-1] {
			s &= 0b101
		}
	}

	switch s {
	case 0b111:
		fmt.Print("CONSTANT")
	case 0b010:
		fmt.Print("ASCENDING")
	case 0b110:
		fmt.Print("WEAKLY ASCENDING")
	case 0b001:
		fmt.Print("DESCENDING")
	case 0b101:
		fmt.Print("WEAKLY DESCENDING")
	default:
		fmt.Print("RANDOM")
	}
}
