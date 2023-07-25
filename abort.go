package piscine

func Abort(a, b, c, d, e int) int {
	lens := 0
	number := []int{a, b, c, d, e}
	for range number {
		lens++
	}
	for i := 0; i < lens; i++ {
		for j := 0; j < lens; j++ {
			if number[i] < number[j] {
				number[j], number[i] = number[i], number[j]
			}
		}
	}
	return number[2]
}
