package main

import "fmt"

func main() {
	var prefix, name1, name2 string
	fmt.Scan(&prefix, &name1, &name2)

	for _, line := range Greeting(prefix, name1, name2) {
		fmt.Println(line)
	}
}

func Greeting(prefix string, name ...string) []string {
	names := make([]string, len(name))
	for _, v := range name {
		names = append(names, prefix+" "+v)
	}
	return names
}
