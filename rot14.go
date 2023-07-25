package piscine

func rot14(letter rune) rune {
	if letter >= 'A' && letter < 'M' || letter >= 'a' && letter < 'm' {
		return letter + 14
	}
	if letter >= 'M' && letter <= 'Z' || letter >= 'm' && letter <= 'z' {
		return letter - 12
	}
	return letter
}

func Rot14(str string) string {
	resultat := ""
	for _, res := range str {
		resultat += string(rot14(res))
	}
	return resultat
}
