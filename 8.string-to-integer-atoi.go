/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	ans, isNeg, hasNum := 0, false, false
	for _, char := range s {
		if char == ' ' && !hasNum {
			continue
		} else if char == '-' && !hasNum {
			hasNum = true
			isNeg = true
		} else if char == '+' && !hasNum {
			hasNum = true
			isNeg = false
		} else if char >= '0' && char <= '9' {
			hasNum = true
			ans = ans*10 + int(char-'0')
			if ans > math.MaxInt32 && !isNeg {
				return math.MaxInt32
			} else if -ans < math.MinInt32 && isNeg {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	if isNeg {
		ans = -ans
	}

	return ans
}

// @lc code=end

