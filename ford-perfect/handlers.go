package main

import (
	"log"
	"net/http"

	pb "github.com/Xennis/ford-prefect-bot/telegram-bot-api"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

func readinessProbeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func livenessProbeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *server) ReceiveUpdates(ctx context.Context, in *pb.Update) (*empty.Empty, error) {
	log.Printf("Receive update: %v\n", in.UpdateId)
	m := in.Message
	if m == nil {
		return new(empty.Empty), nil
	}

	c := m.Chat
	if c != nil {
		log.Printf("Chat id: %v\n", c.Id)
	}
	return new(empty.Empty), nil
}
