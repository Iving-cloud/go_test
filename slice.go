package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	fmt.Println(len(s))
	s = append(s, "d")
	fmt.Println(s)

	var a []int
	a = append(a, 3)
	fmt.Println(a, len(a))
	a = append(a, 4)
	fmt.Println(a, len(a))
}
