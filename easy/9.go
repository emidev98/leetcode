package easy

// https://leetcode.com/problems/palindrome-number/solutions/6538773/9-palindrome-number-by-emidev98-qjlo/
func IsPalindrome(x int) bool {
	// if the number is negative it does not read the same
	// way reversed
	if x < 0 {
		return false
	}

	// alway start with 0
	reversed := 0
	// temp is 51
	temp := x

	for temp != 0 {
		// [iter1] (0 = 0 * 10)
		// [iter2] (0 = 1 * 10)
		reversedMultByTen := (reversed * 10)

		// [iter1] (1 = 51 % 10)
		// [iter2] (5 = 5 % 10)
		tempModuleTen := (temp % 10)

		// [iter1] (1 = 0 + 1)
		// [iter2] (1 = 10 + 5)
		reversed = reversedMultByTen + tempModuleTen

		// [iter1] (5 = 51 / 10)
		// [iter1] (1 = 51 / 10)
		temp = temp / 10
	}

	return x == reversed
}
