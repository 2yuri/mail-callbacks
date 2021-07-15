package usecases

import (
	"mail-callbacks/app/smtp"
	"mail-callbacks/domain"
	"errors"
)

type MessageService struct {
	repo domain.MessageRepository
}

func NewMessageService(repo domain.MessageRepository) *MessageService {
	return &MessageService{
		repo: repo,
	}
}

func (e *MessageService) ProduceMessage(dto MessageDTO) error {
	message, err := getStatusMessageByCode(dto.Status)
	if err != nil {
		return err
	}

	return e.repo.ProduceMessage(domain.NewMessage(dto.MessageId, dto.Mail, dto.Action, dto.Status, message, dto.Date))
}

func getStatusMessageByCode(status string) (string, error) {
	statusMessage := smtp.GetSTMPStatus().Codes()
	if status == "" {
		return "", errors.New("cannnot get statusMessage from empty status")
	}
	code := formatCode(status)
	message := statusMessage["X."+code]

	if message.Description == "" {
		return "", errors.New("cannnot find a messge for this code: X." + code)
	}

	return message.Description, nil
}

func formatCode(s string) string {
	m := 0
	for i := range s {
		if m >= 2 {
			return s[i:]
		}
		m++
	}
	return s[:0]
}
