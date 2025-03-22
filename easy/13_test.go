package easy_test

import (
	"testing"

	"github.com/emidev98/leetcode/easy"
	"github.com/stretchr/testify/assert"
)

type testCase13 = struct {
	input          string
	expectedReturn int
}

func Test13(t *testing.T) {
	tests := []testCase13{
		{input: "III", expectedReturn: 3},
		{input: "LVIII", expectedReturn: 58},
		{input: "MCMXCIV", expectedReturn: 1994},
	}

	for _, test := range tests {
		res := easy.RomanToInt(test.input)
		assert.EqualValues(t, test.expectedReturn, res)
	}

}
