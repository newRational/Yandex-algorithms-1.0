package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)
	fmt.Scanf("%d", &n3)
	if n1+n2 <= n3 {
		fmt.Print("NO")
	} else if n2+n3 <= n1 {
		fmt.Print("NO")
	} else if n1+n3 <= n2 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
