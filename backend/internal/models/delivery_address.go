package models

type DeliveryAddress struct {
	RecipientName string `json:"recipient_name"`
	ContactNumber string `json:"contact_number"`
	AddressLine   string `json:"address_line"`
	Ward          string `json:"ward"`
	District      string `json:"district"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	IsDefault     bool   `json:"is_default"`
}
