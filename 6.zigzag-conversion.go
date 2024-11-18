/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	index, step := 0, 1
	for _, char := range s {
		rows[index] += string(char)
		if index == 0 {
			step = 1
		} else if index == numRows-1 {
			step = -1
		}
		index += step
	}

	ans := ""
	for _, row := range rows {
		ans += row
	}
	return ans
}

// @lc code=end

