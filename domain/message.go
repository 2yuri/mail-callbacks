package domain

import (
	"encoding/json"
	"time"
)

type Message struct {
	messageId     string
	mails         string
	action        string
	status        string
	statusMessage string
	date          time.Time
}

func NewMessage(messageId, destinations, action, status, statusMessage string, datetime time.Time) *Message {
	return &Message{
		messageId:     messageId,
		mails:         destinations,
		action:        action,
		status:        status,
		statusMessage: statusMessage,
		date:          datetime,
	}
}

func (s *Message) MessageId() string {
	return s.messageId
}

func (s *Message) ToBytes() ([]byte, error) {
	message := struct {
		MessageId     string    `json:"message_id"`
		Mails         string    `json:"destinations"`
		Action        string    `json:"action"`
		Status        string    `json:"status"`
		StatusMessage string    `json:"status_message"`
		Date          time.Time `json:"datetime"`
	}{
		MessageId:     s.messageId,
		Mails:         s.mails,
		Action:        s.action,
		Status:        s.status,
		StatusMessage: s.statusMessage,
		Date:          s.date,
	}

	return json.Marshal(message)
}
