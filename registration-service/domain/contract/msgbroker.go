package contract

type MSGBroker interface {
	Publish(subject string, event []byte)
}
