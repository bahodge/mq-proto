package main

const initialQueueCapacity = 10_000_000

type Queue[T any] interface {
	Push(T)
	Peek() T
	Next() T
	IsEmpty() bool
}

type queue[T any] struct {
	contents     []T
	currentIndex int
	queueCap     int
}

func (q queue[T]) IsEmpty() bool {
	return q.currentIndex == len(q.contents)
}

func (q *queue[T]) Push(element T) {
	q.contents = append(q.contents, element)
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{
		// TODO: Convert the slice into a fixed array
		contents:     make([]T, 0, initialQueueCapacity),
		queueCap:     initialQueueCapacity,
		currentIndex: 0,
	}
}

func NewQueueCap[T any](queueCap int) Queue[T] {
	return &queue[T]{
		// TODO: Convert the slice into a fixed array
		contents:     make([]T, 0, queueCap),
		queueCap:     queueCap,
		currentIndex: 0,
	}
}

func (q queue[T]) Peek() T {
	return q.contents[q.currentIndex]
}

var idxResetCount int = 0

func (q *queue[T]) Next() T {
	item := q.contents[q.currentIndex]
	if q.currentIndex+1 == q.queueCap {
		q.currentIndex = 0
		idxResetCount++
	} else {
		q.currentIndex++
	}
	return item
}
