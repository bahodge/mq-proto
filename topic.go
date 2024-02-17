package main

import (
	"errors"
)

type Topic interface {
	Publish(*Message)
	Subscribe(sub Subscriber) error
}

type topic struct {
	name        string
	queue       Queue[*Message]
	subscribers map[Subscriber]bool
}

func NewTopic(t string) Topic {
	return &topic{
		name:        t,
		queue:       NewQueue[*Message](),
		subscribers: make(map[Subscriber]bool),
	}
}

func (t *topic) Publish(msg *Message) {
	t.queue.Push(msg)
	t.tick()
}

func (t *topic) Subscribe(sub Subscriber) error {
	_, ok := t.subscribers[sub]
	if ok {
		return errors.New("already subscribed to topic")
	}

	t.subscribers[sub] = true

	return nil
}

// There should be a pool of workers to handle sending messages to subscribers

func (t *topic) tick() {
	if t.queue.IsEmpty() {
		return
	}
	msg := t.queue.Next()
	for sub := range t.subscribers {
		sub.ProcessMessage(msg)
	}
}
