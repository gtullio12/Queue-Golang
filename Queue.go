package main

type Queue struct {
	store []any
}

func (queue *Queue) Add(n any) {
	queue.store = append(queue.store, n)
}

func (queue *Queue) Remove() (elem any) {
	elem = queue.store[0]
	queue.store = queue.store[1:]
	return
}

func (queue *Queue) Peek() (elem any) {
	elem = queue.store[0]
	return
}
