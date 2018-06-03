package main

import (
	"fmt"
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
	fmt.Printf("Receive update: %v", in.UpdateId)
	return new(empty.Empty), nil
}
