package dpd

import (
	"mail-callbacks/config"
	"mail-callbacks/domain"
	"mail-callbacks/infra"
	"mail-callbacks/usecases"
)

var (
	MessageRepository domain.MessageRepository
	MessageService    usecases.MessageRepository
	ImapService       usecases.ImapRepository
)

func init() {
	MessageRepository = infra.NewKafkaProducer(config.GetConfig().KafkaProduceTopic())
	MessageService = usecases.NewMessageService(MessageRepository)
	ImapService = usecases.NewImapService(config.GetConfig().ImapAddress(), config.GetConfig().AmazonEmail(), config.GetConfig().AmazonPassword(), config.GetConfig().RequestTime(), config.GetConfig().RequestAmount())
}
