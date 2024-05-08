package ximmio

import "net/http"

type Client struct {
	client      *http.Client
}

// NewClient creates a new client to work with the Ximmio Waste API.
func NewClient() *Client {
	client := &Client{
		client: http.DefaultClient,
	}

	return client
}
