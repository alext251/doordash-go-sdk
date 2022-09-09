// API Spec: https://developer.doordash.com/en-US/api/drive#tag/Delivery
package doordash

import "time"

// Object for creating a delivery NewQuote
type NewQuote struct {
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

// API Doc: https://developer.doordash.com/en-US/api/drive#tag/Delivery/operation/DeliveryQuote
func (c *Client) CreateDeliveryQuote(q *NewQuote) (*DeliveryInfo, error) {
	return c.makeDeliveryRequest("POST", "drive/v2/quotes", q)
}

// API Doc: https://developer.doordash.com/en-US/api/drive#tag/Delivery/operation/DeliveryQuoteAccept
func (c *Client) AcceptDeliveryQuote(externalDeliveryID string) (*DeliveryInfo, error) {
	return c.makeDeliveryRequest("POST", ("drive/v2/quotes" + externalDeliveryID + "/accept"), nil)
}
