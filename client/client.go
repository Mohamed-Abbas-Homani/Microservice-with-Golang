package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mohamed-Abbas-Homani/microservice/proto"
	"github.com/Mohamed-Abbas-Homani/microservice/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//GRPC client
func NewGRPCClient(remoteAddr string) (proto.PriceFetcherClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := proto.NewPriceFetcherClient(conn)

	return c, nil
}

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

//Client fetching logic
func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)

	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(res.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service respond with non OK status code %s", httpErr["error"])
	}

	priceRes := new(types.PriceResponse)

	if err := json.NewDecoder(res.Body).Decode(priceRes); err != nil {
		return nil, err
	}

	return priceRes, nil
}
