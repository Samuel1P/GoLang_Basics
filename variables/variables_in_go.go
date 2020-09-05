package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func saystring(x, y string) (string, string) {
	return x, y
}

func main() {
	var a int = 4
	var b int = 7
	c, d := "Hello", "There"
	fmt.Println("Added result ", add(a, b))
	fmt.Println(saystring(c, d))

}
