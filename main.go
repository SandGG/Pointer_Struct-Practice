package main

import "fmt"

type figure struct {
	height, width *int
	name          string
}

func main() {
	var height, width int = 20, 10
	rectangle := figure{&height, &width, "Rectangle"}
	fmt.Println(rectangle)
	area(&rectangle)
	fmt.Println(rectangle)
}

func area(shape *figure) {
	*shape.height = 30
	area := *shape.height * *shape.width

	fmt.Println(shape.name, area, "area")
}
