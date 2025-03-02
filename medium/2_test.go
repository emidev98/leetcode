package medium_test

import (
	"testing"

	"github.com/emidev98/leetcode/medium"
	"github.com/stretchr/testify/assert"
)

type testCase2 = struct {
	inputList      *medium.ListNode
	inputList1     *medium.ListNode
	expectedReturn *medium.ListNode
}

func Test1(t *testing.T) {
	tests := []testCase2{{
		inputList: &medium.ListNode{
			Val: 2,
			Next: &medium.ListNode{
				Val: 4,
				Next: &medium.ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
		inputList1: &medium.ListNode{
			Val: 5,
			Next: &medium.ListNode{
				Val: 6,
				Next: &medium.ListNode{
					Val: 4,
					Next: &medium.ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
		expectedReturn: &medium.ListNode{
			Val: 7,
			Next: &medium.ListNode{
				Val: 0,
				Next: &medium.ListNode{
					Val: 4,
					Next: &medium.ListNode{
						Val: 0,
						Next: &medium.ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	},
		{
			inputList: &medium.ListNode{
				Val: 2,
				Next: &medium.ListNode{
					Val: 4,
					Next: &medium.ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			inputList1: &medium.ListNode{
				Val: 5,
				Next: &medium.ListNode{
					Val: 6,
					Next: &medium.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expectedReturn: &medium.ListNode{
				Val: 7,
				Next: &medium.ListNode{
					Val: 0,
					Next: &medium.ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			inputList: &medium.ListNode{
				Val:  0,
				Next: nil,
			},
			inputList1: &medium.ListNode{
				Val:  0,
				Next: nil,
			},
			expectedReturn: &medium.ListNode{
				Val:  0,
				Next: nil,
			},
		},
		{
			inputList: &medium.ListNode{
				Val: 9,
				Next: &medium.ListNode{
					Val: 9,
					Next: &medium.ListNode{
						Val: 9,
						Next: &medium.ListNode{
							Val: 9,
							Next: &medium.ListNode{
								Val: 9,
								Next: &medium.ListNode{
									Val: 9,
									Next: &medium.ListNode{
										Val: 9,
										Next: &medium.ListNode{
											Val:  9,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
			inputList1: &medium.ListNode{
				Val: 9,
				Next: &medium.ListNode{
					Val: 9,
					Next: &medium.ListNode{
						Val: 9,
						Next: &medium.ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
			},
			expectedReturn: &medium.ListNode{
				Val: 8,
				Next: &medium.ListNode{
					Val: 9,
					Next: &medium.ListNode{
						Val: 9,
						Next: &medium.ListNode{
							Val: 9,
							Next: &medium.ListNode{
								Val: 0,
								Next: &medium.ListNode{
									Val: 0,
									Next: &medium.ListNode{
										Val: 0,
										Next: &medium.ListNode{
											Val: 0,
											Next: &medium.ListNode{
												Val:  1,
												Next: nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		res := medium.AddTwoNumbers(test.inputList, test.inputList1)
		assert.EqualValues(t, res, test.expectedReturn)
	}

}
