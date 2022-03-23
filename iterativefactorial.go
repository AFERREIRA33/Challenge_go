package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb >= 0 && nb < 100 {
		for i := 1; i <= nb; i++ {
			if result > 0 {
				result = result * i
			}
		}
	} else {
		result = 0
	}
	return result
}
