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
	left, right, max := 0, len(height)-1, 0
	for left < right {
		tmp := (right - left) * min(height[left], height[right])
		if tmp > max {
			max = tmp
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

// @lc code=end

