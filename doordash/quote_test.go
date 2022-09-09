package doordash

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestCreateDeliveryQuote(t *testing.T) {
	timeStamp, _ := time.Parse(time.RFC3339, "2018-08-22T17:20:28Z")
	payload := &NewQuote{
		PickupAddress:                   "901 Market Street 6th Floor San Francisco, CA 94103",
		PickupBusinessName:              "Wells Fargo SF Downtown",
		PickupPhoneNumber:               "+16505555555",
		PickupInstructions:              "Enter gate code 1234 on the callbox.",
		PickupReferenceTag:              "Order number 61",
		PickupExternalBusinessID:        "ase-243-dzs",
		PickupExternalStoreID:           "ase-243-dzs",
		DropoffAddress:                  "901 Market Street 6th Floor San Francisco, CA 94103",
		DropoffBusinessName:             "Wells Fargo SF Downtown",
		DropoffPhoneNumber:              "+16505555555",
		DropoffInstructions:             "Enter gate code 1234 on the callbox.",
		DropoffContactGivenName:         "John",
		DropoffContactFamilyName:        "Doe",
		DropoffContactSendNotifications: true,
		ContactlessDropoff:              false,
		ActionIfUndeliverable:           "return_to_pickup",
		Tip:                             599,
		OrderValue:                      1999,
		Currency:                        "USD",
		PickupWindow: TimeWindow{
			StartTime: timeStamp,
			EndTime:   timeStamp,
		},
		DropoffWindow: TimeWindow{
			StartTime: timeStamp,
			EndTime:   timeStamp,
		},
	}
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != "/drive/v2/quotes" {
			t.Errorf("expected request URL to be /drive/v2/quotes, got %s", reqURL)
		}
		// Send response to be tested
		rw.Write(deliveryResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.CreateDeliveryQuote(payload)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &DeliveryInfo{}
	json.Unmarshal(deliveryResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}

func TestAcceptDeliveryQuote(t *testing.T) {
	testID := "D-12345"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != ("/drive/v2/quotes" + testID + "/accept") {
			t.Errorf("expected request URL to be /drive/v2/quotes/%s/accept, got %s", testID, reqURL)
		}
		// Send response to be tested
		rw.Write(deliveryResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.AcceptDeliveryQuote(testID)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &DeliveryInfo{}
	json.Unmarshal(deliveryResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}
