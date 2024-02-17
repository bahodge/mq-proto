package main

import (
	"errors"
)

type Node interface {
	AddClient(c Client) error
	Topics() map[string]Topic
	Subscribe(c Client, topic string, handler SubscriptionHandler) (Subscriber, error)
	Publish(topic string, msg *Message) error
}

type node struct {
	id      int
	clients map[Client]bool
	topics  map[string]Topic
}

var nodeId = 0

func NewNode() *node {
	nodeId++

	return &node{
		id:      nodeId,
		clients: make(map[Client]bool),
		topics:  make(map[string]Topic),
	}
}

func (n *node) AddClient(c Client) error {
	_, ok := n.clients[c]
	if ok {
		return errors.New("client already added")
	}

	n.clients[c] = true

	return nil
}

func (n *node) Publish(topic string, msg *Message) error {
	var t Topic
	var ok bool
	var err error

	t, ok = n.topics[topic]
	if !ok {
		t, err = n.CreateTopic(topic)
		if err != nil {
			return err
		}
	}

	t.Publish(msg)
	return nil

}

func (n *node) Subscribe(c Client, topic string, handler SubscriptionHandler) (Subscriber, error) {
	var t Topic
	var ok bool
	var err error

	t, ok = n.topics[topic]
	if !ok {
		t, err = n.CreateTopic(topic)
		if err != nil {
			return nil, err
		}
	}

	sub := NewSubscriber(c, t, handler)
	err = t.Subscribe(sub)
	if err != nil {
		return nil, err
	}

	return sub, nil

}

func (n node) Topics() map[string]Topic {
	return n.topics
}

func (n *node) CreateTopic(topic string) (Topic, error) {
	_, ok := n.topics[topic]
	if !ok {
		t := NewTopic(topic)
		n.topics[topic] = t
		return t, nil
	}

	return nil, errors.New("topic already exists")
}

func (n *node) DeleteTopic(topic string) error {
	_, ok := n.topics[topic]
	if !ok {
		return errors.New("topic does not exist")
	}

	delete(n.topics, topic)

	return nil
}
