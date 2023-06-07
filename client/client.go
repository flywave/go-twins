package client

import "github.com/flywave/go-twins/protocol"

type Handler func(requestId string, message *protocol.Envelope)

type Client interface {
	Connect() error
	Disconnect()
	Reply(requestId string, message *protocol.Envelope) error
	Send(message *protocol.Envelope) error
	Subscribe(handlers ...Handler)
	Unsubscribe(handlers ...Handler)
}
