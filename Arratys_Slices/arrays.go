package arrays

func Sum(array []int) int {
	sum := 0
	for _, el := range array {
		sum += el
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	result := make([]int, len(arrays))
	for i, arr := range arrays {
		result[i] = Sum(arr)
	}
	return result
}

func SumAllTails(arrays ...[]int) []int {
	var result []int

	for _, arr := range arrays {
		if len(arr) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(arr[1:]))
		}
	}

	return result
}
