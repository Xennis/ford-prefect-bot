package main

import (
	"log"
	"net"
	"net/http"

	pb "github.com/Xennis/ford-prefect-bot/telegram-bot-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	errRPC := make(chan error)
	errHTTP := make(chan error)
	go listenRPC(errRPC)
	go listenHTTP(errHTTP)

	select {
	case err := <-errRPC:
		log.Fatalf("Exited rpc with error: %s", err.Error())
	case err := <-errHTTP:
		log.Fatalf("Exited http with error: %s", err.Error())
	}
}

func listenRPC(errChan chan error) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		errChan <- err
		return
	}
	s := grpc.NewServer()
	pb.RegisterTelegramBotApiServer(s, &server{})
	reflection.Register(s) // Register reflection service on gRPC server.
	errChan <- s.Serve(lis)
}

func listenHTTP(errchan chan error) {
	mux := http.NewServeMux()
	mux.Handle("/ready", http.HandlerFunc(readinessProbeHandler))
	mux.Handle("/healthz", http.HandlerFunc(livenessProbeHandler))
	errchan <- http.ListenAndServe(":8080", mux)
}
