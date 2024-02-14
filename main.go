package main

import (
	"flag"
)

func main() {
	jsonAddr := flag.String("json", ":3000", "the port the json server is running on")
	grpcAddr := flag.String("grpc", ":4000", "the port the grpc server is running on")

	flag.Parse()
	svc := NewLoggingService(&priceFetcher{})

	go makeGRPCServer(*grpcAddr, svc)
	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
