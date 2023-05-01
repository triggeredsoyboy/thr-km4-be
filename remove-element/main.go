package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	//yout code here!
	if position == "left" {
		return arr[1:]
	} else if position == "right" {
		return arr[:len(arr)-1]
	} else {
		fmt.Println("Invalid position parameter")
		return nil
	}
}

func main() {
	arr := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(removeLeftRight(arr, "left"))
	fmt.Println(removeLeftRight(arr, "right"))

	strArr := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeLeftRight(strArr, "left"))
	fmt.Println(removeLeftRight(strArr, "right"))
}
