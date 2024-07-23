package main

import (
	"testing"
)

func TestAddToQueue(t *testing.T) {
	var queue *Queue = &Queue{
		store: []any{1, 2, 3, 4},
	}

	queue.Add(5)

	if len(queue.store) != 5 {
		t.Errorf("Did not add element correctly")
	}

	if queue.store[4] != 5 {
		t.Errorf("Queue should've added: %v, but was: %v\n", 5, queue.store[4])
	}
}

func TestRemoveFromQueue(t *testing.T) {
	var queue *Queue = &Queue{
		store: []any{},
	}
	queue.Add("1")
	queue.Add("2")
	queue.Add("3")
	elem := queue.Remove()

	if len(queue.store) != 2 {
		t.Errorf("Did not remove element correctly")
	}

	if elem != "1" {
		t.Errorf("Queue should've removed: %v, but was: %v\n", "1", elem)
	}
}

func TestPeekFromQueue(t *testing.T) {
	var queue *Queue = &Queue{
		store: []any{},
	}
	queue.Add("1")
	queue.Add("2")
	queue.Add("3")
	elem := queue.Peek()

	if len(queue.store) != 3 {
		t.Errorf("Queue should not remove element")
	}

	if elem != "1" {
		t.Errorf("Queue should've removed: %v, but was: %v\n", "1", elem)
	}
}
