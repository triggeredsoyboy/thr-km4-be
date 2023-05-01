package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	//your code here!
	newDataSlice := []int{newData}

	if position == "up" {
		return append(newDataSlice, data...)
	} else if position == "down" {
		return append(data, newDataSlice...)
	} else {
		panic("Invalid position parameter")
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	newData := 6

	position := "up"
	result := AddElement(data, newData, position)
	fmt.Println(result)

	position = "down"
	result = AddElement(data, newData, position)
	fmt.Println(result)
}
