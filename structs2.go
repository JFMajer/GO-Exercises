package main

import (
	"fmt"
)

type car struct {
	name     string
	topSpeed float64
}

func nitroBoost(p *car) {
	p.topSpeed += 50
}

func main() {
	var mustang car
	mustang.name = "Mustanc Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang)
}
