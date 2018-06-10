package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	log.Printf("Start service")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go listenSignals(sigs)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	handler := createHandler()
	log.Fatal(http.ListenAndServeTLS(":8443", "/etc/nginx/ssl/nginx.crt", "/etc/nginx/ssl/nginx.key", handler))
}

func listenSignals(sigs chan os.Signal) {
	sig := <-sigs
	log.Printf("Exit service: %v\n", sig)
	os.Exit(1)
}
