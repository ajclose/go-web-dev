package main

import "fmt"

type gator int

type flamingo bool

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
	greeting()
}

func bayou(sc swampCreature) {
	sc.greeting()
}
func main() {
	var g1 gator
	g1 = 8
	bayou(g1)
	var f1 flamingo
	f1 = false
	bayou(f1)

}
