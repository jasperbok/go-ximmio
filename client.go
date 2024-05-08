package ximmio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://wasteapi.ximmio.com/api"
)

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

func (c *Client) GetAddress(postCode string, houseNumber int) (Address, error) {
	url := fmt.Sprintf("%s/GetAddress", baseURL)

	requestData := struct {
		CompanyCode string `json:"companyCode"`
		PostCode    string `json:"postCode"`
		HouseNumber int    `json:"houseNumber"`
	}{c.CompanyCode, postCode, houseNumber}

	data, err := json.Marshal(requestData)
	if err != nil {
		return Address{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return Address{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return Address{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Address{}, err
	}

	responseStruct := struct {
		DataList []Address `json:"dataList"`
	}{}

	err = json.Unmarshal(body, &responseStruct)
	if err != nil {
		return Address{}, err
	}

	return responseStruct.DataList[0], nil
}
