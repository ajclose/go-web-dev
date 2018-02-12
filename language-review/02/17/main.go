package main

import "fmt"

func main() {
	s := "I'm sorry dave i can't do that"
	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	fmt.Println(string([]byte(s[:14])))
	fmt.Println(string([]byte(s[10:22])))
	fmt.Println(string([]byte(s[17:])))
	for _, v := range []byte(s) {
		fmt.Println(string(v))
	}
}
