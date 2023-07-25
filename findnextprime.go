package piscine

func Prime(a int) bool {
	if a <= 1 {
		return false
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for {
		if Prime(nb) {
			return nb
		}
		nb++
	}
}
