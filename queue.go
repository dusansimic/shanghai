package shanghai

import (
	"fmt"
)

type Queue interface {
	Enqueue(elem *node)
	Dequeue() (*node, error)
	Empty() bool
	Peek() (*node, error)
}

type queue struct {
	elems []*node
}

func NewQueue() Queue {
	return &queue{
		elems: []*node{},
	}
}

func (q *queue) Enqueue(elem *node) {
	q.elems = append(q.elems, elem)
}

func (q *queue) Dequeue() (*node, error) {
	if q.Empty() {
		return nil, fmt.Errorf("empty queue")
	}

	element := q.elems[0]
	q.elems = q.elems[1:]
	return element, nil
}

func (q *queue) Empty() bool {
	return len(q.elems) == 0
}

func (q *queue) Peek() (*node, error) {
	if q.Empty() {
		return nil, fmt.Errorf("empty queue")
	}
	return q.elems[0], nil
}
