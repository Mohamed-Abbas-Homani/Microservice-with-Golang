build:
	go build -o bin/priceFetcher

run: build
	bin/priceFetcher
