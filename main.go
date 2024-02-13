package main

import (
	// "context"
	"flag"
	// "fmt"
	// "log"

	// "github.com/Mohamed-Abbas-Homani/microservice/client"
)

func main() {
	// clt := client.New("http://localhost:3000")
	// price, err := clt.FetchPrice(context.Background(), "ETfH")
	// if err != nil {
	// 	log.Fatal(err) 
	// }
	// fmt.Printf("%+v\n", price)
	// return
	port := flag.String("port", ":3000", "the port the server is running on")
	flag.Parse()
	svc := NewLoggingService(&priceFetcher{})

	server := NewJSONAPIServer(*port, svc)
	server.Run()
}
