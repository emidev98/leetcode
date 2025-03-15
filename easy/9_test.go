package easy_test

import (
	"testing"

	"github.com/emidev98/leetcode/easy"
	"github.com/stretchr/testify/assert"
)

type testCase9 = struct {
	input          int
	expectedReturn bool
}

func Test9(t *testing.T) {
	tests := []testCase9{
		{input: -121, expectedReturn: false},
		{input: 51, expectedReturn: false},
		{input: 120, expectedReturn: false},
		{input: 121, expectedReturn: true},
		{input: -10, expectedReturn: false},
	}

	for _, test := range tests {
		res := easy.IsPalindrome(test.input)
		assert.EqualValues(t, test.expectedReturn, res)
	}

}
