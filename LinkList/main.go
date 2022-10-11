package main

type linklist interface {
	Reverse(k int)
	Length() int
	IsEmpty() bool
	InsertFromHead(data interface{}) *LinkList
	InsertFromTail(data interface{}) *LinkList
	InsertByIndex(index int, data interface{}) *LinkList
	Delete(data interface{}) *LinkList
	DeleteAtIndex(index int) bool
	Contain(data interface{}) bool
}

type Node struct {
	Value interface{}
	Next  *Node
}

// LinkList 单链表
type LinkList struct {
	Head *Node //头结点
	Tail *Node //尾结点
	Size int
}

// Reserve 反转链表 k为将多少个结点反转
func (l *LinkList) Reserve(k int) {
	cnt := 1
	newNode := l.Head //newNode为第一个结点
	oldNode := newNode.Next
	for cnt < k {
		tmp := oldNode.Next    //tmp中间变量记录oldNode所指向的下一个结点
		oldNode.Next = newNode //实现反转，将oldNode的下一个结点置为newNode
		newNode = oldNode
		oldNode = tmp
		cnt++
	}
}

// Length 获取链表长度
func (l *LinkList) Length() int {
	length := 0
	node := l.Head
	for node != nil {
		length++
		node = node.Next
	}
	if l.Size != length {
		l.Size = length
	}
	return length
}

// IsEmpty 是否为空
func (l *LinkList) IsEmpty() bool {
	if l.Head == nil {
		return true
	} else {
		return false
	}
}

// InsertFromHead 从链表头部添加元素
func (l *LinkList) InsertFromHead(data interface{}) *LinkList {
	node := &Node{Value: data}
	node.Next = l.Head
	l.Head = node
	l.Size++ //链表数量加1
	return l
}

// InsertFromTail 从链表尾部添加元素
func (l *LinkList) InsertFromTail(data interface{}) *LinkList {
	node := &Node{Value: data}
	if l.IsEmpty() { //如果链表为空，那么直接将元素作为头结点
		l.Head = node
		return l
	}
	l.Tail.Next = node
	l.Tail = node
	l.Size++
	return l
}

// InsertByIndex 根据位置添加元素
func (l *LinkList) InsertByIndex(index int, data interface{}) *LinkList {
	if index < 0 { //如果index小于0，在头部进行插入
		l.InsertFromHead(data)
	} else if index > l.Length() { //index大于链表长度，在尾部进行插入
		l.InsertFromTail(data)
	} else {
		pre := l.Head
		cnt := 0
		for cnt < index-1 {
			pre = pre.Next
			cnt++
		}
		node := &Node{Value: data}
		node.Next = pre.Next
		pre.Next = node
		l.Size++
	}
	return l
}

// Delete 删除指定元素
func (l *LinkList) Delete(data interface{}) *LinkList {
	pre := l.Head
	if pre.Value == data {
		l.Head = pre.Next
	} else {
		//删除的不是头结点
		for pre.Next != nil {
			if pre.Next.Value == data { //删除节点
				pre.Next = pre.Next.Next
				l.Size--
			} else {
				pre = pre.Next
			}
		}
	}
	return l
}

// DeleteAtIndex 删除指定位置的元素
func (l *LinkList) DeleteAtIndex(index int) bool {
	pre := l.Head
	if index <= 0 {
		//删除头节点
		l.Head = pre.Next
	} else if index > l.Length() {
		//超出链表长度
		return false
	} else {
		cnt := 0
		for cnt != (index-1) && pre.Next != nil {
			cnt++
			pre.Next = pre.Next.Next
		}
	}
	return true
}

// Contain 查看是否包含某个元素
func (l *LinkList) Contain(data interface{}) bool {
	pre := l.Head
	for pre != nil {
		if pre.Value == data {
			return true
		}
		pre = pre.Next
	}
	return false
}
