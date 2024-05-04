type HashSet map[int]bool

func permuteUnique(nums []int) [][]int {
    var ans [][]int
    helper(nums, 0, &ans)
    return ans
}

func helper(nums []int, i int, ans *[][]int) {
    if i == len(nums) {
        permutation := make([]int, len(nums))
        copy(permutation, nums)
        *ans = append(*ans, permutation)
    } else {
        set := make(HashSet)
        for j := i; j < len(nums); j++ {
            if !set[nums[j]] {
                set[nums[j]] = true
                swap(nums, i, j)
                helper(nums, i + 1, ans)
                swap(nums, i, j)
            }
        }
    }
}

func swap(nums []int, i, j int) {
    if i != j {
        nums[i], nums[j] = nums[j], nums[i]
    }
    
}
