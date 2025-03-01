package easy

// https://leetcode.com/problems/apply-operations-to-an-array/solutions/6481672/2460-apply-operations-to-an-array-by-emi-bx6m/
func ApplyOperations2460(nums []int) []int {
	numsLenght := len(nums)

	for currentIndex := 0; currentIndex < numsLenght; currentIndex++ {
		currentValue := nums[currentIndex]

		nextIndex := currentIndex + 1
		if nextIndex < numsLenght {

			nextValue := nums[nextIndex]
			if currentValue == nextValue {
				nums[currentIndex] = currentValue * 2
				nums[nextIndex] = 0
			}
		}

	}

	numsGreaterThanZero := []int{}
	zeros := []int{}
	for i := 0; i < numsLenght; i++ {
		if nums[i] == 0 {
			zeros = append(zeros, 0)
		} else {
			numsGreaterThanZero = append(numsGreaterThanZero, nums[i])
		}
	}

	return append(numsGreaterThanZero, zeros...)
}
