package ximmio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

func (c *Client) request(method, path string, requestData []byte) ([]byte, error) {
	url := fmt.Sprintf("%s%s", baseURL, path)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestData))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, err
}

func (c *Client) GetAddress(postCode string, houseNumber int) (Address, error) {
	path := "/GetAddress"

	requestData := struct {
		CompanyCode string `json:"companyCode"`
		PostCode    string `json:"postCode"`
		HouseNumber int    `json:"houseNumber"`
	}{c.CompanyCode, postCode, houseNumber}

	data, err := json.Marshal(requestData)
	if err != nil {
		return Address{}, err
	}

	body, err := c.request(http.MethodPost, path, data)
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

func (c *Client) GetCalendars(start, end time.Time, addressId string) (Calendars, error) {
	calendars := Calendars{}

	path := "/GetCalendar"

	requestData := struct {
		CompanyCode string `json:"companyCode"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
		AddressID   string `json:"uniqueAddressId"`
	}{c.CompanyCode, start.Format("2006-01-02"), end.Format("2006-01-02"), addressId}

	data, err := json.Marshal(requestData)
	if err != nil {
		return calendars, err
	}

	body, err := c.request(http.MethodPost, path, data)
	if err != nil {
		return calendars, err
	}

	err = json.Unmarshal(body, &calendars)
	return calendars, err
}

func (c *Client) GetWasteTypes() ([]WasteType, error) {
	wasteTypes := []WasteType{}

	path := "/GetConfigOption"

	requestData := struct {
		CompanyCode string `json:"companyCode"`
		ConfigName  string `json:"configName"`
	}{c.CompanyCode, "ALL"}

	data, err := json.Marshal(requestData)
	if err != nil {
		return wasteTypes, err
	}

	body, err := c.request(http.MethodPost, path, data)
	if err != nil {
		return wasteTypes, err
	}

	responseStruct := struct {
		DataList []configOption `json:"dataList"`
	}{}

	err = json.Unmarshal(body, &responseStruct)
	if err != nil {
		return wasteTypes, err
	}

	for _, config := range responseStruct.DataList {
		wasteTypes = append(wasteTypes, config.toWasteType())
	}

	return wasteTypes, nil
}
