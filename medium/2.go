package medium

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stringifiedNumber1 := getListValues(l1, "")
	stringifiedNumber2 := getListValues(l2, "")

	number1, err := strconv.ParseInt(stringifiedNumber1, 10, 64)
	if err != nil {
		panic(err)
	}
	number2, err := strconv.ParseInt(stringifiedNumber2, 10, 64)
	if err != nil {
		panic(err)
	}

	totalSum := number1 + number2

	totalStringifiedSum := strconv.FormatInt(totalSum, 10)
	return toListNode(totalStringifiedSum, nil)
}

func toListNode(strSum string, res *ListNode) *ListNode {
	if len(strSum) == 0 {
		return res
	}

	firstElement := strSum[:1]
	firstNumber, err := strconv.Atoi(firstElement)
	if err != nil {
		panic(err)
	}

	if res == nil {
		return toListNode(strSum[1:], &ListNode{
			Val:  firstNumber,
			Next: nil,
		})
	}

	return toListNode(strSum[1:], &ListNode{
		Val:  firstNumber,
		Next: res,
	})
}

func getListValues(l *ListNode, stringifiedItems string) string {
	stringifiedItems = strconv.Itoa(l.Val) + stringifiedItems
	if l.Next == nil {
		return stringifiedItems
	}

	return getListValues(l.Next, stringifiedItems)
}
