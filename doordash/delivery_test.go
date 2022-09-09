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

var deliveryResponse = []byte(`{
	"external_delivery_id": "D-12345",
	"locale": "en-US",
	"pickup_address": "901 Market Street 6th Floor San Francisco, CA 94103",
	"pickup_business_name": "Wells Fargo SF Downtown",
	"pickup_phone_number": "+16505555555",
	"pickup_instructions": "Enter gate code 1234 on the callbox.",
	"pickup_reference_tag": "Order number 61",
	"pickup_external_business_id": "ase-243-dzs",
	"pickup_external_store_id": "ase-243-dzs",
	"dropoff_address": "901 Market Street 6th Floor San Francisco, CA 94103",
	"dropoff_business_name": "Wells Fargo SF Downtown",
	"dropoff_phone_number": "+16505555555",
	"dropoff_instructions": "Enter gate code 1234 on the callbox.",
	"dropoff_contact_given_name": "John",
	"dropoff_contact_family_name": "Doe",
	"dropoff_contact_send_notifications": true,
	"order_value": 1999,
	"currency": "USD",
	"delivery_status": "quote",
	"cancellation_reason": "cancelled_by_creator",
	"pickup_time_estimated": "2018-08-22T17:20:28Z",
	"pickup_time_actual": "2018-08-22T17:20:28Z",
	"dropoff_time_estimated": "2018-08-22T17:20:28Z",
	"dropoff_time_actual": "2018-08-22T17:20:28Z",
	"return_time_estimated": "2018-08-22T17:20:28Z",
	"return_time_actual": "2018-08-22T17:20:28Z",
	"return_address": "901 Market Street 6th Floor San Francisco, CA 94103",
	"fee": 1900,
	"support_reference": "86313",
	"tracking_url": "https://doordash.com/tracking?id=",
	"dropoff_verification_image_url": "https://doordash-static.s3...",
	"pickup_verification_image_url": "https://doordash-static.s3...",
	"contactless_dropoff": false,
	"action_if_undeliverable": "return_to_pickup",
	"tip": 599
}`)

func TestCreateDelivery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != "/drive/v2/deliveries" {
			t.Errorf("expected request URL to be /drive/v2/deliveries, got %s", reqURL)
		}
		// Send response to be tested
		rw.Write(deliveryResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	payload := &NewDelivery{
		ExternalDeliveryID:  "D-12345",
		PickupAddress:       "901 Market Street 6th Floor San Francisco, CA 94103",
		PickupBusinessName:  "Wells Fargo SF Downtown",
		PickupPhoneNumber:   "+16505555555",
		PickupInstructions:  "Enter gate code 1234 on the callbox.",
		DropoffAddress:      "901 Market Street 6th Floor San Francisco, CA 94103",
		DropoffBusinessName: "Wells Fargo SF Downtown",
		DropoffPhoneNumber:  "+16505555555",
		DropoffInstructions: "Enter gate code 1234 on the callbox.",
		OrderValue:          1999,
	}
	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.CreateDelivery(payload)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &DeliveryInfo{}
	json.Unmarshal(deliveryResponse, want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}

func TestGetDeliveryStatus(t *testing.T) {
	testID := "D-12345"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != ("/drive/v2/deliveries/" + testID) {
			t.Errorf("expected request URL to be /drive/v2/deliveries/%s, got %s", testID, reqURL)
		}
		// Send response to be tested
		rw.Write(deliveryResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.GetDeliveryStatus(testID)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &DeliveryInfo{}
	json.Unmarshal(deliveryResponse, want)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}

}

func TestUpdateDelivery(t *testing.T) {
	testID := "D-12345"
	timeStamp, _ := time.Parse(time.RFC3339, "2018-08-22T17:20:28Z")
	payload := &DeliveryUpdate{
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
		if reqURL != ("/drive/v2/deliveries/" + testID) {
			t.Errorf("expected request URL to be/drive/v2/deliveries/%s, got %s", testID, reqURL)
		}
		// Send response to be tested
		rw.Write(deliveryResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.UpdateDelivery(testID, payload)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &DeliveryInfo{}
	json.Unmarshal(deliveryResponse, want)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}

func TestCancelDelivery(t *testing.T) {
	testID := "D-12345"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		reqURL := req.URL.String()
		if reqURL != ("/drive/v2/deliveries/" + testID) {
			t.Errorf("expected request URL to be/drive/v2/deliveries/%s, got %s", testID, reqURL)
		}
		// Send response to be tested
		rw.Write(deliveryResponse)
	}))
	// Close the server when test finishes
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	client := &Client{url, "token", server.Client()}
	got, err := client.CancelDelivery(testID)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	want := &DeliveryInfo{}
	json.Unmarshal(deliveryResponse, want)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected response to be %v, got %v", want, got)
	}
}
