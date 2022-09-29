package main

import (
	"fmt"
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
}
