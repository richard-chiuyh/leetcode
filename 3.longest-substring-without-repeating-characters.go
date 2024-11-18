/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	start, maxLen, currLen := 0, 0, 0
	charIndex := make(map[rune]int)
	for end, char := range s {
		if idx, ok := charIndex[char]; ok && start <= idx {
			start = idx + 1
			currLen = end - start + 1
		} else {
			currLen++
			if currLen > maxLen {
				maxLen = currLen
			}
		}
		charIndex[char] = end
	}
	return maxLen
}

// @lc code=end

