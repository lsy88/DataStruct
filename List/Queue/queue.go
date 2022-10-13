package Queue

import (
	"errors"
	"sync"
)

type Queue struct {
	lock sync.RWMutex
	Data []interface{}
}

func (q *Queue) Len() int {
	return len(q.Data)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Data) == 0
}

func (q *Queue) Cap() int {
	return cap(q.Data)
}

func (q *Queue) NewQueue() *Queue {
	q.Data = make([]interface{}, 0)
	return q
}

func (q *Queue) Enqueue(value interface{}) {
	q.lock.Lock()
	q.Data = append(q.Data, value)
	q.lock.Unlock()
}

func (q *Queue) Dequeue() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	
	th := q.Data
	if len(th) == 0 {
		return nil, errors.New("queue is empty")
	}
	value := th[0]
	q.Data = th[1:]
	return value, nil
}

//查询队列第一个数，不弹出
func (q *Queue) Front() interface{} {
	q.lock.RLock()
	item := q.Data[0]
	q.lock.RUnlock()
	return item
}
