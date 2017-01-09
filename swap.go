package main

import "fmt"

func main() {
	i, j := 10, 30
	fmt.Println(i, j)
	// call swap here
	swap(&i, &j)
	fmt.Println(i, j)
}

func swap(x, y *int) {
	// swap values without using any temporary variables
	*x, *y = *y, *x
}
