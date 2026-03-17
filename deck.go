package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, j := range d {
		fmt.Println(i, j)
	}
}
