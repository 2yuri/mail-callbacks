package usecases

type MessageRepository interface {
	ProduceMessage(MessageDTO) error
}
