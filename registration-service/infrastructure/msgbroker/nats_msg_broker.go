package msgbroker

import "github.com/nats-io/nats.go"

type NATSMsgBroker struct {
	nc *nats.Conn
}

func New(nc *nats.Conn) *NATSMsgBroker {

	return &NATSMsgBroker{
		nc: nc,
	}
}

func (nmb *NATSMsgBroker) Publish(subject string, event []byte) {

	nmb.nc.Publish(subject, event)
}
