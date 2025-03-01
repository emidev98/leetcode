package easy

// https://leetcode.com/problems/two-sum/solutions/6483074/1-two-sum-by-emidev98-rfxu/
func TwoSum(nums []int, target int) []int {
	pairs := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		currentNumber := nums[i]
		expected := target - currentNumber

		counterpartIndex, ok := pairs[expected]

		if ok {
			return []int{counterpartIndex, i}
		} else {
			pairs[currentNumber] = i
		}
	}

	return []int{0}
}
