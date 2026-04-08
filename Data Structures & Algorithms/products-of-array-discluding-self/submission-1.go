func product(nums... int) int {
	res := 1
	for _, n := range nums {
		res *= n
	}

	return res
}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	result := make([]int, l, l)

	for i := range l {
		tmp := make([]int, 0, l-1)
		tmp = append(tmp, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		result[i] = product(tmp...)
	}

	return result
}
