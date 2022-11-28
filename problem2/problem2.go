package main

import (
	"fmt"
)

func GetMinMax(numbers ...*int) (min int, max int) { // cari nilai minimal sama maximum dari 6 data
	// your code here

	for i := 0; i < 6; i++ { // 6x perulangan
		if *numbers[i] > max { // data pertama > nilai
			max = *numbers[i]
		}

	}
	min = max
	for i := 0; i < 6; i++ { // 6x perulangan
		if *numbers[i] < min { // data pertama < nilai max num[]->
			min = *numbers[i]
		}

	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
