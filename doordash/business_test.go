package doordash

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var businessResponse = []byte(`{
	"name": "Neighborhood Deli",
	"description": "A neighborhood deli serving many tasty sandwiches and soups.",
	"external_business_id": "a0720d55-7cbe-41ce-8185-58285b7985cd",
	"activation_status": "active",
	"created_at": "2022-04-25T17:21:43Z",
	"last_updated_at": "2022-04-25T17:21:43Z",
	"is_test": false,
	"external_metadata": {
		"number_of_stores": 10,
		"client_email": "merchant@doordash.com",
		"client_phone_number": "+12065551212"
	}
}`)

var businessListResponse = []byte(`{
	"result": [
	  {
		"name": "Neighborhood Deli",
		"description": "A neighborhood deli serving many tasty sandwiches and soups.",
		"external_business_id": "a0720d55-7cbe-41ce-8185-58285b7985cd",
		"activation_status": "active",
		"created_at": "2022-04-25T17:21:43Z",
		"last_updated_at": "2022-04-25T17:21:43Z",
		"is_test": false,
		"external_metadata": {
		  "number_of_stores": 10,
		  "client_email": "merchant@doordash.com",
		  "client_phone_number": "+12065551212"
		}
	  }
	],
	"continuation_token": "string",
	"result_count": 0
  }`)

func TestCreateBusiness(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != "/developer/v1/businesses" {
			t.Errorf("expected request URL to be /developer/v1/businesses, got %s", reqURL)
		}
		// Send response to be tested
		rw.Write(businessResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	payload := &NewBusiness{
		ExternalBusinessID: "B-12345",
		Name:               "Neighborhood Deli",
		Description:        "A neighborhood deli serving many tasty sandwiches and soups.",
		ActivationStatus:   "active",
	}

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.CreateBusiness(payload)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &BusinessInfo{}
	json.Unmarshal(businessResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}

func TestListBusinesses(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != "/developer/v1/businesses" {
			t.Errorf("expected request URL to be /developer/v1/businesses, got %s", reqURL)
		}
		// Send response to be tested
		rw.Write(businessListResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.ListBusinesses("active", "token")
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &BusinessInfoList{}
	json.Unmarshal(businessListResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}

func TestGetBusiness(t *testing.T) {
	testID := "B-12345"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != ("/developer/v1/businesses/" + testID) {
			t.Errorf("expected request URL to be /developer/v1/businesses/%s, got %s", testID, reqURL)
		}
		// Send response to be tested
		rw.Write(businessResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.GetBusiness(testID)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &BusinessInfo{}
	json.Unmarshal(businessResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}

func TestUpdateBusiness(t *testing.T) {
	testID := "B-12345"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != ("/developer/v1/businesses/" + testID) {
			t.Errorf("expected request URL to be /developer/v1/businesses/%s, got %s", testID, reqURL)
		}
		// Send response to be tested
		rw.Write(businessResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	payload := &BusinessUpdate{
		Name:             "Neighborhood Deli",
		Description:      "A neighborhood deli serving many tasty sandwiches and soups.",
		ActivationStatus: "active",
	}

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.UpdateBusiness(testID, payload)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &BusinessInfo{}
	json.Unmarshal(businessResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}
