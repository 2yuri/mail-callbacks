package usecases

import "time"

type MessageDTO struct {
	MessageId string
	Mail      string
	Action    string
	Status    string
	Date      time.Time
}
