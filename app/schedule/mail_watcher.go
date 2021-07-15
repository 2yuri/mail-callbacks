package schedule

import (
	"mail-callbacks/config"
	"mail-callbacks/dpd"
	"log"
	"time"
)

type mailWatcherService struct {
}

func NewMailWatcherService() *mailWatcherService {
	return &mailWatcherService{}
}

func (s *mailWatcherService) WatchMails() {
	for {
		log.Print("Runnin Watcher!")

		messsages, err := dpd.ImapService.ReadMessages()
		if err != nil {
			log.Printf("cannot get messages: %v", err)
			s.wait()
			continue
		}

		for _, val := range messsages {
			err := dpd.MessageService.ProduceMessage(val)
			if err != nil {
				log.Printf("cannot produce message: %v", err)
				continue
			}
		}

		s.wait()
	}
}

func (s *mailWatcherService) wait() {
	log.Print("awaiting 1min...")
	time.Sleep(time.Duration(config.GetConfig().RequestTime()) * time.Minute)
	log.Print("finished await time.")
}
