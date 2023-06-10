package recentstack

import (
	"errors"

	deque "gopkg.in/dnaeon/go-deque.v1"
)

type (
	RecentStack[T any] interface {
		Push(T)
		// Pop() (T,error)
		// IsEmpty() bool
		Length() int
	}

	recentStack[T any] struct {
		dq       *deque.Deque[T]
		capacity int
	}
)

// ErrNoRows is returned by New when capacity is less than 2.
// For a capacity of 1 you might as well just use a variable.
var ErrTooSmall = errors.New("recentstack: minimum sensible capacity is 2")

func New[T any](capacity int) (RecentStack[T], error) {
	if capacity < 2 {
		return nil, ErrTooSmall
	}
	return &recentStack[T]{
		dq:       deque.New[T](),
		capacity: capacity,
	}, nil
}

func (r *recentStack[T]) Length() int {
	return r.dq.Length()
}

func (r *recentStack[T]) Push(x T)  {
	r.dq.PushFront(x)
}
