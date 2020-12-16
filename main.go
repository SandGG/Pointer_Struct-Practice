package main

import "fmt"

type figure struct {
	height, width *int
	name          string
}

func main() {
	var height, width int = 10, 50
	area(&figure{height: &height, width: &width, name: "Rectangle"})
}

func area(shape *figure) {
	*shape.height = 30
	area := *shape.height * *shape.width

	fmt.Println(shape.name, area, "area")
}
