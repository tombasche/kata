package runningsum

func RunningSum(input []int) (result []int) {

	lastElem := input[0]
	result = []int{lastElem}
	for _, val := range input[1:] {
		result = append(result, val+lastElem)
		lastElem += val
	}
	return
}
