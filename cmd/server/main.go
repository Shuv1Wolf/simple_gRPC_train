package main

import (
	"log"
	"net"
	adder "simple_gRPC_train/adder/gRPC"

	trainv1 "github.com/Shuv1Wolf/train_protos/gen/go"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	trainv1.RegisterAdderServer(s, srv)

	log.Print("Server started")

	l, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
