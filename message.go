package main

type Request struct {
	SenderId      int
	SenderNodeId  int
	Service       string
	Data          []byte
	TransactionId string
}

type Reply struct {
	SenderId       int
	SenderNodeId   int
	ReceiverId     int
	ReceiverNodeId int
	Service        string
	Data           []byte
	TransactionId  string
}

type Message struct {
	SenderId int
	Topic    string
	Data     []byte
}
