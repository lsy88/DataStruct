package main

import (
	"container/heap"
)

//优先级队列要实现Len(), Less(), Swap()三个方法

//struct队列
type PriorityQueue []*Node
type Node struct {
	Priority int    //队列的优先级
	Value    string //值
	index    int
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

//按照优先级排序
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

//交换次序
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Node, value string, priority int) {
	item.Value = value
	item.Priority = priority
	//Fix在索引i处的元素更新后重新建立堆排序
	heap.Fix(pq, item.index)
}

//数组实现
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
