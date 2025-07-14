package main

import "fmt"

func add(x int, y int) int { // someone can also write x,y int
	return x + y
}

func swap(x, y string) (string, string) { // multiple return values possible
	return y, x
}

func main() {
	fmt.Println(add(42, 58))

	fmt.Println(swap("Hello", "World"))
}
