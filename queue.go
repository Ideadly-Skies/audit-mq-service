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
	subscribers []chan Message
	mu          sync.RWMutex
}

func NewQueue() *InMemoryQueue {
	return &InMemoryQueue{
		subscribers: make([]chan Message, 0),
	}
}

func (q *InMemoryQueue) Publish(msg Message) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	// Send to all subscribers
	for _, sub := range q.subscribers {
		select {
		case sub <- msg:
			// Message sent successfully
		default:
			// Skip if subscriber's channel is full
		}
	}
}

func (q *InMemoryQueue) Subscribe() <-chan Message {
	ch := make(chan Message, 10)

	q.mu.Lock()
	q.subscribers = append(q.subscribers, ch)
	q.mu.Unlock()

	return ch
}

// Global queue instance
var messageQueue = NewQueue()
