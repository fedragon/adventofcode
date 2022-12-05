package shared

import (
	"fmt"
)

type Stack struct {
	elems []string
}

func (s *Stack) Push(elem string) {
	s.elems = append(s.elems, elem)
}

func (s *Stack) PushN(elems ...string) {
	s.elems = append(s.elems, elems...)
}

func (s *Stack) Pop() string {
	if len(s.elems) == 0 {
		return ""
	}

	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]

	return elem
}

func (s *Stack) PopN(count int) []string {
	if len(s.elems) == 0 || len(s.elems) < count {
		return nil
	}

	i := len(s.elems) - count
	elems := s.elems[i:]
	s.elems = s.elems[:i]

	return elems
}

func (s *Stack) Peek() string {
	if len(s.elems) == 0 {
		return ""
	}

	return s.elems[len(s.elems)-1]
}

func (s *Stack) Reverse() {
	length := len(s.elems)
	for i := 0; i < length/2; i++ {
		j := length - i - 1

		s.elems[i], s.elems[j] = s.elems[j], s.elems[i]
	}
}

func (s *Stack) String() string {
	return fmt.Sprint(s.elems)
}
