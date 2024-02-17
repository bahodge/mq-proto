package main

import (
	"errors"
)

type client struct {
	id   int
	node Node
}

type Client interface {
	Id() int
	Connect(n Node) error
	Publish(topic string, msg *Message) error
	Subscribe(topic string, handler SubscriptionHandler) (Subscriber, error)
}

var id = 0

func NewClient() Client {
	id++
	return &client{
		id:   id,
		node: nil,
	}
}

func (c client) Id() int {
	return c.id
}

func (c *client) Connect(n Node) error {
	c.node = n
	return n.AddClient(c)
}

func (c client) Publish(topic string, msg *Message) error {
	if c.node == nil {
		return errors.New("not connected to node")
	}

	err := c.node.Publish(topic, msg)
	if err != nil {
		return err
	}

	return nil

}

func (c *client) Subscribe(topic string, handler SubscriptionHandler) (Subscriber, error) {
	if c.node == nil {
		return nil, errors.New("client not connected")
	}

	sub, err := c.node.Subscribe(c, topic, handler)
	if err != nil {
		return nil, err
	}

	return sub, nil

}

// func (c *client) Request(service string, data []byte) (*Reply, error) {
// 	txid++
// 	req := &Request{
// 		SenderId:      c.id,
// 		SenderNodeId:  c.node.Id(),
// 		Service:       service,
// 		Data:          data,
// 		TransactionId: fmt.Sprintf("txid %d", txid),
// 	}
// 	// Make the request
// 	err := c.node.Request(req)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return nil, errors.New("not implemented")
//
// }
//
// func (c *client) HandleRequest(req *Request) error {
// 	fmt.Printf("%s - %s received message\n", c.id, c.service)
// 	msg := fmt.Sprintf("reply: %s", string(req.Data))
// 	rep := &Reply{
// 		SenderId:       c.id,
// 		SenderNodeId:   c.node.Id(),
// 		ReceiverId:     req.SenderId,
// 		ReceiverNodeId: req.SenderNodeId,
// 		Service:        req.Service,
// 		Data:           []byte(msg),
// 		TransactionId:  req.TransactionId,
// 	}
//
// 	err := c.Reply(rep)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// func (c *client) Reply(rep *Reply) error {
// 	return c.node.Reply(rep)
//
// }
