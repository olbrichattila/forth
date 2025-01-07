// Package stack implement a simplified Stack manager
package stack

import (
	"errors"
)

// New creates a new stack
func New() Stacker {
	return &stack{}
}

// Stacker is the abstract description of the functions available
type Stacker interface {
	Push(v int)
	Pop() (int, error)
	Last() (int, error)
	Drop() error
}

type stack struct{
	values []int
}

var (
	errStackOverflow = errors.New("stack overflow")
)

// Push adds an element to the top of the stack
func (s *stack) Push(v int) {
	s.values = append(s.values, v)
}

// Pop pulls an element from the top of the stack and removes it
func (s *stack) Pop() (int, error) {
	valueCount := len(s.values);
	if valueCount == 0 {
		return 0, errStackOverflow
	}

	last := s.values[valueCount-1]
    s.values = s.values[:valueCount-1]

	return last, nil
}

// Last returns with the top of the stack, but not removed
func (s *stack) Last() (int, error) {
	valueCount := len(s.values);
	if valueCount == 0 {
		return 0, errStackOverflow
	}
	return s.values[valueCount-1], nil
}

// Drop removes the element from the top of the stack without returning
func (s *stack) Drop() error {
	valueCount := len(s.values);
	if valueCount == 0 {
		return errStackOverflow
	}
	 s.values = s.values[:valueCount-1]

	 return nil
}
