package main

import "fmt"

func main() {
	var a, b, c, d, e, f int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)
	fmt.Scanf("%d", &e)
	fmt.Scanf("%d", &f)

	if a == 0 && b == 0 && c == 0 && d == 0 && (e != 0 || f != 0) {
		fmt.Print(0)
		return
	}

	if a == 0 && b == 0 && c == 0 && d == 0 && e == 0 && f == 0 {
		fmt.Print(5)
		return
	}

	if a == 0 && c == 0 && b != 0 && d != 0 {
		if e/b == f/d {
			fmt.Print(4, e/b)
		} else {
			fmt.Print(0)
		}
		return
	}

	if b == 0 && d == 0 && a != 0 && c != 0 {
		if e/a == f/c {
			fmt.Print(3, e/a)
		} else {
			fmt.Print(0)
		}
		return
	}

	if a/b == c/d && e/b == f/d {
		fmt.Print(1, -a/b, e/b)
		return
	}

	fmt.Print(2)
}
