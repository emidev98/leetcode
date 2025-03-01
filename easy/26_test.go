package easy_test

import (
	"testing"

	"github.com/emidev98/leetcode/easy"
	"github.com/stretchr/testify/assert"
)

type testCase26 = struct {
	input          []int
	expectedValue  []int
	expectedReturn int
}

func Test26(t *testing.T) {
	tests := []testCase26{
		{input: []int{1, 1, 2}, expectedValue: []int{1, 2}, expectedReturn: 2},
		{input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expectedValue: []int{0, 1, 2, 3, 4}, expectedReturn: 5},
	}

	for _, test := range tests {
		res := easy.RemoveDuplicates(test.input)
		assert.EqualValues(t, test.expectedReturn, res)
	}

}
