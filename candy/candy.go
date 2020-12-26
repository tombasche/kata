package candy

func DistributeCandies(input []int, extraCandies int) []bool {

	var result []bool

	most := max(input)
	for _, c := range input {
		total := c + extraCandies
		result = append(result, total >= most)
	}
	return result
}

func max(input []int) int {
	m := 0
	for _, val := range input {
		if val > m {
			m = val
		}
	}
	return m
}
