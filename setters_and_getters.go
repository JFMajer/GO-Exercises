package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) setYear(year int) error {
	if year < 0 {
		return errors.New("Year can't be a negative number")
	}
	d.Year = year
	return nil
}

func main() {
	newDate := Date{2005, 6, 7}
	fmt.Println(newDate)
	newDate.setYear(3005)
	fmt.Println(newDate)
	err := newDate.setYear(-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newDate)
}
