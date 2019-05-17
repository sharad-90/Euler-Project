package main

import "fmt"

var t = map[float64]struct{}{}
var p = map[float64]struct{}{}
var h = map[float64]struct{}{}
var empty struct{}

func main() {
	for n := 1; n < 200000; n++ {
		con := float64(n)
		n1 := (con*con + con) / 2
		n2 := (3*con*con - con) / 2
		n3 := 2*con*con - con
		t[n1] = empty
		p[n2] = empty
		h[n3] = empty
		_, tbool := t[n1]
		_, pbool := p[n1]
		_, hbool := h[n1]
		if tbool && pbool && hbool {
			fmt.Println(n1)
		}
	}
}
