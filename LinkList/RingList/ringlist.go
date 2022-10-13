package main

import (
	"fmt"
	"reflect"
)

//循环链表

type Ring struct {
	Data interface{}
	Next *Ring //后驱节点
	Prev *Ring //前驱节点
}

//初始化空的循环链表，前去和后驱都指向自己
func (r *Ring) Init() *Ring {
	r.Next = r
	r.Prev = r
	return r
}

// NewRing 创建指定大小为N的循环链表
func NewRing(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.Next = &Ring{Prev: p}
		p = p.Next
	}
	p.Next = r
	r.Prev = p
	return r
}

//获取上一个节点
func (r *Ring) PrevNode() *Ring {
	if r.Next == nil {
		return r.Init()
	}
	return r.Prev
}

//获取下一个节点
func (r *Ring) NextNode() *Ring {
	if r.Next == nil {
		return r.Init()
	}
	return r.Next
}

//获取循环链表长度
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.NextNode(); p != r; p = p.Next {
			n++
		}
	}
	return n
}

//获取第n个节点
//因为链表是循环的，当n为负数，表示往前遍历，否则就是往后遍历
func (r *Ring) Search(n int) *Ring {
	if r.Next == nil {
		return r.Init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.Prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.Next
		}
	}
	return r
}

func (r *Ring) Insert(s *Ring) *Ring {
	n := r.NextNode() //找到r的后驱节点
	if s != nil {
		p := s.PrevNode() //此时p指向的还是s
		r.Next = s
		s.Prev = r
		n.Prev = p
		p.Next = n
	}
	return n
}

//删除节点后面的n个节点
func (r *Ring) Delete(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Insert(r.Search(n + 1))
}

func main() {
	// 第一个节点
	r := &Ring{Data: -1}
	
	// 链接新的五个节点
	r.Insert(&Ring{Data: 1})
	r.Insert(&Ring{Data: 2})
	r.Insert(&Ring{Data: 3})
	r.Insert(&Ring{Data: 4})
	
	node := r
	for {
		// 打印节点值
		fmt.Println(node.Data)
		
		// 移到下一个节点
		node = node.NextNode()
		
		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}

//将切片数组转换成map
func SliceToMap(i interface{}) (interface{}, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to converts %#v of type %T to map[interface{}]struct{}", i, i)
	}
	
	t := reflect.TypeOf(i)
	kind := t.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice or array", i, i)
	}
	v := reflect.ValueOf(i)
	mt := reflect.MapOf(t.Elem(), reflect.TypeOf(struct{}{}))
	mv := reflect.MakeMapWithSize(mt, v.Len())
	for j := 0; j < v.Len(); j++ {
		mv.SetMapIndex(v.Index(j), reflect.ValueOf(struct{}{}))
	}
	return mv.Interface(), nil
}
