package graph

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list: list.New()}
}

func (s *Stack) Push(v string) {
	s.list.PushBack(v)
}

func (s *Stack) Pop() (string, bool) {
	if s.list.Len() == 0 {
		return "", false
	}
	e := s.list.Back()
	s.list.Remove(e)
	return e.Value.(string), true
}

func (s *Stack) Peek() (string, bool) {
	if s.list.Len() == 0 {
		return "", false
	}
	return s.list.Back().Value.(string), true
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}
