package main

func plusOne(digits []int) []int {
	for i:= len(digits)-1; i >=0;i--{
		digits[i] = digits[i] + 1
		if digits[i] != 10{
			break
		}
		digits[i] = 0

		if i == 0{
			result := make([]int,1)
			result[0] = 1
			digits = append(result, digits...)
		}
	}
	return digits
}
