package main

import (
	"sync"
	"time"
)

// Simple in-memory queue for demonstration
type Message struct {
	ID        int
	Content   string
	Timestamp time.Time
}

type InMemoryQueue struct {
	messages []Message
	mu       sync.Mutex
	cond     *sync.Cond
}

func NewQueue() *InMemoryQueue {
	q := &InMemoryQueue{
		messages: make([]Message, 0),
	}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *InMemoryQueue) Publish(msg Message) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.messages = append(q.messages, msg)
	q.cond.Signal() // Wake up waiting subscriber
}

func (q *InMemoryQueue) Subscribe() <-chan Message {
	ch := make(chan Message, 10)
	go func() {
		for {
			q.mu.Lock()
			for len(q.messages) == 0 {
				q.cond.Wait() // Wait for new messages
			}
			msg := q.messages[0]
			q.messages = q.messages[1:]
			q.mu.Unlock()
			ch <- msg
		}
	}()
	return ch
}

// Global queue instance
var messageQueue = NewQueue()
