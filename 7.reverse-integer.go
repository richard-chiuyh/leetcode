/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}
	return reversed
}

// @lc code=end

