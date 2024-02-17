package main

import (
	"testing"
)

func TestQueueIsEmpty(t *testing.T) {
	// Make an empty queue
	q := NewQueue[string]()

	// Should be empty
	if !q.IsEmpty() {
		t.Fatal("The queue should be empty")
	}

	// Add something
	q.Push("Hello")
	q.Push("world")
	q.Push(":)")

	// Should have 3 elements
	if q.IsEmpty() {
		t.Fatal("The queue should not be empty")
	}

	q.Next()
	q.Next()
	q.Next()

	// Should be empty because we used all its elements
	if !q.IsEmpty() {
		t.Fatal("The queue should be empty")
	}

}

func TestQueueFIFO(t *testing.T) {
	// Test the queue is FIFO
	q := NewQueueCap[int](10000)

	// Add 10k elements to the queue
	for i := 0; i < 10000; i++ {
		q.Push(i)
	}

	// Retrieve 10k elements from the queue.
	// they SHOULD be in order

	for i := 0; i < 10000; i++ {
		// Get the next element
		x := q.Next()

		// The next element should be i, because it is fifo
		if x != i {
			t.Fatalf("The popped element should be %d, got %d", i, x)
		}
	}

}
