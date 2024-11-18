/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	dp, max, low := make([]int, len(height)), 0, math.MaxInt64
	for i := 0; i < len(height)-1; i++ {
		tmp := 0
		if height[i] > height[i+1] {
			tmp = height[i+1]
		} else {
			tmp = height[i]
		}
		
	}
}

// @lc code=end

