package wealth

func Wealthiest(input [][]int) (sum int) {

	for _, arr := range input {
		t := sumSlice(arr)
		if t > sum {
			sum = t
		}
	}

	return
}

func sumSlice(input []int) int {
	var total int
	for _, i := range input {
		total += i
	}
	return total
}
