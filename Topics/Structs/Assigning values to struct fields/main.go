package main

import "fmt"

type Country struct {
	Name     string
	Capital  string
	Currency string
}

func main() {
	germany := Country{Name: "Germany", Capital: "Berlin", Currency: "Euro"}
	fmt.Printf("%s\n%s\n%s\n", germany.Name, germany.Capital, germany.Currency)
}
