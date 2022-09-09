// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store
package doordash

import (
	"net/url"
	"time"
)

// Object for creating a new business
type NewBusiness struct {
	ExternalBusinessID string `json:"external_business_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	ActivationStatus   string `json:"activation_status"`
}

// Object for sending a business update
type BusinessUpdate struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	ActivationStatus string `json:"activation_status"`
}

// Object containing response information for businesses
type BusinessInfo struct {
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	ExternalBusinessID string    `json:"external_business_id"`
	ActivationStatus   string    `json:"activation_status"`
	CreatedAt          time.Time `json:"created_at"`
	LastUpdatedAt      time.Time `json:"last_updated_at"`
	IsTest             bool      `json:"is_test"`
	ExternalMetadata   struct {
		NumberOfStores    int    `json:"number_of_stores"`
		ClientEmail       string `json:"client_email"`
		ClientPhoneNumber string `json:"client_phone_number"`
	} `json:"external_metadata"`
}

// Object containing response information for mulitple businesses
type BusinessInfoList struct {
	Result            []BusinessInfo `json:"result"`
	ContinuationToken string         `json:"continuation_token"`
	ResultCount       int            `json:"result_count"`
}

// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store/operation/CreateBusiness
func (c *Client) CreateBusiness(b *NewBusiness) (*BusinessInfo, error) {
	res := &BusinessInfo{}
	if err := c.makeRequest("POST", "/developer/v1/businesses", nil, b, res); err != nil {
		return nil, err
	}

	return res, nil
}

// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store/operation/ListBusiness
func (c *Client) ListBusinesses(activationStatus string, paginationToken string) (*BusinessInfoList, error) {
	params := url.Values{
		"activation_status": []string{activationStatus},
		"pagination_token":  []string{paginationToken},
	}

	res := &BusinessInfoList{}
	if err := c.makeRequest("GET", "/developer/v1/businesses", params, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store/operation/GetBusiness
func (c *Client) GetBusiness(externalBusinessID string) (*BusinessInfo, error) {
	res := &BusinessInfo{}
	if err := c.makeRequest("GET", ("/developer/v1/businesses/" + externalBusinessID), nil, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store/operation/UpdateBusiness
func (c *Client) UpdateBusiness(externalBusinessID string, b *BusinessUpdate) (*BusinessInfo, error) {
	res := &BusinessInfo{}
	if err := c.makeRequest("PATCH", ("/developer/v1/businesses/" + externalBusinessID), nil, b, res); err != nil {
		return nil, err
	}

	return res, nil
}
