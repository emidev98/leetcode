package easy_test

import (
	"testing"

	"github.com/emidev98/leetcode/easy"
	"github.com/stretchr/testify/assert"
)

type testCase1 = struct {
	input          []int
	target         int
	expectedReturn []int
}

func Test1(t *testing.T) {
	tests := []testCase1{
		{input: []int{2, 7, 11, 15}, target: 9, expectedReturn: []int{0, 1}},
		{input: []int{3, 2, 4}, target: 6, expectedReturn: []int{1, 2}},
		{input: []int{3, 3}, target: 6, expectedReturn: []int{0, 1}},
		{input: []int{3, 2, 3}, target: 6, expectedReturn: []int{0, 2}},
		{input: []int{4, 0, 2, 0}, target: 6, expectedReturn: []int{0, 2}},
	}

	for _, test := range tests {
		res := easy.TwoSum(test.input, test.target)
		assert.EqualValues(t, test.expectedReturn, res)
	}

}
