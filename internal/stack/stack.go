// Package stack implements a simplified Stack manager
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

type stack struct {
	values []int
}

var (
	errStackUnderflow = errors.New("stack underflow")
)

// Push adds an element to the top of the stack
func (s *stack) Push(v int) {
	s.values = append(s.values, v)
}

// Pop pulls an element from the top of the stack and removes it
func (s *stack) Pop() (int, error) {
	valueCount := len(s.values)
	if valueCount == 0 {
		return 0, errStackUnderflow
	}

	last := s.values[valueCount-1]
	s.values = s.values[:valueCount-1]

	return last, nil
}

// Last returns the top element of the stack, but does not remove it
func (s *stack) Last() (int, error) {
	valueCount := len(s.values)
	if valueCount == 0 {
		return 0, errStackUnderflow
	}
	return s.values[valueCount-1], nil
}

// Drop removes the top element of the stack without returning it
func (s *stack) Drop() error {
	valueCount := len(s.values)
	if valueCount == 0 {
		return errStackUnderflow
	}
	s.values = s.values[:valueCount-1]

	return nil
}
