package utils

func BinaryConvertor(number int, bits int) []int {
	result := make([]int, 0)
	for number > 0 {
		result = append(result, number%2)
		number /= 2
	}
	for i := len(result) - 1; len(result) != bits; i++ {
		result = append(result, 0)
	}
	return result
}
