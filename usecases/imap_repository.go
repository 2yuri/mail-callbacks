package usecases

type ImapRepository interface {
	ReadMessages() ([]MessageDTO, error)
}
