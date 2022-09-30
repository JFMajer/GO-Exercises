package main

import (
	"fmt"
	"sort"
)

func main() {
	var v int
	var ok bool

	m := map[string]int{"a": 0, "b": 1, "c": 2}
	for v := range m {
		fmt.Println(m[v])
	}

	for v, k := range m {
		fmt.Println(k, v)
	}
	v, ok = m["d"]
	fmt.Println(v, ok)
	v, ok = m["a"]
	fmt.Println(v, ok)
	value, exist := m["b"]
	fmt.Println(value, exist)
	delete(m, "b")
	value, exist = m["b"]
	fmt.Println(value, exist)

	newMap := map[string]int{"Teddy": 5, "Jenny": 4, "Bob": 3}
	fmt.Println(newMap)
	//maps are printed in random order
	//therefore we are sorting keys first
	var names []string
	for k := range newMap {
		names = append(names, k)
	}
	fmt.Println(names)
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s got grade %v\n", name, newMap[name])
	}

	//doesn't work, need to use make function or use map literal
	// var yetAnotherMap map[string]string
	// yetAnotherMap["key"] = "value"
	// fmt.Println(yetAnotherMap)  //panic: assignment to entry in nil map

}
