func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    numSet := make(map[int]struct{}, len(nums))
    for _, num := range nums {
        numSet[num] = struct{}{}
    }

    maxLen := 0
    for num := range numSet {
        if _, exists := numSet[num-1]; !exists {
            currentNum, currentLen := num, 1
            
            for {
                if _, exists := numSet[currentNum+1]; !exists {
                    break
                }
                currentNum++
                currentLen++
            }
            
            if currentLen > maxLen {
                maxLen = currentLen
            }
        }
    }
    return maxLen
}