package ximmio

import "net/http"

type Client struct {
	client      *http.Client
	CompanyCode string
}

// NewClient creates a new client to work with the Ximmio Waste API.
func NewClient(companyCode string) *Client {
	client := &Client{
		client:      http.DefaultClient,
		CompanyCode: companyCode,
	}

	return client
}
