package main

import "fmt"

// func swap(x, y int) (int, int) {
// 	return y, x
// }

// func main() {
// 	// x, y := 42, 24
// 	// we can return multiple values from a function
// 	// and assign them to multiple variables
// 	// x, y = swap(x, y)

// 	// fmt.Printf(`Type of value %T, %T`, x, y)

// 	// we can also use the blank identifier to ignore a value
// 	// _, y = swap(x, y)

// 	// we can also use the blank identifier to ignore multiple values
// 	// _, _ = swap(x, y)

// 	// we can also write functions in variables like this
// 	x := func(a ...int) int {
// 		if len(a) > 0 {
// 			return a[0]
// 		}
// 		return 0
// 	}

// 	y := func () int {
// 		return 24
// 	}

// 	fmt.Println(x(1,2,4), y)
// }

// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	fmt.Printf("Type of value %T:%p, %T:%p \n", pos, neg, pos, neg)
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-1*i),
		)
	}
}