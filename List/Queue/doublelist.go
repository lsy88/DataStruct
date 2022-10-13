package Queue

import "sync"

//双端列表，双端队列
type DoubleList struct {
	Head *ListNode //指向链表头部
	Tail *ListNode //指向链表尾部
	Len  int       //链表长度
	lock sync.Mutex
}

type ListNode struct {
	pre   *ListNode // 前驱节点
	next  *ListNode // 后驱节点
	value string    // 值
}

// 获取节点值
func (node *ListNode) GetValue() string {
	return node.value
}

// 获取节点前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// 获取节点后驱节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// 是否存在后驱节点
func (node *ListNode) HasNext() bool {
	return node.pre != nil
}

// 是否存在前驱节点
func (node *ListNode) HasPre() bool {
	return node.next != nil
}

// 是否为空节点
func (node *ListNode) IsNil() bool {
	return node == nil
}

//添加节点到链表头部的第N个元素之前，N=0成为新节点的头部
func (list *DoubleList) AddNodeFromHead(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()
	
	//索引长度超过列表长度
	if n > list.Len {
		panic("index out")
	}
	
	//先找出头部
	node := list.Head
	
	//往后遍历拿到第n+1个位置元素
	for i := 1; i <= n; i++ {
		node = node.next
	}
	
	newNode := &ListNode{value: v}
	
	//如果定位到的节点为空，表示链表为空。将新节点设置成心头不和欣慰不
	if node.IsNil() {
		list.Head = newNode
		list.Tail = newNode
	} else {
		//定位到节点它的前驱
		pre := node.pre
		//如果定位到的节点前驱为nil，那么定位到的节点为链表头部
		if pre.IsNil() {
			//将新节点连接在老头部之前
			newNode.next = node
			node.pre = newNode
			//新节点成为头部
			list.Head = newNode
		} else {
			//将新节点插入到定位的节点之前
			//定位到的节点的前去节点pre连接到新节点上
			pre.next = newNode
			newNode.pre = pre
			//定位到的节点的后驱节点node.next链接到新节点上
			node.next.pre = newNode
			newNode.next = node.next
		}
	}
	
	//列表长度加一
	list.Len++
}

//添加节点到链表尾部的第N个元素之后，N=0表示新节点成为新的尾部
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	//加锁
	list.lock.Lock()
	defer list.lock.Unlock()
	
	//当索引超过列表长度，报错
	if n > list.Len {
		panic("index out")
	}
	
	//先找出尾部
	node := list.Tail
	
	//往前遍历拿到第N+1个元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	
	newNode := &ListNode{value: v}
	//如果定位到节点为空，表示链表为空，将新节点设置为头部和尾部
	if node.IsNil() {
		list.Head = newNode
		list.Tail = newNode
	} else {
		//定位到节点的后驱
		next := node.next
		
		//如果定位到节点后驱为nil，那么定位到的节点为尾部，需要换尾部
		if next.IsNil() {
			node.next = newNode
			newNode.pre = node
			//新节点成为尾部
			list.Tail = newNode
		} else {
			//将新节点插入到定位的节点后
			newNode.pre = node
			node.next = newNode
			
			newNode.next = next
			next.pre = newNode
		}
		
	}
	list.Len++
}

//从头部开始获取某个位置列表节点
func (list *DoubleList) IndexFromHead(n int) *ListNode {
	//索引过长
	if n > list.Len {
		return nil
	}
	node := list.Head
	for i := 0; i < n; i++ {
		node = node.next
	}
	return node
	
}

//从尾部开始获取某个位置节点
func (list *DoubleList) IndexFromTail(n int) *ListNode {
	if n > list.Len {
		return nil
	}
	node := list.Head
	for i := 0; i < n; i++ {
		node = node.pre
	}
	return node
}

//从头部开始移除并返回某个位置的节点
func (list *DoubleList) PopFromHead(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()
	
	//索引长度超过列表长度
	if n > list.Len {
		return nil
	}
	//获取头部
	node := list.Head
	
	for i := 1; i <= n; i++ {
		node = node.next
	}
	//移除节点的前驱和后驱
	pre := node.pre
	next := node.next
	//如果前驱和后驱是nil,表示是列表中的唯一节点
	if pre.IsNil() && next.IsNil() {
		list.Head = nil
		list.Tail = nil
	} else if pre.IsNil() {
		//表示移除的是头结点，下一个节点成头结点
		list.Head = next
		next.pre = nil
	} else if next.IsNil() {
		//表示删除的时尾结点
		list.Tail = pre
		pre.next = nil
	} else {
		//删除的是中间节点
		pre.next = next
		next.pre = pre
	}
	list.Len--
	return node
}

// 从尾部开始往前找，获取第N+1个位置的节点，并移除返回
func (list *DoubleList) PopFromTail(n int) *ListNode {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()
	
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.Len {
		return nil
	}
	
	// 获取尾部
	node := list.Tail
	
	// 往前遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	
	// 移除的节点的前驱和后驱
	pre := node.pre
	next := node.next
	
	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.Head = nil
		list.Tail = nil
	} else if pre.IsNil() {
		// 表示移除的是头部节点，那么下一个节点成为头节点
		list.Head = next
		next.pre = nil
	} else if next.IsNil() {
		// 表示移除的是尾部节点，那么上一个节点成为尾节点
		list.Tail = pre
		pre.next = nil
	} else {
		// 移除的是中间节点
		pre.next = next
		next.pre = pre
	}
	
	// 节点减一
	list.Len--
	return node
}
