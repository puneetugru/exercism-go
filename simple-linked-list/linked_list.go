// API:
//
// type Element struct {
//  data int
//  next *Element
// }
//
// type List struct {
//  head *Element
//  size int
// }
//
// func New([]int) *List
// func (*List) Size() int
// func (*List) Push(int)
// func (*List) Pop() (int, error)
// func (*List) Array() []int
// func (*List) Reverse() *List

package linkedlist

import "errors"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(value []int) *List {
	var result List
	for _, elem := range value {
		result.Push(elem)
	}
	return &result
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Push(value int) {
	newNode := Element{value, nil}
	if list.head == nil {
		list.head = &newNode
	} else {
		node := list.head
		for node.next != nil {
			node = node.next
		}
		node.next = &newNode
	}
	list.size++
}

func (list *List) Pop() (int, error) {
	var value int
	if list.size == 0 {
		return 0, errors.New("Cannot pop from empty list")
	} else if list.size == 1 {
		value = list.head.data
		list.head = nil
	} else {
		node := list.head
		for node.next.next != nil {
			node = node.next
		}
		value = node.next.data
		node.next = nil
	}
	list.size--
	return value, nil
}

func (list *List) Array() []int {
	result := []int{}
	for node := list.head; node != nil; node = node.next {
		result = append(result, node.data)
	}
	return result
}

func (list *List) Reverse() *List {
	array, size := list.Array(), list.size
	for i := 0; i < size/2; i++ {
		array[i], array[size-i-1] = array[size-i-1], array[i]
	}
	return New(array)
}
