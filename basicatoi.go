package piscine

func BasicAtoi(s string) int {
	array1 := []rune(s)
	num1 := 0
	num2 := 0
	for _, val := range array1 {
		for i := '0'; i < val; i++ {
			num1++
		}
		num2 = num1 + num2*10
		num1 = 0
	}
	return num2
}
