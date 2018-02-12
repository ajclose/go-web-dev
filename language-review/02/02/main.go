package main

import "fmt"

func main() {
	m := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
	}

	fmt.Println(m)
	for k, _ := range m {
		fmt.Println(k)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
