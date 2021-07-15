package main

import (
	"fmt"
	"log"
	"mail-callbacks/app/schedule"
	"mail-callbacks/app/smtp"
	"mail-callbacks/config"
)

func main() {
	log.Println("root path ", config.GetConfig().RootPath())

	logFile := "mail-callbacks.log"

	if len(config.GetConfig().RootPath()) > 0 {
		logFile += fmt.Sprintf(`%s\%s`, config.GetConfig().RootPath(), logFile)
	}

	log.Printf("Defining logfile %s", logFile)

	config.DefineLogFile(logFile)
	smtp.InitSMTPCodes()

	log.Printf("\n\n")
	log.Println("================== starting the service mail-callbacks ==========================")

	svc := schedule.NewMailWatcherService()
	svc.WatchMails()
}
