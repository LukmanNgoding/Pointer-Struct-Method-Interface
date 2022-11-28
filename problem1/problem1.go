package main

import "fmt"

func Swap(a, b *int) (int, int) {
	//your code here
	*a, *b = *b, *a
	// a,b 20, 10
	return *a, *b
}

func main() {
	a := 10
	b := 20

	Swap(&a, &b)
	fmt.Println(a, b)
}
