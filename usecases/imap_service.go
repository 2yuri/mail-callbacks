package usecases

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type ImapService struct {
	imapAddress   string
	mail          string
	password      string
	requestAmount int
}

func NewImapService(imapAddress, mail, password string, requestTime int64, requestAmount int) *ImapService {
	return &ImapService{
		imapAddress:   imapAddress,
		mail:          mail,
		password:      password,
		requestAmount: requestAmount,
	}
}

func (s *ImapService) ReadMessages() ([]MessageDTO, error) {
	conn, err := s.connect()
	if err != nil {
		return nil, err
	}

	defer conn.Logout()
	loginErr := s.login(conn)
	if loginErr != nil {
		return nil, loginErr
	}

	mails, set, err := s.getMails(conn)
	if err != nil {
		return nil, err
	}

	messages, err := s.parseIntoMessage(mails)
	if err != nil {
		return nil, err
	}

	callbackErr := s.runningAfterRead(conn, set)
	if err != nil {
		return nil, callbackErr
	}

	return messages, nil
}

func (s *ImapService) connect() (*client.Client, error) {
	return client.DialTLS(s.imapAddress, nil)
}

func (s *ImapService) login(conn *client.Client) error {
	return conn.Login(s.mail, s.password)
}

func (s *ImapService) getMails(conn *client.Client) (chan *imap.Message, *imap.SeqSet, error) {
	_, err := conn.Select("[Gmail]/Todos os e-mails", false)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot select mailbox: %v", err)
	}

	seqset := new(imap.SeqSet)

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{"\\Seen"}
	uids, err := conn.Search(criteria)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot select message: %v", err)
	}

	for i := 0; i < s.requestAmount; i++ {
		seqset.AddNum(uids[len(uids)-i-1])
	}

	messages := make(chan *imap.Message, s.requestAmount)
	done := make(chan error, 1)
	go func() {
		done <- conn.Fetch(seqset, []imap.FetchItem{imap.FetchRFC822Text}, messages)
	}()

	return messages, seqset, nil
}

func (s *ImapService) parseIntoMessage(messages chan *imap.Message) ([]MessageDTO, error) {
	var response []MessageDTO

	for msg := range messages {
		for _, value := range msg.Body {
			len := value.Len()
			buf := make([]byte, len)
			n, err := value.Read(buf)
			if err != nil {
				return nil, err
			}
			if n != len {
				return nil, err
			}
			var imapDto ImapDTO

			jsonError := json.Unmarshal(buf, &imapDto)
			if err != nil {
				return nil, fmt.Errorf("cannot parse value into imapDto: %v", jsonError)
			}

			var imapContent ImapContentDTO
			jsonError = json.Unmarshal([]byte(imapDto.Message), &imapContent)
			if err != nil {
				return nil, fmt.Errorf("cannot parse value into imapContentDto: %v", jsonError)
			}

			for _, m := range imapContent.Bounce.BouncedRecipients {
				response = append(response, MessageDTO{
					MessageId: imapDto.MessageID,
					Mail:      m.EmailAddress,
					Action:    m.Action,
					Status:    m.Status,
					Date:      imapContent.Mail.Timestamp,
				})
			}
		}
	}
	return response, nil
}

func (s *ImapService) runningAfterRead(conn *client.Client, seqset *imap.SeqSet) error {
	flags := []interface{}{imap.SeenFlag}
	err := conn.UidStore(seqset, "+FLAGS.SILENT", flags, nil)
	if err != nil {
		log.Print("Erro ao dar uidstore: ", err)
		return err
	}

	return nil
}
