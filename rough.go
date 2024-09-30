// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// x := func () int {
// 	// 	return 42
// 	// }

// 	// y := func () int {
// 	// 	return 24
// 	// }

// 	// println(x(), y())

// 	t := time.Now()

// 	fmt.Printf("%T", t.UTC())
// }

package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	for _, char : range s {
		m[char] += 1;
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
