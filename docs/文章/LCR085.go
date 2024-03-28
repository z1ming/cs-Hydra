func generateParenthesis(n int) []string {
    var ans []string 
    helper(n, n, "", &ans)
    return ans
}

func helper(left, right int, temp string, ans *[]string) {
    if left == 0 && right == 0 {
        *ans = append(*ans, temp)
        return
    }
    if left > 0 {
        helper(left - 1, right, temp + "(", ans)
    }

    if left < right {
        helper(left, right - 1, temp + ")", ans)
    }
    
}

