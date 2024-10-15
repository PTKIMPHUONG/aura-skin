package models

type User struct {
	ID                string            `bson:"_id" json:"_id" form:"_id"`
	Username          string            `bson:"username" json:"username" form:"username"`
	Email             string            `bson:"email" json:"email" form:"email"`
	Password          string            `bson:"password" json:"password" form:"password"`
	PhoneNumber       string            `bson:"phone_number" json:"phone_number" form:"phone_number"`
	User_image        string            `bson:"user_image" json:"user_image" form:"user_image"`
	IsAdmin           bool              `bson:"is_admin" json:"is_admin" form:"is_admin"`
	IsActive          bool              `bson:"is_active" json:"is_active" form:"is_active"`
	Gender            string            `bson:"gender" json:"gender" form:"gender"`       
	BirthDate         string            `bson:"birth_date" json:"birth_date" form:"birth_date"` 
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
