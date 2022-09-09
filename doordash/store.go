// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store
package doordash

import (
	"net/url"
	"time"
)

// Object for creating a new stores
type NewStore struct {
	ExternalStoreID string `json:"external_store_id"`
	Name            string `json:"name"`
	PhoneNumber     string `json:"phone_number"`
	Address         string `json:"address"`
}

// Object for sending a store update
type StoreUpdate struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

// Object containing response information for stores
type StoreInfo struct {
	Name               string    `json:"name"`
	ExternalBusinessID string    `json:"external_business_id"`
	ExternalStoreID    string    `json:"external_store_id"`
	PhoneNumber        string    `json:"phone_number"`
	Address            string    `json:"address"`
	Status             string    `json:"status"`
	IsTest             bool      `json:"is_test"`
	CreatedAt          time.Time `json:"created_at"`
	LastUpdatedAt      time.Time `json:"last_updated_at"`
}

// Object containing response information for mulitple stores
type StoreInfoList struct {
	Result            []StoreInfo `json:"result"`
	ContinuationToken string      `json:"continuation_token"`
	ResultCount       int         `json:"result_count"`
}

// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store/operation/CreateStore
func (c *Client) CreateStore(externalBusinessID string, body *NewStore) (*StoreInfo, error) {
	res := &StoreInfo{}
	if err := c.makeRequest("POST", ("/developer/v1/businesses/" + externalBusinessID + "/stores"), nil, body, res); err != nil {
		return nil, err
	}
	return res, nil
}

// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store/operation/ListStore
func (c *Client) ListStores(externalBusinessID string, activationStatus string, paginationToken string) (*StoreInfoList, error) {
	params := url.Values{
		"activation_status": []string{activationStatus},
		"pagination_token":  []string{paginationToken},
	}

	res := &StoreInfoList{}
	if err := c.makeRequest("GET", ("/developer/v1/businesses/" + externalBusinessID + "/stores"), params, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store/operation/GetStore
func (c *Client) GetStore(externalBusinessID string, externalStoreID string) (*StoreInfo, error) {
	res := &StoreInfo{}
	if err := c.makeRequest("GET", ("/developer/v1/businesses/" + externalBusinessID + "/stores/" + externalStoreID), nil, nil, res); err != nil {
		return nil, err
	}
	return res, nil
}

// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Business-and-Store/operation/UpdateStore
func (c *Client) UpdateStore(externalBusinessID string, externalStoreID string, body *StoreUpdate) (*StoreInfo, error) {
	res := &StoreInfo{}
	if err := c.makeRequest("PATCH", ("/developer/v1/businesses/" + externalBusinessID + "/stores/" + externalStoreID), nil, body, res); err != nil {
		return nil, err
	}
	return res, nil
}
