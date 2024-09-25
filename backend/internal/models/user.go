package models

import "time"

type User struct {
	ID        string    `json:"_id" bson:"_id,omitempty" form:"_id,omitempty"`
	Username  string    `json:"username" bson:"username" form:"username"`
	Email     string    `json:"email" bson:"email" form:"email"`
	Password  string    `json:"password" bson:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" form:"created_at"`
}
