package models

type User struct {
	ID                string           `json:"id"`                
	Username          string           `json:"username"`
	Email             string           `json:"email"`
	Password          string           `json:"password"`
	PhoneNumber       string           `json:"phone_number"`
	IsAdmin           bool             `json:"is_admin"`
	DeliveryAddresses []DeliveryAddress `json:"delivery_addresses"` 
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
