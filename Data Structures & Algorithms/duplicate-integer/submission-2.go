func hasDuplicate(nums []int) bool {
    memory := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := memory[n]; ok {
			return true
		}
		memory[n] = struct{}{}
	} 
	return false
}
