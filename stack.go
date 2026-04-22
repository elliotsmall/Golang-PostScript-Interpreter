// Operand stack and its methods

package main

import "fmt"

type Stack struct {
	data []PSObject
}

func (s *Stack) Push(obj PSObject) {
	s.data = append(s.data, obj)
}

func (s *Stack) Pop() (PSObject, error) {
	if len(s.data) == 0 {
		return PSObject{}, fmt.Errorf("stack underflow error")
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, nil
}

func (s *Stack) Peek() (PSObject, error) {
	if len(s.data) == 0 {
		return PSObject{}, fmt.Errorf("stack underflow error")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Clear() {
	s.data = []PSObject{}
}

func (s *Stack) PopN(n int) ([]PSObject, error) {
	if len(s.data) < n {
		return nil, fmt.Errorf("stack underflow error: need %d elements, only have %d", n, len(s.data))
	}
	items := make([]PSObject, n)
	copy(items, s.data[len(s.data)-n:])
	s.data = s.data[:len(s.data)-n]
	return items, nil
}

func (s *Stack) Index(i int) (PSObject, error) {
	idx := len(s.data) - 1 - i
	if idx < 0 || idx >= len(s.data) {
		return PSObject{}, fmt.Errorf("stack index out of range")
	}
	return s.data[idx], nil
}
