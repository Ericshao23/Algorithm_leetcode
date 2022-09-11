// https://codeforces.com/contest/1728/problem/C
package main

import "fmt"

type Element interface{}

type Queue interface {
	Push(e Element)
	Pop() Element
	Top() Element
	Size() int
	IsEmpty() bool
}

type sliceEntry struct {
	element []Element
}

func NewQueue() *sliceEntry {
	return &sliceEntry{}
}

// 向队列中添加元素
func (entry *sliceEntry) Push(e Element) {
	entry.element = append(entry.element, e)
}

// 移除队列中最前面元素
func (entry *sliceEntry) Pop() Element {
	if entry.IsEmpty() {
		fmt.Println("queue is empty")
		return nil
	}
	firstElement := entry.element[0]
	entry.element = entry.element[1:]
	return firstElement
}

// 队列长度
func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

// 队列首个元素
func (entry *sliceEntry) Top() Element {
	if entry.IsEmpty() {
		return nil
	}
	firstElement := entry.element[0]
	//entry.element = entry.element[1:]
	return firstElement
}

// 队列是否为空
func (entry *sliceEntry) IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		b := make([]int, n)
		qa := NewQueue()
		for j := 0; j < n; j++ {
			fmt.Scan(&a[j])
			qa.Push(a[j])
		}
		qb := NewQueue()
		for k := 0; k < n; k++ {
			fmt.Scan(&b[k])
			qb.Push(b[k])
		}
		ans := 0
		for qa.IsEmpty() != true {
			if qa.Top() == qb.Top() {
				qa.Pop()
				qb.Pop()
				continue
			}
			ans += 1
			// 存在部分问题待解决
			if qa.Top() > qb.Top() {
				qa.Push(qa.Top())
				qa.Pop()
			} else {
				qb.Push(qb.Pop())
				qb.Pop()
			}
		}
		fmt.Println(ans)
	}
}
