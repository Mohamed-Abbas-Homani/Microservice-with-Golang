package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Mohamed-Abbas-Homani/microservice/client"
	"github.com/Mohamed-Abbas-Homani/microservice/proto"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3000", "the port the json server is running on")
		grpcAddr = flag.String("grpc", ":4000", "the port the grpc server is running on")
		svc = NewLoggingService(&priceFetcher{})
		ctx = context.Background()
	)


	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}
	
	go func ()  {
		time.Sleep(3*time.Second)
		res, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "ETH"})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", res)
	}()

	go makeGRPCServer(*grpcAddr, svc)
	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
