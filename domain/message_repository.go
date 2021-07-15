package domain

type MessageRepository interface {
	ProduceMessage(*Message) error
}
