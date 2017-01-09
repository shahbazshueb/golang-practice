package main

import "fmt"

type Integer int

func (x Integer) add(y int) Integer {
	return x + Integer(y)
}
func (x Integer) subtract(y int) Integer {
	return x - Integer(y)
}
func (x Integer) divide(y int) Integer {
	return x / Integer(y)
}
func (x Integer) multiply(y int) Integer {
	return x * Integer(y)
}
func (x Integer) sq() Integer {
	return x * x
}
func (x Integer) cube() Integer {
	return x * x * x
}
func main() {
	x := Integer(100)
	fmt.Println(x.add(12))
	fmt.Println(x.cube())

}
