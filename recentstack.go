package recentstack

import (
	deque "gopkg.in/dnaeon/go-deque.v1"
)

type (
	RecentStack[T any] interface {
		// Push(T)
		// Pop() (T,error)
		// IsEmpty() bool
		// Size() int
	}

	recentStack[T any] struct {
		dq       *deque.Deque[T]
		capacity int
	}
)

func New[T any](capacity int) RecentStack[T] {
	return recentStack[T]{
		dq:       deque.New[T](),
		capacity: capacity,
	}
}

// func (r *recentStack[T]) Length() int {
// return r.Length()
// }
