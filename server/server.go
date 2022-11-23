package main

import (
	"fmt"
	"gogrpc/schema/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("error while create listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("calculator service listened")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("error while serve %v", err)
	}
}
