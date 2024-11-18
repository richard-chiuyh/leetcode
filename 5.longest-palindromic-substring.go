/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	var longest string
	for i := range s {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		if len(l1) > len(longest) {
			longest = l1
		}
		if len(l2) > len(longest) {
			longest = l2
		}
	}
	return longest
}

func expandAroundCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

// @lc code=end

