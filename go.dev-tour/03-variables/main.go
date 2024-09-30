package main

import (
	"fmt"
	"math"
)

type shrut float64

func (f shrut) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	x := shrut(-math.Sqrt2)
	fmt.Println(x.Abs())
}
