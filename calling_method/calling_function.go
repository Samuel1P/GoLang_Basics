package main

import (
	"fmt"
	"math/rand"
)

func generateint() {
	fmt.Println("This is your random int --> ", rand.Intn(25))
}

func main() {
	generateint()
}
