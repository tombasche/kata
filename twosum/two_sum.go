package twosum

func TwoSum(input []int, target int) []int {

	cache := make(map[int]int)
	for idx, val := range input {
		complement := target - val
		if cachedIdx, ok := cache[complement]; ok {
			return []int{cachedIdx, idx}
		}
		cache[val] = idx
	}
	return []int{}
}
