package Stack

import (
	"errors"
)

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Cap() int {
	return cap(s)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (interface{}, error) {
	th := *s
	if len(th) == 0 {
		return nil, errors.New("len is 0")
	}
	value := th[len(th)-1]
	*s = th[:len(th)-1]
	return value, nil
}

func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("len is 0")
	}
	return s[len(s)-1], nil
}
