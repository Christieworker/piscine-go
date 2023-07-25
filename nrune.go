package piscine

func NRune(s string, n int) rune {
	runeArray := []rune(s)
	arrayCount := 0
	for a := range s {
		arrayCount = a + 1
	}
	if n <= arrayCount && n-1 >= 0 {
		return runeArray[n-1]
	} else {
		return 0
	}
}
