package main

import (
	"fmt"
)

//哈夫曼树(最优二叉树)-带权路径值最小的树
//通过小顶堆来每次确定找到最小的两个数
/*
	哈夫曼树没有度为1的节点
	n个叶子节点的哈夫曼树共有2n-1个节点
*/

//哈夫曼树结构体
type HuffmanTree struct {
	Value     interface{} //字符或是数值
	Weight    uint        //权值
	LeftTree  *HuffmanTree
	RightTree *HuffmanTree
	Parent    *HuffmanTree //父节点
}

type MinHeap struct {
	Size int
	Heap []*HuffmanTree //存放元素的数组
}

// Init 初始化哈夫曼树
func Init() *MinHeap {
	T := &MinHeap{
		Size: 0,
		Heap: make([]*HuffmanTree, 1),
	}
	T.Heap[0] = &HuffmanTree{}
	return T
}

// Insert 插入节点
func (h *MinHeap) Insert(item *HuffmanTree) {
	h.Size++
	i := h.Size
	h.Heap = append(h.Heap, &HuffmanTree{})
	for h.Heap[i/2].Weight > item.Weight {
		h.Heap[i] = h.Heap[i/2]
		i = i / 2
	}
	h.Heap[i] = item
}

//判断堆是否为空
func (h *MinHeap) IsEmpty() bool {
	return h.Size == 0
}

//最小堆的删除
func (h *MinHeap) Delete() *HuffmanTree {
	if h.IsEmpty() {
		return nil
	}
	var parent, child int
	minItem := h.Heap[1]
	for parent = 1; parent*2 <= h.Size; parent = child {
		child = parent * 2
		if child != h.Size && h.Heap[child].Weight > h.Heap[child+1].Weight {
			child++
		}
		if h.Heap[h.Size].Weight <= h.Heap[child].Weight {
			break
		} else {
			h.Heap[parent] = h.Heap[child]
		}
	}
	h.Heap[parent] = h.Heap[h.Size]
	h.Size--
	return minItem
}

//获取哈夫曼树
func (h *MinHeap) Huffman() *HuffmanTree {
	T := &HuffmanTree{}
	for i := 1; i < h.Size; i++ {
		T.LeftTree = h.Delete()                           //从小顶堆中删除一个节点，作为T的左子结点
		T.RightTree = h.Delete()                          //从小顶堆中删除一个节点，作为T的右子节点
		T.Weight = T.LeftTree.Weight + T.RightTree.Weight //计算新权值
		h.Insert(T)                                       //将新T插入到小顶堆
	}
	T = h.Delete()
	return T
}

func (h *HuffmanTree) Traversal() {
	if h != nil {
		fmt.Printf("%v\t", h.Weight)
		h.LeftTree.Traversal()
		h.RightTree.Traversal()
	}
}

func main() {
	h := Init()
	h.Insert(&HuffmanTree{Weight: 6})
	h.Insert(&HuffmanTree{Weight: 1})
	h.Insert(&HuffmanTree{Weight: 8})
	h.Insert(&HuffmanTree{Weight: 7})
	h.Insert(&HuffmanTree{Weight: 5})
	h.Insert(&HuffmanTree{Weight: 3})
	h.Insert(&HuffmanTree{Weight: 2})
	huffman := h.Huffman()
	huffman.Traversal()
}
