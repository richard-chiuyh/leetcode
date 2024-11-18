/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := numMap[complement]; found {
			return []int{index, i}
		}

		numMap[num] = i
	}

	return nil
}

// @lc code=end

