package main

type Calculator struct{}

func (*Calculator) Add(x, y int) int {
	return x + y
}

func (*Calculator) Sub(x, y int) int {
	return x - y
}

func (*Calculator) Mult(x, y int) int {
	return x * y
}

func (*Calculator) Div(x, y int) int {
	return x / y
}
