package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)
	fmt.Scanf("%d", &e)

	ad := a <= d
	ae := a <= e
	bd := b <= d
	be := b <= e
	cd := c <= d
	ce := c <= e

	if ad && (be || ce) || bd && (ae || ce) || cd && (ae || be) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

	return
}
