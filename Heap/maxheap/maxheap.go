package maxheap

import "math"

//堆节点
type Heap struct {
	Size     int
	Capacity int
	Node     []interface{} //存储堆元素
}

//实现大顶堆
type MaxHeap Heap

//创建堆
func (h Heap) Init(m int) MaxHeap {
	heap := MaxHeap{
		Size:     0,
		Capacity: m,
		Node:     make([]interface{}, m+1),
	}
	//定义哨兵为最大值
	heap.Node[0] = math.MaxInt
	return heap
}

func (m *MaxHeap) IsFull() bool {
	return m.Size >= m.Capacity
}
func (m *MaxHeap) IsEmpty() bool {
	return m.Size == 0
}

//最大堆插入元素
func (m *MaxHeap) Insert(item interface{}) {
	if m.IsFull() {
		panic("MaxHeap Is Full")
	}
	i := m.Size + 1 //i指向插入后堆中最后一个元素
	//如果插入元素比父节点大，那就与父节点换位
	//i>1是因为Node[0]是哨兵位置，已经超出了堆的范围
	for ; m.Node[i/2].(int) < item.(int) /* && i > 1*/ ; i /= 2 {
		m.Node[i] = m.Node[i/2]
	}
	m.Node[i] = item
	m.Size++
}

// DeleteMax 最大堆删除元素,取出根节点(最大值)，然后删除堆的一个节点
func (m *MaxHeap) DeleteMax() interface{} {
	if m.IsEmpty() {
		panic("MaxEmpty Is Empty")
	}
	max := m.Node[1] //取出根节点最大值
	//用最大堆中最后一个元素从根节点开始向上过滤下层节点
	tmp := m.Node[m.Size-1]
	var parent, child int
	for parent = 1; parent*2 <= m.Size; parent = child {
		child = parent * 2
		if child != m.Size && m.Node[child].(int) < m.Node[child+1].(int) {
			child++ //child指向左右子树较大者
		}
		if tmp.(int) >= m.Node[child].(int) {
			break
		} else {
			//移动tmp元素到下一层
			m.Node[parent] = m.Node[child]
		}
	}
	m.Node[parent] = tmp
	m.Size--
	return max
}
