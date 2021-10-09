package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	userID   string             `json:"user_id,omitempty" bson:"userId,omitempty"`
	mobile   string             `json:"mobile,omitempty" bson:"mobile"`
	email    string             `json:"email" bson:"email"`
	password string             `json:"password" bson:"password"`
	post     string             `json:"post,omitempty" bson:"post"`
}
