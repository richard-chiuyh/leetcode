/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	value := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ans := ""
	for i, v := range value {
		for num >= v {
			num = num - v
			ans = ans + symbol[i]
		}
	}
	return ans
}

// @lc code=end
