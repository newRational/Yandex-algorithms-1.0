package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScanInts() (a []int64) {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	parts := strings.Split(s.Text(), " ")
	for _, p := range parts {
		i, _ := strconv.Atoi(p)
		a = append(a, int64(i))
	}
	return
}

// a = list(map(int,input().split()))
// max1, max2, max3 = -1000000000000, -1000000000000, -1000000000000
// min1, min2 = 1000000000000, 1000000000000

// for i in a:
//    if i > max1:
//        max3 = max2
//        max2 = max1
//        max1 = i
//    elif i > max2:
//        max3 = max2
//        max2 = i
//    elif i > max3:
//        max3 = i
//    if i < min1:
//        min2 = min1
//        min1 = i
//    elif i < min2:
//        min2 = i

// if max1 * max2 * max3 > max1 * min1 * min2:
//    print(max3, max2, max1)
// else:
//    print(min1, min2, max1)

func task(a []int64) {
	max1 := int64(-1000000000000)
	max2 := int64(-1000000000000)
	max3 := int64(-1000000000000)

	min1 := int64(1000000000000)
	min2 := int64(1000000000000)

	// for i in a:
	//    if i > max1:
	//        max3 = max2
	//        max2 = max1
	//        max1 = i
	//    elif i > max2:
	//        max3 = max2
	//        max2 = i
	//    elif i > max3:
	//        max3 = i
	//
	//    if i < min1:
	//        min2 = min1
	//        min1 = i
	//    elif i < min2:
	//        min2 = i

	for _, i := range a {
		if i > max1 {
			max3 = max2
			max2 = max1
			max1 = i
		} else if i > max2 {
			max3 = max2
			max2 = i
		} else if i > max3 {
			max3 = i
		}

		if i < min1 {
			min2 = min1
			min1 = i
		} else if i < min2 {
			min2 = i
		}
	}

	// if max1 * max2 * max3 > max1 * min1 * min2:
	//    print(max3, max2, max1)
	// else:
	//    print(min1, min2, max1)

	if max1*max2*max3 > max1*min1*min2 {
		fmt.Println(max3, max2, max1)
	} else {
		fmt.Println(min1, min2, max1)
	}
}

func main() {
	task(ScanInts())
}
