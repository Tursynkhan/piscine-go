package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	i := 0
	var nums []int
	number_n := n
	toolong := false

	if n < -9223372036854775807 {
		n /= 10
		toolong = true
	}

	if n < 0 {
		n *= -1
	}

	for i >= 0 {
		nums = append(nums, n%10)
		if n < 10 {
			break
		}
		n = n / 10
	}

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	if toolong {
		nums = append(nums, 8)
	}
	if number_n < 0 {
		z01.PrintRune('-')
	}
	for i := range nums {
		z01.PrintRune(rune(nums[i] + 48))
	}
}
