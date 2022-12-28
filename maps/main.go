package main

import "fmt"

func main() {
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3

	fmt.Println(intMap)

	for k, v := range intMap {
		fmt.Println(k, v)
	}

	addToMap(intMap)

	fmt.Println("after add function: ")
	for k, v := range intMap {
		fmt.Println(k, v)
	}

	delete(intMap, "addition")
	for k, v := range intMap {
		fmt.Println(k, v)
	}

	fmt.Println("element elephant exist in map?")
	fmt.Println(checkOIfExists(intMap, "elephant"))

}

func addToMap(map1 map[string]int) {
	map1["addition"] = 4
}

func checkOIfExists(map1 map[string]int, key1 string) bool {
	_, ok := map1[key1]
	return ok
}
