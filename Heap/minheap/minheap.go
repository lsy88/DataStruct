package minheap

import "math"

//堆节点
type Heap struct {
	Size     int
	Capacity int
	Node     []interface{} //存储堆元素
}

//实现小顶堆
type MinHeap Heap

//创建堆
func (h Heap) Init(n int) MinHeap {
	heap := MinHeap{
		Size:     0,
		Capacity: n,
		Node:     make([]interface{}, n+1),
	}
	//定义哨兵为最小值
	heap.Node[0] = math.MinInt
	return heap
}

func (m *MinHeap) IsFull() bool {
	return m.Size >= m.Capacity
}

func (m *MinHeap) IsEmpty() bool {
	return m.Size == 0
}

//最小堆插入元素
func (m *MinHeap) Insert(item interface{}) {
	if m.IsFull() {
		panic("MinHeap Is Full")
	}
	m.Size++
	i := m.Size //i指向插入后堆中最后一个元素
	//如果插入元素比父节点小，那就与父节点换位
	//i>1是因为Node[0]是哨兵位置，已经超出了堆的范围
	for ; m.Node[i/2].(int) > item.(int) /* && i > 1*/ ; i /= 2 {
		m.Node[i] = m.Node[i/2]
	}
	m.Node[i] = item
}

// DeleteMin 最小堆删除元素,取出根节点(最大值)，然后删除堆的一个节点
func (m *MinHeap) DeleteMin() interface{} {
	if m.IsEmpty() {
		panic("MinHeap Is Empty")
	}
	min := m.Node[1] //取出根节点最大值
	//用最大堆中最后一个元素从根节点开始向上过滤下层节点
	tmp := m.Node[m.Size-1]
	var parent, child int
	for parent = 1; parent*2 <= m.Size; parent = child {
		child = parent * 2
		if child != m.Size && m.Node[child].(int) > m.Node[child+1].(int) {
			child++ //child指向左右子树较大者
		}
		if tmp.(int) <= m.Node[child].(int) {
			break
		} else {
			//移动tmp元素到下一层
			m.Node[parent] = m.Node[child]
		}
	}
	m.Node[parent] = tmp
	m.Size--
	return min
}
