package twosum

func TwoSum(input []int, target int) []int {

	for compareIdx := range input {
		firstElem := input[compareIdx]
		for idx, val := range input {
			if idx == compareIdx {
				continue
			}
			if firstElem+val == target {
				return []int{compareIdx, idx}
			}
		}
	}

	return []int{}
}
