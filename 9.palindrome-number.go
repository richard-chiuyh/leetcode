/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	temp := x
	for temp != 0 {
		reversed = reversed*10 + temp%10
		temp /= 10
	}
	fmt.Println(reversed)
	return (reversed == x)
}

// @lc code=end

