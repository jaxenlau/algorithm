package stack

import "sync"

type Stack interface {
	IsEmpty() bool
	IsFull() bool
	Push(element interface{})
	Pop() interface{}
	Top() interface{}
}

func New(size int) *StackImpl {
	return &StackImpl{
		RWMutex: sync.RWMutex{},
		size:    size,
		buf:     make([]Element, 0, size),
	}
}

type Element interface{}

type StackImpl struct {
	sync.RWMutex
	size int
	buf  []Element
}

func (s *StackImpl) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return len(s.buf) == 0
}

func (s *StackImpl) IsFull() bool {
	s.RLock()
	defer s.RUnlock()
	return len(s.buf) == s.size
}

func (s *StackImpl) Push(element Element) {
	s.Lock()
	defer s.Unlock()
	s.buf = append(s.buf, element)
}

func (s *StackImpl) Pop() *Element {
	s.Lock()
	defer s.Unlock()
	topIdx := len(s.buf) - 1
	element := s.buf[topIdx]
	s.buf = s.buf[:topIdx]
	return &element
}
func (s *StackImpl) Top() *Element {
	s.RLock()
	defer s.RUnlock()
	topIdx := len(s.buf) - 1
	element := s.buf[topIdx]
	return &element
}
