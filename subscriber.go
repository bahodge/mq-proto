package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

type SubscriptionHandler func(msg *Message) error
type Subscriber interface {
	ProcessMessage(msg *Message)
	Subscribe(ctx context.Context)
	Close() error
}

type subscriber struct {
	client     Client
	topic      Topic
	handler    SubscriptionHandler
	subscribed bool
	errch      chan error
	done       chan struct{}
}

func NewSubscriber(c Client, topic Topic, handler SubscriptionHandler) Subscriber {
	return &subscriber{
		client:     c,
		topic:      topic,
		handler:    handler,
		subscribed: false,
		errch:      make(chan error),
		done:       make(chan struct{}),
	}
}

func (s *subscriber) ProcessMessage(msg *Message) {
	err := s.handler(msg)
	if err != nil {
		s.errch <- err
	}

	// TODO: add some better error handling for unprocessable messages
}

func (s *subscriber) Subscribe(ctx context.Context) {
	if s.subscribed {
		s.errch <- errors.New("already subscribed to topic")

		return
	}

	go func() {
		// TODO: actually reach out to the remote topic
		for {
			select {
			case <-ctx.Done():
				s.errch <- ctx.Err()
			case <-s.done:
				fmt.Println("done subscribing")
				return
			case err := <-s.errch:
				log.Fatalf("received error! %s", err.Error())
			}
		}

	}()

}

func (s *subscriber) Close() error {
	if s.subscribed {
		s.done <- struct{}{}
	}
	return nil
}
