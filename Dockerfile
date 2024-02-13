FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN  go build -o /priceFetcher

EXPOSE 3000

CMD [ "/priceFetcher" ]