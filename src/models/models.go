package models

import (
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	_ "time"
)

type User struct {
	ID
	First_Name
	Last_Name
	Password
	Email
	Phone
	Token
	Refresh_Token
}
