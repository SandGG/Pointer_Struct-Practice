package main

import "fmt"

type figure struct {
	height, width int
	name          string
}

func main() {
	rectangle := figure{5, 10, "Rectangle"}
	fmt.Println(rectangle)
	area(&rectangle)
	fmt.Println(rectangle)
}

func area(shape *figure) {
	shape.height = 30
	area := shape.height * shape.width

	fmt.Println(shape.name, area, "area")
}
