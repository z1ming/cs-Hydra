func combinationSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var temp []int
	helper(nums, target, 0, temp, &ans)
	return ans
}

func helper(nums []int, target, i int, temp []int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, temp...))
		return
	} else if target > 0 && i < len(nums) {
		next := getNext(nums, i)
		helper(nums, target, next, temp, ans)
		temp = append(temp, nums[i])
		helper(nums, target-nums[i], i+1, temp, ans)
		temp = temp[:len(temp)-1]
	}
}

func getNext(nums []int, index int) int {
	next := index
	for next < len(nums) && nums[next] == nums[index] {
		next++
	}
	return next
}
