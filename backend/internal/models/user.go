package models

type User struct {
	ID                string            `bson:"_id" json:"_id" form:"_id"`
	Username          string            `bson:"username" json:"username" form:"username"`
	Email             string            `bson:"email" json:"email" form:"email"`
	Password          string            `bson:"password" json:"password" form:"password"`
	PhoneNumber       string            `bson:"phone_number" json:"phone_number" form:"phone_number"`
	IsAdmin           bool              `bson:"is_admin" json:"is_admin" form:"is_admin"`
	DeliveryAddresses []DeliveryAddress `bson:"delivery_addresses" json:"delivery_addresses" form:"delivery_addresses"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
