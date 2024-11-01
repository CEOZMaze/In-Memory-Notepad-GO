package main

import "fmt"

func main() {
	var a, b, c bool
	fmt.Scan(&a, &b, &c)

	fmt.Println(Any(a, b, c))
}

func Any(values ...bool) bool {

	for _, v := range values {
		if v {
			return true
		}
	}
	return false
}
