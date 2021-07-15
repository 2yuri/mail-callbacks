package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	rootPath          string
	kafkaBrokers      string
	amazonEmail       string
	amazonPassword    string
	requestTime       int64
	responseAmount    int
	kafkaProduceTopic string
	imapAddress       string
}

var config *Config

func init() {
	serviceFile, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatalf("Não foi possivel resgatar o caminho do serviço %v", err)
	}
	rootPath := extractPath(serviceFile)
	if err != nil {
		log.Fatalf("Não foi possivel definir o caminho do executável: %v", err)
	}
	log.Print("RootPath: ", rootPath)

	if err = createTemp(rootPath); err != nil {
		log.Fatalf("Error when trying to create the temp folder [%v]", err)
	}

	requestTime, err := strconv.ParseInt(os.Getenv("REQUEST_TIME"), 10, 64)
	if err != nil {
		log.Fatalf("Não foi converter o tempo de request: %v", err)
	}

	responseAmount, err := strconv.Atoi(os.Getenv("RESPONSE_AMOUNT"))
	if err != nil {
		log.Fatalf("Não foi converter o valor de emails: %v", err)
	}

	config = &Config{
		rootPath:          rootPath,
		kafkaBrokers:      os.Getenv("KAFKA_BROKERS"),
		amazonEmail:       os.Getenv("AWS_MAIL"),
		amazonPassword:    os.Getenv("AWS_MAIL_PASSWORD"),
		requestTime:       requestTime,
		responseAmount:    responseAmount,
		kafkaProduceTopic: os.Getenv("KAFKA_PRODUCE_TOPIC"),
		imapAddress:       os.Getenv("IMAP_ADDRESS"),
	}
}

func (c *Config) RootPath() string {
	return c.rootPath
}

func (c *Config) KafkaBrokers() string {
	return c.kafkaBrokers
}

func (c *Config) ImapAddress() string {
	return c.imapAddress
}

func (c *Config) AmazonEmail() string {
	return c.amazonEmail
}

func (c *Config) AmazonPassword() string {
	return c.amazonPassword
}

func (c *Config) RequestTime() int64 {
	return c.requestTime
}

func (c *Config) RequestAmount() int {
	return c.responseAmount
}

func (c *Config) KafkaProduceTopic() string {
	return c.kafkaProduceTopic
}

func GetConfig() *Config {
	return config
}

func extractPath(servicePath string) string {
	path := strings.Split(servicePath, `\`)
	return strings.Join(path[0:len(path)-1], `\`)
}

func createTemp(path string) error {
	_, err := os.Stat(fmt.Sprintf("%s/temp", path))
	if err != nil {
		err := os.Mkdir(fmt.Sprintf("%s/temp", path), 0666)
		if err != nil {
			return fmt.Errorf("cannot create temp folder: %s", err)
		}
	}

	return nil
}
