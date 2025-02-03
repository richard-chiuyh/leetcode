/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1,
	}
	ans := 0
	for i := 1; i < len(s); i++ {
		x, y := romanMap[s[i-1]], romanMap[s[i]]
		if x < y {
			ans -= x
		} else {
			ans += x
		}
	}
	return ans + romanMap[s[len(s)-1]]
}

// @lc code=end
