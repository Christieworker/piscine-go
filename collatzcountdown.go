package piscine

// La fonction CollatzCountdown renvoie le nombre de pas
// nécessaires pour atteindre 1 en utilisant la conjecture de Collatz.
// Elle renvoie -1 si start est égal à 0 ou négatif.
func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	n := start
	steps := 0

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			steps++
		} else {
			n = 3*n + 1
			steps++
		}
	}
	return steps
}
