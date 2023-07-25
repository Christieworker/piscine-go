package piscine

import (
	"fmt"
)

func bort(a, b, c, d int) int {
	lens := 0
	number := []int{a, b, c, d}
	for range number {
		lens++
	}
	fmt.Println(lens)

	if lens%2 != 0 {
		for i := 0; i < lens; i++ {
			for j := 0; j < lens; j++ {
				if number[i] < number[j] {
					number[j], number[i] = number[i], number[j]
				}
			}
		}
		return number[2]
	}
	// return number[2]

	m := lens / 2
	n := (lens / 2) + 1
	l := (number[m] + number[n]) // 2

	if lens%2 == 0 {
		for i := 0; i < lens; i++ {
			for j := 0; j < lens; j++ {
				if number[i] < number[j] {
					number[j], number[i] = number[i], number[j]
				}
			}
		}
	}
	return l
}

func main() {
	middle := bort(2, 3, 8, 5)
	fmt.Println(middle)
}
