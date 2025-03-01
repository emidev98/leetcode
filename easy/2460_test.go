package easy_test

import (
	"testing"

	"github.com/emidev98/leetcode/easy"
	"github.com/stretchr/testify/assert"
)

type testCase2460 = struct {
	input         []int
	expectedValue []int
}

func Test2460(t *testing.T) {
	tests := []testCase2460{
		{input: []int{1, 2, 2, 1, 1, 0}, expectedValue: []int{1, 4, 2, 0, 0, 0}},
		{input: []int{0, 1}, expectedValue: []int{1, 0}},
	}

	for _, test := range tests {
		value := easy.ApplyOperations2460(test.input)
		assert.EqualValues(t, test.expectedValue, value)
	}

}
