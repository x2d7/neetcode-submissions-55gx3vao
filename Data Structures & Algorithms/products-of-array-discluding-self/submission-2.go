func productExceptSelf(nums []int) []int {
	l := len(nums)
	result := make([]int, l, l)

	up := make([]int, l, l)
	down := make([]int, l, l)

	for i, _ := range nums {
		j := l - i - 1

		up[i] = 1
		down[i] = 1
		if i > 0 {
			up[i] *= up[i-1] * nums[i-1]
			down[i] *= down[i-1] * nums[j+1]
		}
	}

	for i, upV := range up {
		downV := down[l - i - 1]
		result[i] = downV * upV
	}

	return result
}
