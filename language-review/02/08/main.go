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

func main() {
	t1 := truck{vehicle{2, "blue"}, true}
	fmt.Println(t1)
	fmt.Println("four wheel drive:", t1.fourWheel)
	s1 := sedan{vehicle{4, "silver"}, false}
	fmt.Println(s1)
	fmt.Println("sedan color:", s1.vehicle.color)
}
