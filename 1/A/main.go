package main

import "fmt"

func main() {
	var t1, t2 int
	var mode string
	fmt.Scanf("%d%d", &t1, &t2)
	fmt.Scanf("%s", &mode)

	switch mode {
	case "freeze":
		if t1 > t2 {
			t1 = t2
		}
	case "heat":
		if t1 < t2 {
			t1 = t2
		}
	case "auto":
		t1 = t2
	}

	fmt.Print(t1)
}
