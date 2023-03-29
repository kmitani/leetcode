package main

import "fmt"

// https://qiita.com/miyyuk/items/af9d3352f89f11365dc4
// https://medium.com/since-i-want-to-start-blog-that-looks-like-men-do/%E5%88%9D%E5%BF%83%E8%80%85%E3%81%AB%E9%80%81%E3%82%8A%E3%81%9F%E3%81%84interface%E3%81%AE%E4%BD%BF%E3%81%84%E6%96%B9-golang-48eba361c3b4
// interfaceを使って同じ関数群を持つものをまとめることができる。まとめることで別の関数の再利用性が高めることができる。

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	val, carry := addVal(l1, l2, carry)
	output := &ListNode{Val: val, Next: nil}
	for l1.Next != nil || l2.Next != nil {
		l1 = moveNext(l1)
		l2 = moveNext(l2)
		val, carry = addVal(l1, l2, carry)
		output = insert(output, val)
	}
	if carry == 1 {
		output = insert(output, carry)
	}

	return output
}

func moveNext(l *ListNode) *ListNode {
	if l.Next != nil {
		l = l.Next
	} else {
		l.Val = 0
		l.Next = nil
	}
	return l
}

func addVal(l1, l2 *ListNode, carry int) (val, carryNext int) {
	val = (l1.Val + l2.Val + carry) % 10
	carryNext = (l1.Val + l2.Val + carry) / 10
	return val, carryNext
}

func insert(l *ListNode, val int) *ListNode {
	new := &ListNode{Val: val, Next: nil}
	if l.Next == nil {
		l.Next = new
	} else {
		next := l.Next
		for next.Next != nil {
			next = next.Next
		}
		next.Next = new
	}
	return l
}

func (l *ListNode) show() {
	next := l.Next
	for next.Next != nil {
		fmt.Print(next.Val, "->")
		next = next.Next
	}
	fmt.Println(next.Val)
}

func main() {
	l1_nums := []int{2, 4, 3}
	l1 := &ListNode{}
	for _, v := range l1_nums {
		l1 = insert(l1, v)
	}
	l1.show()
	l2_nums := []int{5, 6, 4}
	l2 := &ListNode{}
	for _, v := range l2_nums {
		l2 = insert(l2, v)
	}
	l2.show()

	output := addTwoNumbers(l1, l2)
	// Explanation: 342 + 465 = 807.
	// ListNodeを表示するときは逆順になる (2->4->3, ...)
	output.show()
}
