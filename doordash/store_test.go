package doordash

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var storeResponse = []byte(`{
	"name": "Neighborhood Deli #10",
	"external_business_id": "a0720d55-7cbe-41ce-8185-58285b7985cd",
	"external_store_id": "ed178ef3-b486-4ce8-8baa-5bc9f0f3fa4a",
	"phone_number": "+12065551212",
	"address": "901 Market Street, 6th Floor, San Francisco, CA, 94103",
	"status": "active",
	"is_test": false,
	"created_at": "2022-04-25T17:21:43Z",
	"last_updated_at": "2022-04-25T17:21:43Z"
  }`)

var storeListResponse = []byte(`{
	"result": [
	  {
		"name": "Neighborhood Deli #10",
		"external_business_id": "a0720d55-7cbe-41ce-8185-58285b7985cd",
		"external_store_id": "ed178ef3-b486-4ce8-8baa-5bc9f0f3fa4a",
		"phone_number": "+12065551212",
		"address": "901 Market Street, 6th Floor, San Francisco, CA, 94103",
		"status": "active",
		"is_test": false,
		"created_at": "2022-04-25T17:21:43Z",
		"last_updated_at": "2022-04-25T17:21:43Z"
	  }
	],
	"continuation_token": "string",
	"result_count": 0
  }`)

func TestCreateStore(t *testing.T) {
	testID := "B-12345"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != ("/developer/v1/businesses/" + testID + "/stores") {
			t.Errorf("expected request URL to be /developer/v1/businesses/%s/stores, got %s", testID, reqURL)
		}
		// Send response to be tested
		rw.Write(storeResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	payload := &NewStore{
		ExternalStoreID: "S-12345",
		Name:            "Neighborhood Deli",
		PhoneNumber:     "+12065551212",
		Address:         "901 Market Street, 6th Floor, San Francisco, CA, 94103",
	}

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.CreateStore(testID, payload)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &StoreInfo{}
	json.Unmarshal(storeResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}

func TestListStores(t *testing.T) {
	testID := "B-12345"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != ("/developer/v1/businesses/" + testID + "/stores") {
			t.Errorf("expected request URL to be /developer/v1/businesses/%s/stores, got %s", testID, reqURL)
		}
		// Send response to be tested
		rw.Write(storeListResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.ListStores(testID, "active", "token")
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &StoreInfoList{}
	json.Unmarshal(storeListResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}

func TestGetStore(t *testing.T) {
	testBID := "B-12345"
	testID := "S-12345"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != fmt.Sprintf("/developer/v1/businesses/%s/stores/%s", testBID, testID) {
			t.Errorf("expected request URL to be /developer/v1/businesses/%s/stores/%s, got %s", testBID, testID, reqURL)
		}
		// Send response to be tested
		rw.Write(storeResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.GetStore(testBID, testID)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &StoreInfo{}
	json.Unmarshal(storeResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}

func TestUpdateStore(t *testing.T) {
	testBID := "B-12345"
	testID := "S-12345"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != fmt.Sprintf("/developer/v1/businesses/%s/stores/%s", testBID, testID) {
			t.Errorf("expected request URL to be /developer/v1/businesses/%s/stores/%s, got %s", testBID, testID, reqURL)
		}
		// Send response to be tested
		rw.Write(storeResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	payload := &StoreUpdate{
		Name:        "Neighborhood Deli",
		PhoneNumber: "+12065551212",
		Address:     "901 Market Street, 6th Floor, San Francisco, CA, 94103",
	}

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.UpdateStore(testBID, testID, payload)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &StoreInfo{}
	json.Unmarshal(storeResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}
