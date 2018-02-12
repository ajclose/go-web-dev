package main

import "fmt"

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}
func main() {
	var g1 gator
	g1 = 8
	g1.greeting()

}
