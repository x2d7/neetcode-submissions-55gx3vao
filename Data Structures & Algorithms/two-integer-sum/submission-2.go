func twoSum(nums []int, target int) []int {
    result := make([]int, 2, 2)

	passOn := make(map[int]int)
	for i1, v := range nums {
		if i2, ok := passOn[v]; ok {
			result[1] = i1
			result[0] = i2
			return result
		}
		passOn[target-v] = i1
	}
	return result
}
