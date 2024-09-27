package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favrouit random number is this\t", rand.Int31())
	fmt.Printf("this is my favourite number \t\t %d \n", rand.Intn(100))
}
