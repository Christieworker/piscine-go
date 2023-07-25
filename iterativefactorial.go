package piscine

func IterativeFactorial(nb int) int {
	if nb <= 25 {

		if nb == 0 {
			return 1
		}
		if nb < 0 {
			return 0
		}

		res := 1

		for i := 1; i <= nb; i++ {
			res = res * i
		}
		return res
	}
	return 0
}
