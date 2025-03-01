package easy

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/solutions/6481667/26-remove-duplicates-from-sorted-array-b-r9j6/
func RemoveDuplicates(nums []int) int {
	lastCheckedElement := nums[0]
	uniqueNums := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if lastCheckedElement != nums[i] {
			uniqueNums = append(uniqueNums, nums[i])
		}
		lastCheckedElement = nums[i]

	}
	copy(nums, uniqueNums)
	return len(uniqueNums)
}
