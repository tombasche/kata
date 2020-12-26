package shuffle

func Shuffle(input []int, n int) []int {

	mid := len(input) / 2
	l1 := input[:mid]
	l2 := input[mid:]
	i := 0

	var result []int
	for i < n {
		result = append(result, l1[i], l2[i])
		i++
	}
	return result
}
