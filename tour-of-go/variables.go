package main

import "fmt"

var c, python, java bool // variables at package level

func main() {
	var i int // variable at function level

	k := 3 // short assigment statement (instead fo using var)
	fmt.Println(i, c, python, java, k)
}
