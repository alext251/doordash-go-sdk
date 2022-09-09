// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Delivery
package doordash

import (
	"net/url"
	"time"
)

// Object for creating a new delivery
type NewDelivery struct {
	ExternalDeliveryID              string     `json:"external_delivery_id"`
	Locale                          string     `json:"locale"`
	PickupAddress                   string     `json:"pickup_address"`
	PickupBusinessName              string     `json:"pickup_business_name"`
	PickupPhoneNumber               string     `json:"pickup_phone_number"`
	PickupInstructions              string     `json:"pickup_instructions"`
	PickupReferenceTag              string     `json:"pickup_reference_tag"`
	PickupExternalBusinessID        string     `json:"pickup_external_business_id"`
	PickupExternalStoreID           string     `json:"pickup_external_store_id"`
	DropoffAddress                  string     `json:"dropoff_address"`
	DropoffBusinessName             string     `json:"dropoff_business_name"`
	DropoffPhoneNumber              string     `json:"dropoff_phone_number"`
	DropoffInstructions             string     `json:"dropoff_instructions"`
	DropoffContactGivenName         string     `json:"dropoff_contact_given_name"`
	DropoffContactFamilyName        string     `json:"dropoff_contact_family_name"`
	DropoffContactSendNotifications bool       `json:"dropoff_contact_send_notifications"`
	OrderValue                      int        `json:"order_value"`
	Currency                        string     `json:"currency"`
	PickupTime                      time.Time  `json:"pickup_time"`
	DropoffTime                     time.Time  `json:"dropoff_time"`
	PickupWindow                    TimeWindow `json:"pickup_window"`
	DropoffWindow                   TimeWindow `json:"dropoff_window"`
	ContactlessDropoff              bool       `json:"contactless_dropoff"`
	ActionIfUndeliverable           string     `json:"action_if_undeliverable"`
	Tip                             int        `json:"tip"`
}

// Object for sending a delivery update
type DeliveryUpdate struct {
	PickupAddress                   string    `json:"pickup_address"`
	PickupBusinessName              string    `json:"pickup_business_name"`
	PickupPhoneNumber               string    `json:"pickup_phone_number"`
	PickupInstructions              string    `json:"pickup_instructions"`
	PickupReferenceTag              string    `json:"pickup_reference_tag"`
	PickupExternalBusinessID        string    `json:"pickup_external_business_id"`
	PickupExternalStoreID           string    `json:"pickup_external_store_id"`
	DropoffAddress                  string    `json:"dropoff_address"`
	DropoffBusinessName             string    `json:"dropoff_business_name"`
	DropoffPhoneNumber              string    `json:"dropoff_phone_number"`
	DropoffInstructions             string    `json:"dropoff_instructions"`
	DropoffContactGivenName         string    `json:"dropoff_contact_given_name"`
	DropoffContactFamilyName        string    `json:"dropoff_contact_family_name"`
	DropoffContactSendNotifications bool      `json:"dropoff_contact_send_notifications"`
	ContactlessDropoff              bool      `json:"contactless_dropoff"`
	ActionIfUndeliverable           string    `json:"action_if_undeliverable"`
	Tip                             int       `json:"tip"`
	OrderValue                      int       `json:"order_value"`
	Currency                        string    `json:"currency"`
	PickupTime                      time.Time `json:"pickup_time"`
	DropoffTime                     time.Time `json:"dropoff_time"`
	PickupWindow                    struct {
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	} `json:"pickup_window"`
	DropoffWindow struct {
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
	} `json:"dropoff_window"`
}

// Object containing response information for deliveries and quotes
type DeliveryInfo struct {
	ExternalDeliveryID              string    `json:"external_delivery_id"`
	Locale                          string    `json:"locale"`
	PickupAddress                   string    `json:"pickup_address"`
	PickupBusinessName              string    `json:"pickup_business_name"`
	PickupPhoneNumber               string    `json:"pickup_phone_number"`
	PickupInstructions              string    `json:"pickup_instructions"`
	PickupReferenceTag              string    `json:"pickup_reference_tag"`
	PickupExternalBusinessID        string    `json:"pickup_external_business_id"`
	PickupExternalStoreID           string    `json:"pickup_external_store_id"`
	DropoffAddress                  string    `json:"dropoff_address"`
	DropoffBusinessName             string    `json:"dropoff_business_name"`
	DropoffPhoneNumber              string    `json:"dropoff_phone_number"`
	DropoffInstructions             string    `json:"dropoff_instructions"`
	DropoffContactGivenName         string    `json:"dropoff_contact_given_name"`
	DropoffContactFamilyName        string    `json:"dropoff_contact_family_name"`
	DropoffContactSendNotifications bool      `json:"dropoff_contact_send_notifications"`
	OrderValue                      int       `json:"order_value"`
	Currency                        string    `json:"currency"`
	DeliveryStatus                  string    `json:"delivery_status"`
	CancellationReason              string    `json:"cancellation_reason"`
	PickupTimeEstimated             time.Time `json:"pickup_time_estimated"`
	PickupTimeActual                time.Time `json:"pickup_time_actual"`
	DropoffTimeEstimated            time.Time `json:"dropoff_time_estimated"`
	DropoffTimeActual               time.Time `json:"dropoff_time_actual"`
	ReturnTimeEstimated             time.Time `json:"return_time_estimated"`
	ReturnTimeActual                time.Time `json:"return_time_actual"`
	ReturnAddress                   string    `json:"return_address"`
	Fee                             int       `json:"fee"`
	SupportReference                string    `json:"support_reference"`
	TrackingURL                     string    `json:"tracking_url"`
	DropoffVerificationImageURL     string    `json:"dropoff_verification_image_url"`
	PickupVerificationImageURL      string    `json:"pickup_verification_image_url"`
	ContactlessDropoff              bool      `json:"contactless_dropoff"`
	ActionIfUndeliverable           string    `json:"action_if_undeliverable"`
	Tip                             int       `json:"tip"`
}

type TimeWindow struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// API Doc: https://developer.doordash.com/en-US/api/drive#tag/Delivery/operation/CreateDelivery
func (c *Client) CreateDelivery(d *NewDelivery) (*DeliveryInfo, error) {
	return c.makeDeliveryRequest("POST", "drive/v2/deliveries", d)
}

// API Doc: https://developer.doordash.com/en-US/api/drive#tag/Delivery/operation/GetDelivery
func (c *Client) GetDeliveryStatus(externalDeliveryID string) (*DeliveryInfo, error) {
	return c.makeDeliveryRequest("GET", ("drive/v2/deliveries/" + externalDeliveryID), nil)
}

// API Doc: https://developer.doordash.com/en-US/api/drive#tag/Delivery/operation/UpdateDelivery
func (c *Client) UpdateDelivery(externalDeliveryID string, d *DeliveryUpdate) (*DeliveryInfo, error) {
	return c.makeDeliveryRequest("PATCH", ("drive/v2/deliveries/" + externalDeliveryID), d)
}

// API Doc: https://developer.doordash.com/en-US/api/drive#tag/Delivery/operation/CancelDelivery
func (c *Client) CancelDelivery(externalDeliveryID string) (*DeliveryInfo, error) {
	return c.makeDeliveryRequest("PUT", ("drive/v2/deliveries/" + externalDeliveryID), nil)

}

func (c *Client) makeDeliveryRequest(method string, endpoint string, body interface{}) (*DeliveryInfo, error) {
	var params url.Values

	res := &DeliveryInfo{}
	err := c.makeRequest(method, endpoint, params, body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
