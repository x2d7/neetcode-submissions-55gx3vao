func topKFrequent(nums []int, k int) []int {
	result := make([]int, k, k)
	m := make(map[int]int)

	for _, v := range nums {
		m[v] += 1
	}

	l := len(m)

	type kv struct {
		k, v int
	}
	kvlist := make([]kv, 0, l)
	for k, v := range m {
		kvlist = append(kvlist, kv{k, v})
	}

	sort.Slice(kvlist, func(i, j int) bool {
		return kvlist[i].v < kvlist[j].v
	})

	kvlist = kvlist[l-k:l]
	for i, v := range kvlist {
		result[i] = v.k
	}

	return result
}
