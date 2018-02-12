package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func (t truck) transportationDevice() string {
	return "Hauls big stuff"
}

func (s sedan) transportationDevice() string {
	return "Makes the ride smooth"
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {
	t1 := truck{vehicle{2, "blue"}, true}
	report(t1)
	s1 := sedan{vehicle{4, "silver"}, false}
	report(s1)
}
