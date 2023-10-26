package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	result := mergeTwoLists(list1, list2)
	fmt.Println(result)

	result = mergeTwoLists(nil, nil)
	fmt.Println(result)

	list3 := &ListNode{Val: 2}
	list4 := &ListNode{Val: 1}
	result = mergeTwoLists(list3, list4)
	fmt.Println(result)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	nodeList1 := list1
	nodeList2 := list2

	result := &ListNode{}
	current := result

	for nodeList2 != nil && nodeList1 != nil {
		if nodeList2.Val > nodeList1.Val {
			current.Next = &ListNode{Val: nodeList1.Val}
			nodeList1 = nodeList1.Next
		} else {
			current.Next = &ListNode{Val: nodeList2.Val}
			nodeList2 = nodeList2.Next
		}

		current = current.Next
	}

	if nodeList1 != nil {
		current.Next = nodeList1
	}
	if nodeList2 != nil {
		current.Next = nodeList2
	}

	return result.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	var data []string
	node := l
	for node != nil {
		data = append(data, strconv.Itoa(node.Val))
		node = node.Next
	}

	return string(strings.Join(data, ", "))
}
