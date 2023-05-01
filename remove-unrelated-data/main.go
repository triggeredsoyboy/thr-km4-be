package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	//your code here
	delete(dataMap, key)
	return dataMap
}

func main() {
	data := map[string]interface{}{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}

	keyToRemove := "address"

	fmt.Println("Data sebelum dihapus:")
	fmt.Println(data)

	newData := removeUnrelated(data, keyToRemove)

	fmt.Println("Data setelah dihapus:")
	fmt.Println(newData)
}
