package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/logging"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var serviceLogger *logging.Logger

func main() {
	ctx := context.Background()

	loggingClient, err := logging.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatalf("Could not create logger: %v", err)
	}
	defer loggingClient.Close()
	serviceLogger = loggingClient.Logger("ford-perfect")
	serviceLogger.Log(logging.Entry{Payload: "Start service"})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go listenSignals(sigs)

	errHTTPS := make(chan error)
	go listenHTTPS(errHTTPS)

	errHTTP := make(chan error)
	go listenHTTP(errHTTP)

	select {
	case err := <-errHTTPS:
		log.Fatalf("Exited https with error: %s", err.Error())
	case err := <-errHTTP:
		log.Fatalf("Exited http with error: %s", err.Error())
	}
}

func listenSignals(sigs chan os.Signal) {
	sig := <-sigs
	log.Fatalf("Exit service: %v\n", sig)
}

func listenHTTP(errorchan chan error) {
	errorchan <- http.ListenAndServe(":8080", createHTTPHandler())
}

func listenHTTPS(errorchan chan error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	handler := createHTTPSHalder(bot.Token)
	errorchan <- http.ListenAndServeTLS(":8443", "/etc/nginx/ssl/nginx.crt", "/etc/nginx/ssl/nginx.key", handler)
}
