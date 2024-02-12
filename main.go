package main

import "flag"

func main() {
	port := flag.String("port", ":3000", "the port the server is running on")
	flag.Parse()
	svc := NewLoggingService(&priceFetcher{})

	server := NewJSONAPIServer(*port, svc)
	server.Run()
}
