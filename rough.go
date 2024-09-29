package main

func main() {
	x := func () int {
		return 42
	}

	y := func () int {
		return 24
	}

	println(x(), y())
}